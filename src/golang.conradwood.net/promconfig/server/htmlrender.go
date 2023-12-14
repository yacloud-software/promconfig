package main

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/htmlserver"
)

func (e *promConfigServer) HTMLRenderer(ctx context.Context, req *htmlserver.SnippetRequest) (*htmlserver.SnippetResponse, error) {
	fmt.Printf("Serving \"%s\"\n", req.Path)
	sr := &htmlserver.SnippetResponse{
		Body:        []byte("foobody"),
		DoNotModify: true,
	}
	return sr, nil
}






