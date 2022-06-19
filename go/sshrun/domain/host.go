package domain

import (
	"fmt"
	"strings"
)

type Host struct {
	Host        string
	User        string
	Name        string
	Port        string
	PortTunnel  string `yaml:"portTunnel"`
	CheckRemote bool
}

func (h Host) ToString() string {
	return strings.ReplaceAll(fmt.Sprintf("%#v", h), "domain.", "")
}
