package shared

type Config struct {
	Hosts                 []Host `yaml:"hosts"`
	IncludeSshConfigHosts bool   `yaml:"includeSshConfigHosts"`
	CheckRemoteForScripts bool   `yaml:"checkRemoteForScripts"`
}

type Host struct {
	Host       string
	User       string
	Name       string
	Port       string
	PortTunnel string `yaml:"portTunnel"`
}
