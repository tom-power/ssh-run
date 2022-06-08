package shared

type Host struct {
	Host            string
	User            string
	Name            string
	Port            string
	PortTunnel      string `yaml:"portTunnel"`
	CheckForScripts bool   `yaml:"checkForScripts"`
}
