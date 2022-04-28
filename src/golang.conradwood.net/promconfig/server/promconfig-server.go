package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/server"
	//	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/promconfig/targets"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
)

var (
	port  = flag.Int("port", 10000, "The grpc server port")
	debug = flag.Bool("debug", false, "debug mode")
)

type promConfigServer struct {
}

func main() {
	flag.Parse()
	fmt.Printf("Starting Promconfig...\n")
	sd := server.NewServerDef()
	sd.Port = *port
	sd.Register = server.Register(
		func(server *grpc.Server) error {
			e := new(promConfigServer)
			pb.RegisterPromConfigServiceServer(server, e)
			return nil
		},
	)
	err := server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/
func (e *promConfigServer) QueryForTargets(ctx context.Context, req *pb.Reporter) (*pb.TargetList, error) {
	tl, err := targets.QueryForTargets(ctx, req)
	return tl, err
}
func (e *promConfigServer) NewTargets(ctx context.Context, req *pb.TargetList) (*common.Void, error) {
	err := targets.UpdateTargets(req)
	if err != nil {
		if *debug {
			fmt.Printf("Failed: %s\n", err)
		}
		return nil, err
	}
	resp := &common.Void{}
	return resp, nil
}
