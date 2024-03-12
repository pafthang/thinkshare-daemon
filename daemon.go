package daemon

import (
	"encoding/base64"
	"fmt"
	"os/exec"
	"sync"

	"github.com/thinkonmay/thinkshare-daemon/childprocess"
	"github.com/thinkonmay/thinkshare-daemon/persistent"
	"github.com/thinkonmay/thinkshare-daemon/persistent/gRPC/packet"
	"github.com/thinkonmay/thinkshare-daemon/utils/log"
	"github.com/thinkonmay/thinkshare-daemon/utils/media"
	"github.com/thinkonmay/thinkshare-daemon/utils/path"
	"github.com/thinkonmay/thinkshare-daemon/utils/system"
)

type internalWorkerSession struct {
	packet.WorkerSession
	childprocess []childprocess.ProcessID
}

type Daemon struct {
	childprocess *childprocess.ChildProcesses
	persist      persistent.Persistent

	mutex *sync.Mutex

	session []internalWorkerSession
	log     int
}

type DaemonOption struct {
	Turn *struct {
		Username string `json:"username"`
		Password string `json:"password"`
		MinPort  int    `json:"min_port"`
		MaxPort  int    `json:"max_port"`
		Port     int
	} `json:"turn"`
}

func WebDaemon(persistent persistent.Persistent,
	options DaemonOption) *Daemon {
	daemon := &Daemon{
		mutex:   &sync.Mutex{},
		session: []internalWorkerSession{},
		persist: persistent,
		childprocess: childprocess.NewChildProcessSystem(func(proc, log string) {
			fmt.Println(proc + " : " + log)
			persistent.Log(proc, "childprocess", log)
		}),
		log: log.TakeLog(func(log string) {
			persistent.Log("daemon.exe", "infor", log)
		}),
	}

	go func() {
		for {
			infor, err := system.GetInfor()
			if err != nil {
				log.PushLog("error get sysinfor : %s", err.Error())
				return
			}

			if options.Turn != nil {
				port := int32(options.Turn.Port)
				infor.TurnPort = &port
			} else {
				infor.TurnPort = nil
			}
			daemon.persist.Infor(infor)
		}
	}()
	daemon.persist.Sessions(func() []packet.WorkerSession {
		sessions := []packet.WorkerSession{}
		for _, iws := range daemon.session {
			sessions = append(sessions, packet.WorkerSession{
				Id:        iws.Id,
				Timestamp: iws.Timestamp,
				Thinkmay:  iws.Thinkmay,
				Sunshine:  iws.Sunshine,
			})
		}
		return sessions
	})

	daemon.persist.RecvSession(func(ss *packet.WorkerSession) error {
		process := []childprocess.ProcessID{}

		err := fmt.Errorf("no session configured")
		if ss.Display != nil {
			name, index := media.StartVirtualDisplay(
				int(ss.Display.ScreenWidth),
				int(ss.Display.ScreenHeight),
			)
			val := int32(index)
			ss.Display.DisplayName, ss.Display.DisplayIndex = &name, &val
		} else {
			ss.Display = &packet.DisplaySession{
				DisplayName:  &media.Displays()[0],
				DisplayIndex: nil,
			}
		}
		if ss.Thinkmay != nil {
			process, err = daemon.handleHub(ss)
		}
		if ss.Sunshine != nil {
			process, err = daemon.handleSunshine(ss)
		}
		if err != nil {
			log.PushLog("session failed")
			return err
		}

		log.PushLog("session creation successful")
		daemon.session = append(daemon.session,
			internalWorkerSession{
				*ss, process,
			})

		return nil
	})

	go func() {
		for {
			ss := daemon.persist.ClosedSession()
			log.PushLog("terminating session %d", ss)
			queue := []internalWorkerSession{}
			for _, ws := range daemon.session {
				if ws.Display.DisplayIndex != nil {
					media.RemoveVirtualDisplay(int(*ws.Display.DisplayIndex))
				}
				if int(ws.Id) == ss {
					for _, pi := range ws.childprocess {
						daemon.childprocess.CloseID(pi)
					}
				} else {
					queue = append(queue, ws)
				}
			}

			if len(daemon.session) == len(queue) {
				log.PushLog("no session terminated, total session : %d", len(daemon.session))
			} else {
				daemon.session = queue
			}
		}
	}()

	return daemon
}

func (daemon *Daemon) Close() {
	daemon.childprocess.CloseAll()
	log.RemoveCallback(daemon.log)
}

func (daemon *Daemon) handleHub(current *packet.WorkerSession) ([]childprocess.ProcessID, error) {
	daemon.mutex.Lock()
	defer daemon.mutex.Unlock()

	webrtcHash, displayHash :=
		string(base64.StdEncoding.EncodeToString([]byte(current.Thinkmay.WebrtcConfig))),
		string(base64.StdEncoding.EncodeToString([]byte(*current.Display.DisplayName)))

	hub_path, err := path.FindProcessPath("", "hub.exe")
	if err != nil {
		return nil, err
	}
	cmd := []string{
		"--webrtc", webrtcHash,
		"--display", displayHash,
	}

	video, err := daemon.childprocess.NewChildProcess(exec.Command(hub_path, cmd...), true)
	if err != nil {
		return nil, err
	}

	return []childprocess.ProcessID{video}, nil
}

func (daemon *Daemon) handleSunshine(current *packet.WorkerSession) ([]childprocess.ProcessID, error) {
	daemon.mutex.Lock()
	defer daemon.mutex.Unlock()

	hub_path, err := path.FindProcessPath("", "sunshine.exe")
	if err != nil {
		return nil, err
	}

	cmd := []string{
		"--username", current.Sunshine.Username,
		"--password", current.Sunshine.Password,
	}

	id, err := daemon.childprocess.NewChildProcess(exec.Command(hub_path, cmd...), true)
	if err != nil {
		return nil, err
	}

	return []childprocess.ProcessID{id}, nil
}
