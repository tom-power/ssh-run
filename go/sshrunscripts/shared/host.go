package shared

type Host struct {
	Name            string
	User            string
	Ip              string
	PortSsh         string `yaml:"portSsh"`
	PortTunnel      string `yaml:"portTunnel"`
	CheckForScripts bool   `yaml:"checkForScripts"`
}
