package service

import (
	"context"
	"time"

	"github.com/ginkgo1981/nft-syncer/internal/biz"
	"github.com/ginkgo1981/nft-syncer/internal/logger"
)

var _ Service = (*CheckInfoCleanerService)(nil)

// to clean the check info
type CheckInfoCleanerService struct {
	checkInfoUsecase *biz.CheckInfoUsecase
	logger           *logger.Logger
	status           chan struct{}
}

// Start implements Service
func (c *CheckInfoCleanerService) Start(ctx context.Context) error {
	c.logger.Info(ctx, "Successfully started the check info cleaner service~")

	var interval time.Duration
	interval = time.Second

	go func() {
		for {
			select {
			case <-ctx.Done():
				c.status <- struct{}{}
				c.logger.Infof(ctx, "receive cancel signal %v", ctx.Err())

				return
			case <-time.After(interval):
				c.process(ctx)
			}
		}
	}()
	return nil
}

func (c *CheckInfoCleanerService) process(ctx context.Context) error {
	return c.checkInfoUsecase.CleanCheckInfo(ctx, 0)
}

// Stop implements Service
func (c *CheckInfoCleanerService) Stop(ctx context.Context) error {
	c.logger.Info(ctx, "Successfully stopped the check info cleaner service~")
	return nil
}

// newCheckInfoCleaner creates a new CheckInfoCleanerService
func NewCheckInfoCleanerService(checkInfoUsecase *biz.CheckInfoUsecase, logger *logger.Logger) *CheckInfoCleanerService {
	return &CheckInfoCleanerService{
		checkInfoUsecase: checkInfoUsecase,
		logger:           logger,
	}
}
