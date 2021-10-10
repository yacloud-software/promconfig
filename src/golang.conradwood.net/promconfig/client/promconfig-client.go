package main

import (
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"os"
)

// this is a very simple client.
// it looks up the rpc endpoint for the logging server
// (which we assume is always running) and lists
// all the apps the logging server knows about
// the client is meant to illustrate a point rather than being
// terribly useful.
func main() {
	flag.Parse()

	// get a Connection
	// (also accepts a path, such as "picoservices/authserver/1" )
	con := client.Connect("promconfig.PromConfigService")
	// once we got a connection, here's the 'client'
	echoClient := pb.NewPromConfigServiceClient(con)
	// a context with authentication
	ctx := tokens.ContextWithToken()

	empty := pb.TargetList{}
	response, err := echoClient.NewTargets(ctx, &empty)
	utils.Bail("Failed to ping server", err)
	fmt.Printf("Response to ping: %v\n", response)

	fmt.Printf("Done.\n")
	os.Exit(0)
}
