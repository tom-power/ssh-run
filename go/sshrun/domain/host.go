package domain

import (
	"encoding/json"
	"fmt"
)

type Host struct {
	Host        string
	User        string
	Name        string
	Port        string
	PortTunnel  string `yaml:"portTunnel"`
	CheckRemote bool   `yaml:"checkRemote"`
}

func (h Host) ToString() string {
	b, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
