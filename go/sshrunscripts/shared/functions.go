package shared

import "strings"

var hostsToHostName = func(host Host) string { return host.Name }
var HostsToHostNames = func(hosts []Host, sep string) string { return strings.Join(Map(hosts, hostsToHostName), sep) }
