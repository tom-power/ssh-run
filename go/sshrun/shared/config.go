package shared

type Config struct {
	Hosts              []Host `yaml:"hosts"`
	HostsFromSshConfig bool   `yaml:"hostsFromSshConfig" default:"false"`
}

type Host struct {
	Host            string
	User            string
	Name            string
	Port            string
	PortTunnel      string `yaml:"portTunnel"`
	CheckForScripts bool   `yaml:"checkForScripts"`
}
