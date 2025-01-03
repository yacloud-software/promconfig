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
	"os"
	"strings"
	"time"

	"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/promconfig/targets"
	"google.golang.org/grpc"
)

var (
	port         = flag.Int("port", 10000, "The grpc server port")
	debug        = flag.Bool("debug", false, "debug mode")
	registries   = flag.String("registries", "", "if set, query these registries regularly and on startup")
	requery_chan = make(chan bool, 2)
)

type promConfigServer struct {
}

func main() {
	flag.Parse()
	fmt.Printf("Starting Promconfig...\n")
	server.SetHealth(common.Health_READY)
	var err error

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
	err = server.ServerStartup(sd)
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
	// update our server list
	for _, target := range tl.Targets {
		for _, adr := range target.Addresses {
			AddServer(adr)
		}
	}

	err = targets.UpdateTargets(tl)
	if err != nil {
		return nil, err
	}
	return tl, nil
}
func (e *promConfigServer) NewTargets(ctx context.Context, req *pb.TargetList) (*common.Void, error) {
	// update our server list
	for _, target := range req.Targets {
		for _, adr := range target.Addresses {
			AddServer(adr)
		}
	}

	// consider the targets themselves
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
func (e promConfigServer) Requery(ctx context.Context, req *common.Void) (*common.Void, error) {
	requery_chan <- true
	return req, nil
}
func reg_query_loop() {
	t := time.Duration(3) * time.Second
	for {
		select {
		case <-time.After(t):
			//
		case <-requery_chan:
			//
		}

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
