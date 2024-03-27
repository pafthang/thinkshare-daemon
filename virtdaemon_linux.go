package daemon

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/melbahja/goph"
	"github.com/thinkonmay/thinkshare-daemon/persistent/gRPC/packet"
	"github.com/thinkonmay/thinkshare-daemon/utils/libvirt"
	"github.com/thinkonmay/thinkshare-daemon/utils/log"
)

type Node struct {
	Ip       string `yaml:"ip"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Role     string `yaml:"role"`

	client *goph.Client
	cancel context.CancelFunc

	internal *packet.WorkerInfor
}

type ClusterConfig struct {
	Nodes []Node `yaml:"nodes"`
}

var (
	dir     = "."
	los     = "./vol.qcow2"
	ldisk   = "./empty.qcow2"
	lbinary = "./daemon"
	virt    *libvirt.VirtDaemon
	network libvirt.Network
	models  []libvirt.VMLaunchModel = []libvirt.VMLaunchModel{}
	nodes   []*Node                 = []*Node{}
)

func init() {
	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	los = fmt.Sprintf("%s/vol.qcow2", dir)
	ldisk = fmt.Sprintf("%s/empty.qcow2", dir)
	lbinary = fmt.Sprintf("%s/daemon", dir)
}

func deinit() {
	for _, node := range nodes {
		node.cancel()
	}
}

func update_gpu(daemon *Daemon) {
	gpus, err := virt.ListGPUs()
	if err != nil {
		log.PushLog("failed to query gpus %s", err.Error())
		return
	}

	gs := []string{}
	for _, g := range gpus {
		if g.Active {
			return
		}
		gs = append(gs, g.Capability.Product.Val)
	}
	daemon.info.GPUs = gs
}

func fileTransfer(node *Node, rfile, lfile string, force bool) error {
	out, err := exec.Command("du", lfile).Output()
	if err != nil {
		return err
	}
	lsize := strings.Split(string(out), "\t")[0]

	out, err = node.client.Run(fmt.Sprintf("du %s", rfile))
	rsize := strings.Split(string(out), "\t")[0]
	if err == nil && force {
		node.client.Run(fmt.Sprintf("rm -f %s", rfile))
	}
	if err != nil || force {
		_, err := exec.Command("sshpass",
			"-p", node.Password,
			"scp", lfile, fmt.Sprintf("%s@%s:%s", node.Username, node.Ip, rfile),
		).Output()
		if err != nil {
			return err
		}

		out, err := node.client.Run(fmt.Sprintf("du %s", rfile))
		if err != nil {
			return err
		}

		rsize = strings.Split(string(out), "\t")[0]
	}

	log.PushLog("%s : local file size %s, remote file size %s", rfile, lsize, rsize)
	return nil
}

func setupNode(node *Node) error {
	client, err := goph.New(node.Username, node.Ip, goph.Password(node.Password))
	if err != nil {
		return err
	}

	node.client = client
	binary := "~/thinkshare-daemon/daemon"
	disk := "~/thinkshare-daemon/empty.qcow2"
	os := "~/thinkshare-daemon/vol.qcow2"

	client.Run("mkdir ~/thinkshare-daemon")

	abs, _ := filepath.Abs(lbinary)
	err = fileTransfer(node, binary, abs, true)
	if err != nil {
		return err
	}

	abs, _ = filepath.Abs(ldisk)
	err = fileTransfer(node, disk, abs, false)
	if err != nil {
		return err
	}

	abs, _ = filepath.Abs(los)
	err = fileTransfer(node, os, abs, false)
	if err != nil {
		return err
	}

	go func() {
		if err != nil {
			return
		}

		log.PushLog("start %s on %s", binary, node.Ip)
		client.Run(fmt.Sprintf("chmod 777 %s", binary))
		client.Run(fmt.Sprintf("chmod 777 %s", disk))

		var ctx context.Context
		ctx, node.cancel = context.WithCancel(context.Background())
		_, err = client.RunContext(ctx, binary)
		if err != nil {
			log.PushLog(err.Error())
		}
	}()
	return nil
}

func (daemon *Daemon) HandleVirtdaemon(cluster *ClusterConfig) {
	if cluster != nil {
		for _, node := range cluster.Nodes {
			err := setupNode(&node)
			if err != nil {
				log.PushLog(err.Error())
				continue
			}

			nodes = append(nodes, &node)
		}
	}

	var err error
	virt, err = libvirt.NewVirtDaemon()
	if err != nil {
		log.PushLog("failed to connect libvirt %s", err.Error())
		return
	}

	network, err = libvirt.NewLibvirtNetwork("enp0s25")
	if err != nil {
		log.PushLog("failed to start network %s", err.Error())
		return
	}
	defer network.Close()

	for {
		update_gpu(daemon)
		QueryInfo(&daemon.info)
		time.Sleep(time.Second * 20)
	}
}

func (daemon *Daemon) DeployVM(session *packet.WorkerSession) (*packet.WorkerInfor, error) {
	if len(session.Vm.GPUs) == 0 {
		return nil, fmt.Errorf("empty gpu")
	}

	g := session.Vm.GPUs[0]
	var gpu *libvirt.GPU = nil
	gpus, err := virt.ListGPUs()
	if err != nil {
		return nil, err
	}
	for _, candidate := range gpus {
		if candidate.Active || candidate.Capability.Product.Val != g {
			continue
		}

		gpu = &candidate
		break
	}

	if gpu == nil {
		return nil, fmt.Errorf("ran out of gpu")
	}

	iface, err := network.CreateInterface(libvirt.Virtio)
	if err != nil {
		return nil, err
	}

	disk := ldisk
	if session.Vm.Volumes != nil && len(session.Vm.Volumes) != 0 {
		disk = fmt.Sprintf("./%s.qcow2", session.Vm.Volumes[0])
	}

	disks, err := prepareVolume(los, disk)
	if err != nil {
		return nil, err
	}

	id := uuid.NewString()
	model := libvirt.VMLaunchModel{
		ID:            id,
		VCPU:          8,
		RAM:           8,
		BackingVolume: disks,
		GPU:           []libvirt.GPU{*gpu},
		Interfaces:    []libvirt.Interface{*iface},
		VDriver:       true,
	}

	models = append(models, model)
	dom, err := virt.DeployVM(model)
	if err != nil {
		return nil, err
	}

	for {
		time.Sleep(time.Second)
		addr, err := network.FindDomainIPs(dom)
		if err != nil {
			log.PushLog("VM ip not available %s", err.Error())
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

		resp, err = client.Get(fmt.Sprintf("http://%s:60000/info", *addr.Ip))
		if err != nil {
			continue
		}
		b, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != 200 {
			log.PushLog(string(b))
			continue
		}

		inf := packet.WorkerInfor{}
		err = json.Unmarshal(b, &inf)
		if err != nil {
			log.PushLog("failed unmarshal reponse body %s", err.Error())
			continue
		} else if inf.PrivateIP == nil || inf.PublicIP == nil {
			log.PushLog("VM address is null, retry")
			continue
		}

		log.PushLog("deployed a new worker %s", *addr.Ip)
		return &inf, nil
	}

}

func (daemon *Daemon) DeployVMonNode(nss *packet.WorkerSession) (*packet.WorkerSession, error) {
	if len(nss.Vm.GPUs) == 0 {
		return nil, fmt.Errorf("empty gpu")
	}

	g := nss.Vm.GPUs[0]
	var node *Node = nil
	for _, n := range nodes {
		resp, err := http.Get(fmt.Sprintf("http://%s:60000/info", n.Ip))
		if err != nil {
			continue
		}

		ss := packet.WorkerInfor{}
		b, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(b, &ss)
		if err != nil {
			log.PushLog(err.Error())
			continue
		}

		for _, gpu := range ss.GPUs {
			if gpu == g {
				cp := *n
				node = &cp
				break
			}
		}

		if node != nil {
			break
		}
	}

	if node == nil {
		return nil, fmt.Errorf("cluster ran out of gpu")
	}

	log.PushLog("deploying VM on node %s", node.Ip)
	b, _ := json.Marshal(nss)
	resp, err := http.Post(
		fmt.Sprintf("http://%s:60000/new", node.Ip),
		"application/json",
		strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	b, _ = io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(string(b))
	}

	err = json.Unmarshal(b, &nss)
	if err != nil {
		return nil, err
	}

	return nss, nil
}

func (daemon *Daemon) ShutdownVM(info *packet.WorkerInfor) error {
	removeVM := func(vm libvirt.Domain) {
		virt.DeleteVM(*vm.Name)
		for _, model := range models {
			if model.ID == *vm.Name {
				for _, v := range model.BackingVolume {
					if v.Disposable {
						v.PopChain()
					}
				}
			}
		}
	}

	vms, err := virt.ListVMs()
	if err != nil {
		return err
	}

	for _, vm := range vms {
		ip, err := network.FindDomainIPs(vm)
		if err != nil {
			continue
		}

		if vm.Running && *ip.Ip == *info.PrivateIP {
			removeVM(vm)
			return nil
		}
	}

	return fmt.Errorf("vm not found")
}

func (daemon *Daemon) HandleSessionForward(ss *packet.WorkerSession, command string) (*packet.WorkerSession, error) {
	if ss.Target == nil {
		for _, node := range nodes {
			for _, session := range node.internal.Sessions {
				if session.Id != ss.Id {
					continue
				}

				log.PushLog("forwarding command %s to node %s", command, node.Ip)

				b, _ := json.Marshal(ss)
				resp, err := http.Post(
					fmt.Sprintf("http://%s:60000/%s", node.Ip, command),
					"application/json",
					strings.NewReader(string(b)))
				if err != nil {
					log.PushLog("failed to request %s", err.Error())
					continue
				}

				b, _ = io.ReadAll(resp.Body)
				if resp.StatusCode != 200 {
					log.PushLog("failed to request %s", string(b))
					continue
				}

				nss := packet.WorkerSession{}
				err = json.Unmarshal(b, &nss)
				if err != nil {
					log.PushLog("failed to request %s", err.Error())
					continue
				}

				return &nss, nil
			}
		}
		return nil, fmt.Errorf("no session found on any node")
	}

	for _, session := range daemon.info.Sessions {
		if session.Id != *ss.Target || session.Vm == nil {
			continue
		}

		log.PushLog("forwarding command %s to vm %s", command, *session.Vm.PrivateIP)

		nss := *ss
		nss.Target = nil
		b, _ := json.Marshal(nss)
		resp, err := http.Post(
			fmt.Sprintf("http://%s:60000/%s", *session.Vm.PrivateIP, command),
			"application/json",
			strings.NewReader(string(b)))
		if err != nil {
			log.PushLog("failed to request %s", err.Error())
			continue
		}

		b, _ = io.ReadAll(resp.Body)
		if resp.StatusCode != 200 {
			log.PushLog("failed to request %s", string(b))
			continue
		}

		err = json.Unmarshal(b, &nss)
		if err != nil {
			log.PushLog("failed to request %s", err.Error())
			continue
		}

		return &nss, nil
	}

	for _, node := range nodes {
		for _, session := range node.internal.Sessions {
			if session.Id != *ss.Target || session.Vm == nil {
				continue
			}

			log.PushLog("forwarding command %s to node %s, vm %s", command, node.Ip, *session.Vm.PrivateIP)

			b, _ := json.Marshal(ss)
			resp, err := http.Post(
				fmt.Sprintf("http://%s:60000/%s", node.Ip, command),
				"application/json",
				strings.NewReader(string(b)))
			if err != nil {
				log.PushLog("failed to request %s", err.Error())
				continue
			}

			b, _ = io.ReadAll(resp.Body)
			if resp.StatusCode != 200 {
				log.PushLog("failed to request %s", string(b))
				continue
			}

			nss := packet.WorkerSession{}
			err = json.Unmarshal(b, &nss)
			if err != nil {
				log.PushLog("failed to request %s", err.Error())
				continue
			}

			return &nss, nil
		}
	}

	return nil, fmt.Errorf("no receiver detected")
}

func (daemon *Daemon) HandleSignaling(token string) (*string, bool) {
	for _, s := range daemon.info.Sessions {
		if s.Id == token && s.Vm != nil {
			return s.Vm.PrivateIP, true
		}
	}
	for _, node := range nodes {
		if node.internal == nil {
			continue
		}

		for _, s := range node.internal.Sessions {
			if s.Id == token && s.Vm != nil {
				return &node.Ip, false
			}

		}

	}
	return nil, false
}

func QueryInfo(info *packet.WorkerInfor) {
	for _, session := range info.Sessions {
		if session.Vm == nil {
			continue
		}

		resp, err := http.Get(fmt.Sprintf("http://%s:60000/info", *session.Vm.PrivateIP))
		if err != nil {
			log.PushLog(err.Error())
			continue
		}

		ss := packet.WorkerInfor{}
		b, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(b, &ss)
		if err != nil {
			log.PushLog(err.Error())
			continue
		}

		session.Vm = &ss
	}

	for _, node := range nodes {
		resp, err := http.Get(fmt.Sprintf("http://%s:60000/info", node.Ip))
		if err != nil {
			continue
		}

		ss := packet.WorkerInfor{}
		b, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(b, &ss)
		if err != nil {
			continue
		}

		node.internal = &ss
	}

	vms, err := virt.ListVMs()
	if err != nil {
		return
	}

	in_use := []string{}
	for _, vm := range vms {
		ip, err := network.FindDomainIPs(vm)
		if err != nil {
			continue
		}

		for _, model := range models {
			if model.ID != *vm.Name {
				continue
			}

			for _, vol := range model.BackingVolume {
				for _, f := range vol.AllFiles() {
					volume_id := strings.Split(filepath.Base(f), ".qcow2")[0]
					if uuid.Validate(volume_id) != nil {
						continue
					}

					in_use = append(in_use, volume_id)
				}

			}

			var volume_id *string = nil
			if len(model.BackingVolume) > 1 {
				volume_id = &strings.Split(filepath.Base(model.BackingVolume[1].Path), ".qcow2")[0]
			}

			if volume_id == nil {
				continue
			}

			for _, ss := range info.Sessions {
				if ss.Vm == nil || *ss.Vm.PrivateIP != *ip.Ip {
					continue
				}

				ss.Vm.Volumes = []string{*volume_id}
			}
		}
	}

	files, err := os.ReadDir(".")
	if err != nil {
		log.PushLog(err.Error())
		return
	}

	vols := []string{}
	for _, f := range files {
		if f.IsDir() || !strings.Contains(f.Name(), "qcow2") {
			continue
		}

		volume_id := strings.Split(filepath.Base(f.Name()), ".qcow2")[0]

		if uuid.Validate(volume_id) != nil {
			continue
		}

		vols = append(vols, volume_id)
	}

	available := []string{}
	for _, vol := range vols {
		found := false
		for _, iu := range in_use {
			if iu == vol {
				found = true
			}
		}
		if found {
			continue
		}

		available = append(available, vol)
	}

	info.Volumes = available
}

func InfoBuilder(cp packet.WorkerInfor) packet.WorkerInfor {
	for _, node := range nodes {
		cp.Sessions = append(cp.Sessions, node.internal.Sessions...)
		cp.GPUs = append(cp.GPUs, node.internal.GPUs...)
	}
	return cp
}

func prepareVolume(los string, ldisk string) ([]libvirt.Volume, error) {
	chain_os := libvirt.NewVolume(los)
	err := chain_os.PushChain(40)
	if err != nil {
		return []libvirt.Volume{}, err
	}

	result, err := exec.Command("qemu-img", "info", ldisk, "--output", "json").Output()
	if err != nil {
		return []libvirt.Volume{}, err
	}

	var chain_disk *libvirt.Volume = nil
	result_data := struct {
		Backing *string `json:"backing-filename"`
	}{}
	err = json.Unmarshal(result, &result_data)
	if err != nil {
		return []libvirt.Volume{}, err
	} else if result_data.Backing != nil {
		chain_disk = libvirt.NewVolume(ldisk, *result_data.Backing)
	} else {
		chain_disk = libvirt.NewVolume(ldisk)
		err = chain_disk.PushChain(150)
		if err != nil {
			return []libvirt.Volume{}, err
		}
	}

	chain_disk.Disposable = false
	return []libvirt.Volume{*chain_os, *chain_disk}, nil
}
