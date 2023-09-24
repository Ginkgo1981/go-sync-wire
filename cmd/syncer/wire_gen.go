// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ginkgo1981/nft-syncer/internal/app"
	"github.com/ginkgo1981/nft-syncer/internal/config"
	"github.com/ginkgo1981/nft-syncer/internal/data"
	"github.com/ginkgo1981/nft-syncer/internal/logger"
	"github.com/ginkgo1981/nft-syncer/internal/service"
)

// Injectors from wire.go:

func initApp(database *config.Database, ckbNode *config.CkbNode, loggerLogger *logger.Logger) (*app.App, func(), error) {
	ckbNodeClient, err := data.NewCkbNodeClient(ckbNode, loggerLogger)
	if err != nil {
		return nil, nil, err
	}
	syncService := service.NewSyncService(loggerLogger, ckbNodeClient)
	dataData, cleanup, err := data.NewData(database, loggerLogger)
	if err != nil {
		return nil, nil, err
	}
	dbMigration := data.NewDBMigration(dataData, loggerLogger)
	appApp := newApp(loggerLogger, syncService, dbMigration)
	return appApp, func() {
		cleanup()
	}, nil
}
