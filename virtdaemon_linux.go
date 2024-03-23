package daemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/thinkonmay/thinkshare-daemon/persistent/gRPC/packet"
	"github.com/thinkonmay/thinkshare-daemon/utils/libvirt"
	"github.com/thinkonmay/thinkshare-daemon/utils/log"
)

var disk_part = "/disk/HHDa/default.qcow2"

func HandleVirtdaemon(daemon *Daemon) {
	virt, err := libvirt.NewVirtDaemon()
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

	models := []libvirt.VMLaunchModel{}
	removeVM := func(vm libvirt.Domain) {
		virt.DeleteVM(*vm.Name)
		for _, model := range models {
			if model.ID == *vm.Name {
				for _, v := range model.BackingVolume {
					v.PopChain()
				}
			}
		}
	}

	stop := false
	go func() {
		for {
			ip := <-daemon.close_vm

			vms, err := virt.ListVMs()
			if err != nil {
				daemon.close_vm <- ip
				continue
			}

			for _, vm := range vms {
				if !vm.Running {
					continue
				}

				add, err := network.FindDomainIPs(vm)
				if err != nil {
					daemon.close_vm <- ip
				} else if add.Ip == nil {
					continue
				} else if ip == "all" {
					removeVM(vm)
				} else if *add.Ip == ip {
					removeVM(vm)
				}
			}

			if ip == "all" {
				stop = true
				break
			}
		}
	}()

	update_vms := func() {
		vms, err := virt.ListVMs()
		if err != nil {
			log.PushLog("failed to query gpus %s", err.Error())
			return
		}

		doms := []*packet.WorkerInfor{}
		for _, vm := range vms {
			if !vm.Running {
				continue
			}

			addr, err := network.FindDomainIPs(vm)
			if err != nil {
				log.PushLog("failed to query gpus %s", err.Error())
				continue
			} else if addr.Ip == nil {
				continue
			}

			client := http.Client{Timeout: time.Second}
			resp, err := client.Post(fmt.Sprintf("http://%s:60000/info", *addr.Ip), "application/json", strings.NewReader("{}"))
			if err != nil {
				continue
			} else if resp.StatusCode != 200 {
				continue
			}

			b, _ := io.ReadAll(resp.Body)
			inf := packet.WorkerInfor{}
			err = json.Unmarshal(b, &inf)
			if err != nil {
				log.PushLog("failed to query gpus %s", err.Error())
				continue
			}

			doms = append(doms, &inf)
		}

		daemon.vms = doms
	}

	deploy_vm := func(g libvirt.GPU) {
		if g.Active {
			return
		}

		iface, err := network.CreateInterface(libvirt.Virtio)
		if err != nil {
			log.PushLog("failed to query gpus %s", err.Error())
			return
		}

		chain := libvirt.NewVolume(disk_part)
		err = chain.PushChain(40)
		if err != nil {
			log.PushLog("failed to query gpus %s", err.Error())
			return
		}

		id := uuid.NewString()
		model := libvirt.VMLaunchModel{
			ID:            id,
			VCPU:          8,
			RAM:           8,
			GPU:           []libvirt.GPU{g},
			BackingVolume: []libvirt.Volume{chain},
			Interfaces:    []libvirt.Interface{*iface},
			VDriver:       true,
		}

		models = append(models, model)
		dom, err := virt.DeployVM(model)
		if err != nil {
			log.PushLog("failed to query gpus %s", err.Error())
			return
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

			client := http.Client{Timeout: time.Second}
			resp, err := client.Get(fmt.Sprintf("http://%s:60000/ping", *addr.Ip))
			if err != nil {
				continue
			} else if resp.StatusCode != 200 {
				continue
			}

			log.PushLog("deployed a new worker %s", *addr.Ip)
			break
		}
	}

	for {
		if stop {
			break
		}

		gpus, err := virt.ListGPUs()
		if err != nil {
			log.PushLog("failed to query gpus %s", err.Error())
			continue
		}
		for _, g := range gpus {
			deploy_vm(g)
		}

		update_vms()
		time.Sleep(time.Second * 20)
	}
}
