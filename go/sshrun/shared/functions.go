package shared

import "strings"

func HostsToHostNames(hosts []Host, sep string) string {
	return strings.Join(Map(hosts, hostsToHostName), sep)
}
func hostsToHostName(host Host) string { return host.Name }

func DefaultString(v string, defaultV string) string {
	if v == "" {
		return defaultV
	}
	return v
}
