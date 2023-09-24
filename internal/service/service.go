package service

import (
	"context"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSyncService)

type Service interface {
	Start(context.Context) error
	Stop(context.Context) error
}
