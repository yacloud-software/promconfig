package main

import (
	"fmt"
	"time"

	"golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/promconfig/targets"
)

func init() {
	go all_server_loop()
}

func all_server_loop() {
	t := time.Duration(3) * time.Second
	last_version := 0
	for {
		time.Sleep(t)
		t = time.Duration(60) * time.Second
		v, ct := GetServerListVersion()
		if time.Since(ct) < time.Duration(30)*time.Second {
			// wait for it to settle
			continue
		}
		if v == last_version {
			// no change
			continue
		}
		sl := GetServerList()
		err := handle_all_server_config(sl)
		if err != nil {
			fmt.Printf("[allserver] failed to config for all servers: %s\n", errors.ErrorString(err))
		} else {
			last_version = v
		}
	}
}

func handle_all_server_config(servers []*known_server) error {
	fmt.Printf("Handling all server update for %d servers\n", len(servers))
	// create node exporter entries
	rep := &promconfig.Reporter{Reporter: "allservers"}
	tl := &promconfig.TargetList{
		Reporter: rep,
	}
	t := &promconfig.Target{
		Name:         "prometheus-node-exporter",
		TargetConfig: &promconfig.EmbeddedTargetConfig{HTTPOnly: true, MetricsPath: "/metrics"},
	}
	tl.Targets = append(tl.Targets, t)
	for _, s := range servers {
		t.Addresses = append(t.Addresses, s.ip+":9100")
	}
	err := targets.UpdateTargets(tl)
	return err
}
