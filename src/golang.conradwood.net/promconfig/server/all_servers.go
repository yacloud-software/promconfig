package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"strings"
	"sync"
	"time"
)

var (
	servers_lock            sync.Mutex
	server_list_version     int
	server_list_last_change time.Time
	all_known_servers       = make(map[string]*known_server)
	ignore_servers          = flag.String("ignore_servers", "", "a comma delimited list of ipaddresses to ignore from the all_servers (e.g. node) list (only). does not affect registry entries")
)

type known_server struct {
	ip string
}

func AddServer(ip string) {
	ip, _, _, err := utils.ParseIP(ip)
	if err != nil {
		fmt.Printf("Server \"%s\" does not have a valid ip: %s\n", ip, errors.ErrorString(err))
		return
	}
	if isIgnored(ip) {
		return
	}
	added := false
	servers_lock.Lock()
	ks := all_known_servers[ip]
	if ks == nil {
		added = true
		ks = &known_server{ip: ip}
		all_known_servers[ip] = ks
	}
	if added {
		server_list_version++
		server_list_last_change = time.Now()
	}
	servers_lock.Unlock()
}
func GetServerListVersion() (int, time.Time) {
	return server_list_version, server_list_last_change
}
func GetServerList() []*known_server {
	var res []*known_server
	servers_lock.Lock()
	for _, ks := range all_known_servers {
		res = append(res, ks)
	}
	servers_lock.Unlock()
	return res
}

func isIgnored(ip string) bool {
	iplist := *ignore_servers
	if len(iplist) == 0 {
		return false
	}
	for _, ipl := range strings.Split(iplist, ",") {
		ipl = strings.Trim(ipl, " ")
		if ipl == ip {
			return true
		}
	}
	return false
}
