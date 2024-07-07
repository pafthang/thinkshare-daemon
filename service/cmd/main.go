package cmd

import (
	"fmt"
	"net/http"
	"os"

	daemon "github.com/thinkonmay/thinkshare-daemon"
	"github.com/thinkonmay/thinkshare-daemon/cluster"
	httpp "github.com/thinkonmay/thinkshare-daemon/persistent/http"
	"github.com/thinkonmay/thinkshare-daemon/utils/log"

	"github.com/thinkonmay/thinkshare-daemon/utils/media"
	"github.com/thinkonmay/thinkshare-daemon/utils/signaling"
	ws "github.com/thinkonmay/thinkshare-daemon/utils/signaling/protocol/websocket"
)

func Start(cluster cluster.ClusterConfig, stop chan os.Signal) {
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

	dm := daemon.WebDaemon(grpc, signaling, cluster)
	defer dm.Close()
	stop <- <-stop
}
