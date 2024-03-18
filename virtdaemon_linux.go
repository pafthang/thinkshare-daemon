package daemon

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/thinkonmay/thinkshare-daemon/persistent/gRPC/packet"
	"github.com/thinkonmay/thinkshare-daemon/utils/libvirt"
	"github.com/thinkonmay/thinkshare-daemon/utils/log"
)

var disk_part = "/home/huyhoang/thinkshare-daemon/utils/libvirt/31134452-4554-4fc8-a0ea-2c88e62ed17f.qcow2"

func HandleVirtdaemon() {
	daemon, err := libvirt.NewVirtDaemon()
	if err != nil {
		log.PushLog("failed to create virtdaemon %s", err.Error())
		return
	}

	network, err := libvirt.NewLibvirtNetwork("enp0s25")
	if err != nil {
		log.PushLog("failed to query gpus %s", err.Error())
		return
	}
	defer network.Close()

	for {
		gpus, err := daemon.ListGPUs()
		if err != nil {
			log.PushLog("failed to query gpus %s", err.Error())
			continue
		}
		for _, g := range gpus {
			if g.Active {
				continue
			}

			iface, err := network.CreateInterface(libvirt.Virtio)
			if err != nil {
				log.PushLog("failed to query gpus %s", err.Error())
				continue
			}

			chain := libvirt.Volume{disk_part, nil}
			err = chain.PushChain(40)
			if err != nil {
				log.PushLog("failed to query gpus %s", err.Error())
				continue
			}

			id := uuid.NewString()
			dom, err := daemon.DeployVM(libvirt.VMLaunchModel{
				ID:            id,
				VCPU:          8,
				RAM:           8,
				GPU:           []libvirt.GPU{g},
				BackingVolume: []libvirt.Volume{chain},
				Interfaces:    []libvirt.Interface{*iface},
				VDriver:       true,
			})

			if err != nil {
				log.PushLog("failed to query gpus %s", err.Error())
				continue
			}

			for {
				time.Sleep(time.Second)
				addr, err := network.FindDomainIPs(dom)
				if err != nil {
					log.PushLog("failed to query gpus %s", err.Error())
					continue
				} else if addr.Ip == nil {
					continue
				}

				client := http.Client{ Timeout: time.Second, }
				resp, err := client.Post(fmt.Sprintf("http://%s:60000/initialize", *addr.Ip), "application/json", strings.NewReader("{}"))
				if err != nil {
					continue
				} else if resp.StatusCode != 200 {
					continue
				}

				log.PushLog("deployed a new worker %s", *addr.Ip)
				break
			}
		}

		time.Sleep(time.Second * 20)
	}
}