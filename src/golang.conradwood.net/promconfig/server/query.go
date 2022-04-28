package main

import (
	"context"
	pb "golang.conradwood.net/apis/promconfig"
)

/*
curl 'http://prometheus:9090/prometheus/api/v1/label/__name__/values'
*/
func (p *promConfigServer) GetSeries(ctx context.Context, req *pb.SeriesMatch) (*pb.SeriesList, error) {
	res := &pb.SeriesList{}
	return res, nil
}
