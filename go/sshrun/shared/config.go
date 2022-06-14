package shared

type Config struct {
	Hosts            []Host `yaml:"hosts"`
	IncludeSshConfig bool   `yaml:"includeSshConfig"`
	CheckForScripts  bool   `yaml:"checkForScripts" default:"false"`
}

type Host struct {
	Host       string
	User       string
	Name       string
	Port       string
	PortTunnel string `yaml:"portTunnel"`
}
