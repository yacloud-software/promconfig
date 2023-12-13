package main

import (
	"context"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/promconfig"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/promconfig/db"
)

func (e *promConfigServer) UpdatePercentageAlert(ctx context.Context, req *pb.PercentAlert) (*pb.PercentAlert, error) {
	if req.TotalMetric == "" {
		return nil, errors.InvalidArgs(ctx, "missing TotalMetric", "missing TotalMetric")
	}
	if req.CountMetric == "" {
		return nil, errors.InvalidArgs(ctx, "missing CountMetric", "missing CountMetric")
	}
	if req.ID == 0 {
		_, err := db.DefaultDBPercentAlert().Save(ctx, req)
		if err != nil {
			return nil, err
		}
	} else {
		err := db.DefaultDBPercentAlert().Update(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return req, nil
}
func (e *promConfigServer) GetAllPercentageAlerts(ctx context.Context, req *common.Void) (*pb.PercentAlertList, error) {
	ls, err := db.DefaultDBPercentAlert().All(ctx)
	if err != nil {
		return nil, err
	}
	res := &pb.PercentAlertList{Alerts: ls}
	return res, nil
}





