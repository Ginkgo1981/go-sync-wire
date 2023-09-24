package service

import (
	"context"
	"time"

	"github.com/ginkgo1981/nft-syncer/internal/biz"
	"github.com/ginkgo1981/nft-syncer/internal/data"
	"github.com/ginkgo1981/nft-syncer/internal/logger"
	ckbTypes "github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

type SyncService struct {
	checkInfoUsecase *biz.CheckInfoUsecase
	logger           *logger.Logger
	client           *data.CkbNodeClient
	status           chan struct{}
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
				s.process(ctx)
			}
		}
	}()
	return nil
}

func (s *SyncService) process(ctx context.Context) error {
	checkInfo, err := s.checkInfoUsecase.FindLastCheckInfo(ctx)
	if err != nil {
		s.logger.Errorf(ctx, "failed to find last check info: %v", err)
		return err
	}

	tipBlockNumber, err := s.client.Rpc.GetTipBlockNumber(ctx)
	if err != nil {
		s.logger.Errorf(ctx, "failed to get tip block number: %v", err)
		return err
	}

	if checkInfo.BlockNumber > tipBlockNumber {
		return nil
	}

	targetBlockNumber := checkInfo.BlockNumber
	if targetBlockNumber > tipBlockNumber {
		targetBlockNumber = tipBlockNumber
	}

	targetBlock, err := s.client.Rpc.GetBlockByNumber(ctx, targetBlockNumber)
	if err != nil {
		s.logger.Errorf(ctx, "failed to get block by number: %v", err)
		return err
	}

	checkInfo.BlockNumber = targetBlock.Header.Number + 1
	checkInfo.BlockHash = targetBlock.Header.Hash.String()[2:]

	err = s.SyncBlock(ctx, targetBlock, *checkInfo)
	if err != nil {
		s.logger.Errorf(ctx, "failed to sync block: %v", err)
		return err
	}

	return nil
}

func (c *SyncService) SyncBlock(ctx context.Context, block *ckbTypes.Block, checkInfo biz.CheckInfo) error {
	for index, tx := range block.Transactions {
		c.logger.Infof(ctx, "tx index: %d, tx: %s", index, tx.Hash)
	}

	return c.checkInfoUsecase.CreateCheckInfo(ctx, &checkInfo)
}

func (s *SyncService) Stop(ctx context.Context) error {
	s.logger.Info(ctx, "Successfully closed the sync service~")
	return nil
}

func NewSyncService(checkInfoUsecase *biz.CheckInfoUsecase, logger *logger.Logger, client *data.CkbNodeClient) *SyncService {
	return &SyncService{
		checkInfoUsecase: checkInfoUsecase,
		logger:           logger,
		client:           client,
	}
}
