package shared

type Config struct {
	Hosts                 []Host `yaml:"hosts"`
	IncludeSshConfigHosts bool   `yaml:"includeSshConfigHosts" default:"true"`
	CheckRemoteForScripts bool   `yaml:"checkRemoteForScripts" default:"false"`
}

type Host struct {
	Host       string
	User       string
	Name       string
	Port       string
	PortTunnel string `yaml:"portTunnel"`
}
