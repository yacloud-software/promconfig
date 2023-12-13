package main

import (
	"context"
	"encoding/json"
	"flag"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/http"
	"strings"
)

var (
	api_url = flag.String("prometheus_api_url", "http://prometheus:9090/prometheus/api/v1/", "url of the prometheus server")
)

/*
curl 'http://prometheus:9090/prometheus/api/v1/label/__name__/values'
*/
func (p *promConfigServer) FindSeries(ctx context.Context, req *pb.SeriesMatch) (*pb.SeriesList, error) {
	sl, err := getAllSeries()
	if err != nil {
		return nil, err
	}
	res := &pb.SeriesList{}
	for _, s := range sl.Series {
		if seriesMatches(req, s) {
			res.Series = append(res.Series, s)
		}
	}
	return res, nil
}
func (p *promConfigServer) GetSeries(ctx context.Context, req *common.Void) (*pb.SeriesList, error) {
	res, err := getAllSeries()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// does a series match a given matcher?
func seriesMatches(req *pb.SeriesMatch, s *pb.Series) bool {
	t := strings.ToLower(s.Name)
	for _, p := range req.Prefix {
		p = strings.ToLower(p)
		if req.PartialMatch {
			if strings.Contains(t, p) {
				return true
			}
		} else {
			if strings.HasPrefix(t, p) {
				return true
			}
		}
	}
	return false
}

type PromJsonSeries struct {
	Status   string   `json:"status"`
	Data     []string `json:"data"`
	Warnings []string `json:"warnings"`
}

func getAllSeries() (*pb.SeriesList, error) {
	ht := &http.HTTP{}
	hr := ht.Get(*api_url + "label/__name__/values")
	err := hr.Error()
	if err != nil {
		return nil, err
	}
	pjs := &PromJsonSeries{}
	err = json.Unmarshal(hr.Body(), pjs)
	if err != nil {
		return nil, err
	}
	res := &pb.SeriesList{}
	for _, d := range pjs.Data {
		s := &pb.Series{Name: d}
		res.Series = append(res.Series, s)
	}
	return res, nil
}





