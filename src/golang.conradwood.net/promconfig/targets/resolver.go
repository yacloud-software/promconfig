package targets

import (
	"fmt"
	"net"
	"sort"
	"strings"
	"time"

	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/utils"
)

var (
	hostcache = cache.New("hostcache", time.Duration(5)*time.Minute, 1000)
)

type hostcache_entry struct {
	hostname string
}

func HostName(ip string) (string, uint32) {
	nip, port, _, err := utils.ParseIP(ip)
	if err != nil {
		fmt.Printf("inv ip (%s) :%s\n", ip, err)
		return "", 0
	}
	o := hostcache.Get(nip)
	if o != nil {
		hce := o.(*hostcache_entry)
		return hce.hostname, port
	}
	addr, err := net.LookupAddr(nip)
	if err != nil {
		utils.PrintStack("DNS Error")
		fmt.Printf("Lookup failed: %s\n", err)
		return "", 0
	}
	if len(addr) == 0 {
		return "", 0
	}
	sort.Slice(addr, func(i, j int) bool {
		return addr[i] < addr[j]
	})
	res := strings.TrimSuffix(addr[0], ".")
	hce := &hostcache_entry{hostname: res}
	hostcache.Put(nip, hce)
	return res, port

}
