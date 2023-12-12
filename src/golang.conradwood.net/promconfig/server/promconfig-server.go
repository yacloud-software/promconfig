package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/server"
	//	"golang.conradwood.net/go-easyops/utils"
	"context"
	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/promconfig/targets"
	"google.golang.org/grpc"
	"os"
	"strings"
	"time"
)

var (
	port       = flag.Int("port", 10000, "The grpc server port")
	debug      = flag.Bool("debug", false, "debug mode")
	registries = flag.String("registries", "", "if set, query these registries regularly and on startup")
)

type promConfigServer struct {
}

func main() {
	flag.Parse()
	fmt.Printf("Starting Promconfig...\n")
	go reg_query_loop()
	sd := server.NewServerDef()
	sd.SetPort(*port)
	sd.SetRegister(server.Register(
		func(server *grpc.Server) error {
			e := new(promConfigServer)
			pb.RegisterPromConfigServiceServer(server, e)
			return nil
		},
	))
	err := server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/
func (e *promConfigServer) QueryForTargets(ctx context.Context, req *pb.Reporter) (*pb.TargetList, error) {
	tl, err := targets.QueryForTargets(ctx, req)
	if err != nil {
		return nil, err
	}
	err = targets.UpdateTargets(tl)
	if err != nil {
		return nil, err
	}
	return tl, nil
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

func reg_query_loop() {
	t := time.Duration(3) * time.Second
	for {
		time.Sleep(t)
		if *registries != "" {
			pcs := &promConfigServer{}
			for _, r := range strings.Split(*registries, ",") {
				ctx := authremote.Context()
				if *debug {
					fmt.Printf("Reporter: %s\n", r)
				}
				req := &pb.Reporter{Reporter: r}
				_, err := pcs.QueryForTargets(ctx, req)
				if err != nil {
					fmt.Printf("failed to query registry \"%s\": %s\n", req.Reporter, err)
				}
			}
		}
		t = time.Duration(15) * time.Minute
	}
}




