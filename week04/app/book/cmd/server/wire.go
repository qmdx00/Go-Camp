// +build wireinject

package main

import (
	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
	"week04/app/book/internal/conf"
	"week04/app/book/internal/server"
)

type App struct {
	HttpServer *server.HttpServer
	GRPCServer *server.GRPCServer
}

// newApp return App struct with server.HttpServer and server.GRPCServer
func newApp(http *server.HttpServer, grpc *server.GRPCServer) *App {
	return &App{HttpServer: http, GRPCServer: grpc}
}

// initApp Inject wire ProvideSet
func initApp(group *errgroup.Group, option conf.Options) *App {
	panic(wire.Build(server.ProvideSet, newApp))
}
