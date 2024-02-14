package targets

import (
	pb "golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/utils"
)

var (
	was_read = false
	cfg      = &pb.Config{}
)

func cfg_start() {
	if was_read {
		return
	}
	err := utils.ReadYaml("configs/targetconfig.yaml", cfg)
	utils.Bail("failed to read config", err)
	was_read = true
}
func get_config_for_service(service string) *pb.TargetConfig {
	cfg_start()
	for _, t := range cfg.Configs {
		if t.ServiceName == service {
			return t
		}
	}
	return nil
}
