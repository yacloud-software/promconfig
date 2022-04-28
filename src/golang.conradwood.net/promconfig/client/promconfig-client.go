package main

import (
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"strings"
)

var (
	find = flag.String("find", "", "find a given series")
)

// this is a very simple client.
// it looks up the rpc endpoint for the logging server
// (which we assume is always running) and lists
// all the apps the logging server knows about
// the client is meant to illustrate a point rather than being
// terribly useful.
func main() {
	flag.Parse()
	if *find != "" {
		Find()
		os.Exit(0)
	}
	// get a Connection
	// (also accepts a path, such as "picoservices/authserver/1" )
	con := client.Connect("promconfig.PromConfigService")
	// once we got a connection, here's the 'client'
	echoClient := pb.NewPromConfigServiceClient(con)
	// a context with authentication
	ctx := authremote.Context()

	response, err := echoClient.QueryForTargets(ctx, &pb.Reporter{Reporter: cmdline.GetRegistryAddress()})
	utils.Bail("Failed to ping server", err)
	fmt.Printf("queried and found %d targets\n", len(response.Targets))
	for _, t := range response.Targets {
		fmt.Printf("%v\n", t)
	}
	fmt.Printf("Done.\n")
	os.Exit(0)
}
func Find() {
	sm := buildMatcher(*find)
	ctx := authremote.Context()
	sl, err := pb.GetPromConfigServiceClient().FindSeries(ctx, sm)
	utils.Bail("failed to get series", err)
	fmt.Printf("Got %d series\n", len(sl.Series))
	for _, s := range sl.Series {
		fmt.Printf("%s\n", s)
	}
}

func buildMatcher(s string) *pb.SeriesMatch {
	ms := strings.Split(s, ",")
	res := &pb.SeriesMatch{PartialMatch: true}
	for _, m := range ms {
		res.Prefix = append(res.Prefix, m)
	}
	return res
}
