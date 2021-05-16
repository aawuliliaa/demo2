// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"project/internal/dao"
	"project/internal/server/grpc"
	"project/internal/server/http"
	"project/internal/service"
)

// Injectors from wire.go:

//go:generate project t wire
func InitApp() (*App, func(), error) {
	daoDao, cleanup, err := dao.NewDB()
	if err != nil {
		return nil, nil, err
	}
	groupService := service.NewGroupService(daoDao)
	namespaceService := service.NewNamespaceService(daoDao)
	serviceService := service.NewService(groupService, namespaceService)
	server, err := http.NewHttpServer(serviceService)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	grpcServer, err := grpc.NewGrpcServer()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	app, cleanup2, err := NewApp(server, grpcServer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
