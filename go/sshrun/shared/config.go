package shared

type Config struct {
	Hosts            []Host `yaml:"hosts"`
	IncludeSshConfig bool   `yaml:"includeSshConfig"`
}

type Host struct {
	Host            string
	User            string
	Name            string
	Port            string
	PortTunnel      string `yaml:"portTunnel"`
	CheckForScripts bool   `yaml:"checkForScripts"`
}
