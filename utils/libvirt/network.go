package libvirt

const (
	Virtio = "virtio"
)

type DomainAddress struct {
	Mac *string `json:"mac"`
	Ip  *string `json:"ip"`
}

type Network interface {
	FindDomainIPs(dom Domain) (DomainAddress, error)
	CreateInterface(driver,network string) (*Interface, error)
	Close()
}