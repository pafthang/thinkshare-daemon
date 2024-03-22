package daemon

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/thinkonmay/thinkshare-daemon/childprocess"
	"github.com/thinkonmay/thinkshare-daemon/persistent"
	"github.com/thinkonmay/thinkshare-daemon/persistent/gRPC/packet"
	"github.com/thinkonmay/thinkshare-daemon/utils/log"
	"github.com/thinkonmay/thinkshare-daemon/utils/media"
	"github.com/thinkonmay/thinkshare-daemon/utils/path"
	"github.com/thinkonmay/thinkshare-daemon/utils/system"
	"github.com/thinkonmay/thinkshare-daemon/utils/turn"
)

type internalWorkerSession struct {
	packet.WorkerSession
	childprocess []childprocess.ProcessID
	turn_server  *turn.TurnServer
}

type Daemon struct {
	vms          []*packet.WorkerInfor
	childprocess *childprocess.ChildProcesses
	persist      persistent.Persistent
	close_vm     chan string

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
	infor, err := system.GetInfor()
	if err != nil {
		log.PushLog("failed to get info %s", err.Error())
		time.Sleep(time.Second)
		return WebDaemon(persistent, options)
	}

	daemon := &Daemon{
		close_vm: make(chan string, 2),
		mutex:    &sync.Mutex{},
		session:  []internalWorkerSession{},
		vms:      []*packet.WorkerInfor{},
		persist:  persistent,
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

			infor.VMs = daemon.vms
			daemon.persist.Infor(infor)
		}
	}()
	daemon.persist.Sessions(func() []packet.WorkerSession {
		sessions := []packet.WorkerSession{}
		for _, iws := range daemon.session {
			sessions = append(sessions, packet.WorkerSession{
				Id:       iws.Id,
				Target:   infor,
				Thinkmay: iws.Thinkmay,
				Sunshine: iws.Sunshine,
			})
		}

		for _, vm := range daemon.vms {
			resp, err := http.Get(fmt.Sprintf("http://%s:60000/sessions", *vm.PrivateIP))
			if err != nil {
				continue
			}

			ss := packet.WorkerSession{}
			b, _ := io.ReadAll(resp.Body)
			err = json.Unmarshal(b, &ss)
			if err != nil {
				continue
			}

			sessions = append(sessions, ss)
		}

		return sessions
	})

	go HandleVirtdaemon(daemon)
	go func() {
		for {
			worker := daemon.persist.ClosedWorker()
			if worker.PrivateIP == infor.PrivateIP &&
				worker.Hostname == infor.Hostname {
				daemon.Close()
				break
			}

			daemon.close_vm <- *worker.PrivateIP
		}
	}()
	daemon.persist.RecvSession(func(ss *packet.WorkerSession) error {
		process := []childprocess.ProcessID{}
		var t *turn.TurnServer = nil

		err := fmt.Errorf("no session configured")
		if ss.Turn != nil {
			t, err = turn.Open(
				ss.Turn.Username,
				ss.Turn.Password,
				int(ss.Turn.MinPort),
				int(ss.Turn.MaxPort),
				int(ss.Turn.Port),
			)
		}

		if ss.Target != nil &&
			(*ss.Target.PrivateIP != *infor.PrivateIP ||
			ss.Target.Hostname != infor.Hostname) {
			for _, vm := range daemon.vms {
				if *vm.PrivateIP == *ss.Target.PrivateIP {
					b, _ := json.Marshal(ss)
					resp, err := http.Post(
						fmt.Sprintf("http://%s:60000/new", *vm.PrivateIP), 
						"application/json", 
						strings.NewReader(string(b)))
					if err != nil {
						return err
					} else if resp.StatusCode != 200 {
						b, _ := io.ReadAll(resp.Body)
						return fmt.Errorf(string(b))
					}

					break
				}
			}
		} else {
			if ss.Display != nil {
				name, index, err := media.StartVirtualDisplay(
					int(ss.Display.ScreenWidth),
					int(ss.Display.ScreenHeight),
				)
				if err == nil {
					return err
				}
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
		}

		if err != nil {
			log.PushLog("session failed")
			return err
		}

		log.PushLog("session creation successful")
		daemon.session = append(daemon.session,
			internalWorkerSession{
				*ss, process, t,
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
				if ws.turn_server != nil {
					ws.turn_server.Close()
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
	daemon.close_vm <- "all"
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

	video, err := daemon.childprocess.NewChildProcess(exec.Command(hub_path, cmd...))
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
		"--port", current.Sunshine.Port,
	}

	id, err := daemon.childprocess.NewChildProcess(exec.Command(hub_path, cmd...))
	if err != nil {
		return nil, err
	}

	return []childprocess.ProcessID{id}, nil
}
