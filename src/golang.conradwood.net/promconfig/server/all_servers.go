package main

import (
	"fmt"
	"sync"
	"time"

	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
)

var (
	servers_lock            sync.Mutex
	server_list_version     int
	server_list_last_change time.Time
	all_known_servers       = make(map[string]*known_server)
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
