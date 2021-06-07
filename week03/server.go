package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type Server struct {
	Name string
	Addr string
}

func NewServer(addr string, name string) *Server {
	return &Server{
		Addr: addr,
		Name: name,
	}
}

func (s *Server) Serve(ctx context.Context, group *errgroup.Group) error {
	srv := &http.Server{
		Addr: s.Addr,
	}

	group.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println(s.Name, "exit ...")
			return srv.Shutdown(context.TODO())
		}
	})

	return srv.ListenAndServe()
}
