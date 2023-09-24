//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ginkgo1981/nft-syncer/internal/app"
	"github.com/ginkgo1981/nft-syncer/internal/config"
	"github.com/ginkgo1981/nft-syncer/internal/data"
	"github.com/ginkgo1981/nft-syncer/internal/logger"
	service "github.com/ginkgo1981/nft-syncer/internal/service"
	"github.com/google/wire"
)

func initApp(*config.Database, *config.CkbNode, *logger.Logger) (*app.App, func(), error) {
	panic(wire.Build(data.ProviderSet, service.ProviderSet, newApp))
}
