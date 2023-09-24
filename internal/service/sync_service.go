package service

import (
	"context"
	"time"

	"github.com/ginkgo1981/nft-syncer/internal/data"
	"github.com/ginkgo1981/nft-syncer/internal/logger"
)

type SyncService struct {
	logger *logger.Logger
	client *data.CkbNodeClient
	status chan struct{}
}

func (s *SyncService) Start(ctx context.Context) error {
	s.logger.Info(ctx, "Successfully started the sync service~")

	var interval time.Duration
	interval = time.Second

	go func() {
		for {
			select {
			case <-ctx.Done():
				s.status <- struct{}{}
				s.logger.Infof(ctx, "receive cancel signal %v", ctx.Err())

				return
			case <-time.After(interval):
				s.sync(ctx)
			}
		}
	}()
	return nil
}

func (s *SyncService) sync(ctx context.Context) error {
	tipBlockNumber, err := s.client.Rpc.GetTipBlockNumber(ctx)
	if err != nil {
		s.logger.Errorf(ctx, "failed to get tip block number: %v", err)
		return err
	}

	s.logger.Infof(ctx, "current block number: %v", tipBlockNumber)
	return nil
}

func (s *SyncService) Stop(ctx context.Context) error {
	s.logger.Info(ctx, "Successfully closed the sync service~")
	return nil
}

func NewSyncService(logger *logger.Logger, client *data.CkbNodeClient) *SyncService {
	return &SyncService{
		logger: logger,
		client: client,
	}
}
