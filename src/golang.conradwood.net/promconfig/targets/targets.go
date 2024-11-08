package targets

// maintain prometheus yaml files
// for the targets

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	pb "golang.conradwood.net/apis/promconfig"
	reg "golang.conradwood.net/apis/registry"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/utils"
	"golang.org/x/sys/unix"
)

const (
	YAML_ID = "# this yaml file was written by promconfig"
)

var (
	targetsdir   = flag.String("prometheus_targets", "/etc/prometheus/promconfig/targets", "Directory to store targets for prometheus in. (empty==no targetfiles are maintained)")
	templatefile = flag.String("prometheus_config_template", "/etc/prometheus/config_template.yaml", "A prometheus config file to use as template (prefix)")
	pmcfgfile    = flag.String("prometheus_config_file", "/etc/prometheus/config.yaml", "If not empty, maintain a prometheus config file")
	sample_limit = flag.Int("prometheus_sample_limit", 4000, "set to non-zero to enable prometheus 'sample_limit' option")
	debug        = flag.Bool("debug_targets", false, "debug target writing")
	targets      *targetCache
	promlock     sync.Mutex
)

type targetCache struct {
	Targets map[string]*targetCacheEntry // reporter -> targetlist
}
type targetCacheEntry struct {
	reporter      string
	list          *pb.TargetList
	lastRefreshed time.Time
}

func (t *targetCache) Names() []string {
	mr := make(map[string]int)
	for _, v := range targets.Targets {
		for _, t := range v.list.Targets {
			mr[t.Name] = 1
		}
	}
	var res []string
	for k, _ := range mr {
		res = append(res, k)
	}
	return res
}

type targetaddress struct {
	Name     string
	Reporter *pb.Reporter
	Address  string
	reported time.Time
	httponly bool
}

func (t *targetCache) TargetsByName(name string) []*targetaddress {
	var res []*targetaddress
	for _, v := range targets.Targets {
		for _, t := range v.list.Targets {
			if t.Name != name {
				continue
			}
			for _, adr := range t.Addresses {
				ta := &targetaddress{
					Name:     t.Name,
					Reporter: t.Reporter,
					Address:  adr,
					reported: v.lastRefreshed,
					httponly: t.HTTPOnly,
				}
				res = append(res, ta)
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Address < res[j].Address
	})
	return res
}
func (t *targetCache) AllTargets() []*targetaddress {
	var res []*targetaddress
	for _, n := range t.Names() {
		tc := t.TargetsByName(n)
		res = append(res, tc...)
	}
	return res
}
func UpdateTargets(req *pb.TargetList) error {
	if *targetsdir == "" {
		if *debug {
			fmt.Printf("no targetsdir!!\n")
		}
		return nil
	}
	if targets == nil {
		targets = &targetCache{Targets: make(map[string]*targetCacheEntry)}
	}
	if req.Reporter == nil || req.Reporter.Reporter == "" {
		return fmt.Errorf("Reporter missing. please fill proto before submitting")
	}
	unix.Umask(000)
	// we don't want to run multi-threaded, we're writing files!
	promlock.Lock()
	defer promlock.Unlock()
	repname := req.Reporter.Reporter
	i := strings.Index(repname, ":")
	if i >= 0 {
		repname = repname[:i]
	}
	repname = strings.ToLower(repname)
	tce := targets.Targets[repname]
	if tce == nil {
		tce = &targetCacheEntry{reporter: repname}
		targets.Targets[repname] = tce
	}
	tce.list = req
	tce.lastRefreshed = time.Now()
	if *debug {
		fmt.Printf("Reported %d new targets\n", len(tce.list.Targets))
	}
	for _, t := range tce.list.Targets {
		t.Reporter = req.Reporter
		t.Reporter.Reporter = repname
	}
	if *debug {
		fmt.Printf("Received %d targets from %s\n", len(targets.Targets), repname)
	}

	err := writeTargets()
	if err != nil {
		fmt.Printf("Failed to write targets: %s\n", err)
		return err
	}
	if *pmcfgfile != "" {
		RewriteConfigFile()
	}
	return nil
}

func writeTargets() error {
	var err error
	dir := fmt.Sprintf("%s/%s", *targetsdir, "registry")
	if !utils.FileExists(dir) {
		fmt.Printf("Dir %s does not (yet) exist\n", dir)
		os.Mkdir(dir, 0777)
	}
	fos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	files := make(map[string]int)
	for _, f := range fos {
		ff := fmt.Sprintf("%s/%s", dir, f.Name())
		files[ff] = 1
	}

	for _, name := range targets.Names() {
		t := targets.TargetsByName(name)
		fname := fmt.Sprintf("%s/%s.yaml", dir, name)
		delete(files, fname)
		s := fmt.Sprintf("%s\n", YAML_ID)
		s = s + fmt.Sprintf("# targets for %s\n", name)
		s = s + fmt.Sprintf("- targets:\n")
		for _, adr := range t {
			comment := fmt.Sprintf(" Reported by \"%s\" on %s", adr.Reporter.Reporter, utils.TimeString(adr.reported))
			s = s + fmt.Sprintf("   # %s\n", comment)
			s = s + fmt.Sprintf("   - %s\n", adr.Address)
		}
		/*
			s = s + fmt.Sprintf("  labels:\n")
			s = s + fmt.Sprintf("     service: %s\n", name)
		*/
		e := ioutil.WriteFile(fname, []byte(s), 0666)
		if e != nil {
			err = e
		}
		//fmt.Println(s)
	}

	// delete yaml files which should not be in there
	if err == nil {
		for f, _ := range files {
			//			fmt.Printf("Empty file: %s\n", f)
			WriteEmptyFile(f)
		}
	}

	return err
}

func WriteEmptyFile(fname string) {
	s := fmt.Sprintf("%s\n", YAML_ID)
	ioutil.WriteFile(fname, []byte(s), 0666)
}
func isOurFile(fname string) bool {
	bs, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		return false
	}
	s := string(bs)
	sx := strings.SplitN(s, "\n", 2)
	if len(sx) < 1 {
		return false
	}
	fmt.Printf("ID: \"%s\" in %s\n", sx[0], fname)
	if sx[0] == YAML_ID {
		return true
	}
	return false
}

func targetName(name string) string {
	x := strings.Split(name, ".")
	return x[0]
}

func RewriteConfigFile() {
	dir := fmt.Sprintf("%s/%s", *targetsdir, "registry")
	var buffer bytes.Buffer

	if *templatefile != "" {
		bs, err := ioutil.ReadFile(*templatefile)
		if err != nil {
			fmt.Println(err)
			return
		}
		buffer.WriteString(string(bs))
	}

	dedup := make(map[string]bool)
	for _, t := range targets.AllTargets() {
		if dedup[t.Name] {
			continue
		}
		dedup[t.Name] = true
		job_cfg := get_config_for_service(t.Name)
		fname := fmt.Sprintf("%s/%s.yaml", dir, t.Name)

		buffer.WriteString(fmt.Sprintf("  - job_name: '%s'\n", t.Name))
		sl := *sample_limit
		if job_cfg != nil {
			sl = int(job_cfg.SampleLimit)
		}
		if sl != 0 {
			buffer.WriteString(fmt.Sprintf("    sample_limit: %d\n", sl))
		}
		buffer.WriteString(fmt.Sprintf("    metrics_path: '/internal/service-info/metrics'\n"))
		if t.httponly {
			buffer.WriteString(fmt.Sprintf("    scheme: 'http'\n"))
		} else {
			buffer.WriteString(fmt.Sprintf("    scheme: 'https'\n"))
			buffer.WriteString(fmt.Sprintf("    tls_config:\n"))
			buffer.WriteString(fmt.Sprintf("      insecure_skip_verify: true\n"))
		}
		buffer.WriteString(fmt.Sprintf("    file_sd_configs:\n"))
		buffer.WriteString(fmt.Sprintf("      - files:\n"))
		buffer.WriteString(fmt.Sprintf("        - '%s'\n", fname))
	}
	if *pmcfgfile == "" {
		return
	}
	err := ioutil.WriteFile(*pmcfgfile, []byte(buffer.String()), 0644)
	if err != nil {
		fmt.Printf("Failed to write config file: %s\n", err)
	}
}

func QueryForTargets(ctx context.Context, req *pb.Reporter) (*pb.TargetList, error) {
	rn := req.Reporter
	if !strings.Contains(rn, ":") {
		rn = rn + ":5001"
	}
	fmt.Printf("Querying \"%s\"\n", rn)
	res := &pb.TargetList{Reporter: &pb.Reporter{Reporter: rn}}
	con, err := client.ConnectWithIP(rn)
	if err != nil {
		return nil, err
	}
	defer con.Close()
	rg := reg.NewRegistryClient(con)
	rl, err := rg.ListRegistrations(ctx, &reg.V2ListRequest{})
	if err != nil {
		return nil, err
	}
	ts := make(map[string][]string)
	for _, r := range rl.Registrations {
		if !r.Targetable {
			continue
		}
		if !r.Running {
			continue
		}
		if !hasStatus(r.Target) {
			continue
		}
		sn := r.Target.ServiceName
		ts[sn] = append(ts[sn], fmt.Sprintf("%s:%d", r.Target.IP, r.Target.Port))
	}
	addressct := 0
	for sn, sl := range ts {
		addressct = addressct + len(sl)
		t := &pb.Target{Name: sn, Addresses: sl}
		res.Targets = append(res.Targets, t)
	}
	fmt.Printf("Queried, got %d services and %d addresses\n", len(ts), addressct)
	return res, nil
}

func hasStatus(r *reg.Target) bool {
	for _, at := range r.ApiType {
		if at == reg.Apitype_status {
			return true
		}
	}
	return false
}
