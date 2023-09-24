package service

import (
	"context"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSyncService, NewCheckInfoCleanerService)

type Service interface {
	Start(context.Context) error
	Stop(context.Context) error
	process(context.Context) error
}
