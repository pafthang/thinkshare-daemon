package cmd

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	daemon "github.com/thinkonmay/thinkshare-daemon"
	"github.com/thinkonmay/thinkshare-daemon/cluster"
	httpp "github.com/thinkonmay/thinkshare-daemon/persistent/http"
	"github.com/thinkonmay/thinkshare-daemon/utils/log"
	"gopkg.in/yaml.v3"

	"github.com/thinkonmay/thinkshare-daemon/utils/media"
	"github.com/thinkonmay/thinkshare-daemon/utils/signaling"
	ws "github.com/thinkonmay/thinkshare-daemon/utils/signaling/protocol/websocket"
)

func Start(stop chan os.Signal) {
	media.ActivateVirtualDriver()
	defer media.DeactivateVirtualDriver()

	grpc, err := httpp.InitHttppServer()
	if err != nil {
		log.PushLog("failed to setup grpc: %s", err.Error())
		return
	}
	defer grpc.Stop()

	signaling := signaling.InitSignallingServer(
		ws.InitSignallingHttp("/handshake/client"),
		ws.InitSignallingHttp("/handshake/server"),
	)

	srv := &http.Server{Addr: fmt.Sprintf(":%d", daemon.Httpport)}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.PushLog(err.Error())
		}
	}()
	defer srv.Close()

	log.PushLog("starting worker daemon")

	exe, _ := os.Executable()
	base_dir, _ := filepath.Abs(filepath.Dir(exe))
	manifest := fmt.Sprintf("%s/cluster.yaml", base_dir)
	if _, err := os.Stat(manifest); errors.Is(err, os.ErrNotExist) {
		i := (*net.Interface)(nil)
		if ifaces, err := net.Interfaces(); err == nil {
			for _, local_if := range ifaces {
				if local_if.Flags&net.FlagLoopback > 0 ||
					local_if.Flags&net.FlagRunning == 0 {
					continue
				}

				i = &local_if
				break
			}
		}

		if i == nil {
			panic(fmt.Errorf("no log file available"))
		}

		content, _ := yaml.Marshal(cluster.ClusterConfigManifest{
			Nodes: []cluster.NodeManifest{},
			Peers: []cluster.PeerManifest{},
			Local: cluster.Host{
				Interface: i.Name,
			},
		})

		os.WriteFile(manifest, content, 0777)
	}

	dm := daemon.WebDaemon(grpc, signaling, manifest, fmt.Sprintf("%s/web/dist", base_dir))
	defer dm.Close()
	stop <- <-stop
}
