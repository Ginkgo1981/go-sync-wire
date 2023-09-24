package biz

import (
	"context"

	"github.com/ginkgo1981/nft-syncer/internal/logger"
)

type CheckType uint8

const (
	SyncEvent CheckType = iota // SyncEvent = 0
)

func (t CheckType) String() string {
	return []string{"sync_event"}[t]
}

type CheckInfo struct {
	Id          uint64
	BlockNumber uint64
	BlockHash   string
	CheckType   CheckType
}

type CheckInfoRepo interface {
	FindOrCreateCheckInfo(ctx context.Context, info *CheckInfo) error
	UpdateCheckInfo(ctx context.Context, info CheckInfo) error
	FindLastCheckInfo(ctx context.Context) (*CheckInfo, error)
	CreateCheckInfo(context.Context, *CheckInfo) error
	CleanCheckInfo(context.Context, CheckType) error
}

type CheckInfoUsecase struct {
	repo   CheckInfoRepo
	logger *logger.Logger
}

func NewCheckInfoUsecase(repo CheckInfoRepo, logger *logger.Logger) *CheckInfoUsecase {
	return &CheckInfoUsecase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *CheckInfoUsecase) FindOrCreate(ctx context.Context, checkInfo *CheckInfo) error {
	return uc.repo.FindOrCreateCheckInfo(ctx, checkInfo)
}

func (uc *CheckInfoUsecase) Update(ctx context.Context, checkInfo *CheckInfo) error {
	return uc.Update(ctx, checkInfo)
}

func (uc *CheckInfoUsecase) FindLastCheckInfo(ctx context.Context) (*CheckInfo, error) {
	return uc.repo.FindLastCheckInfo(ctx)
}

func (uc *CheckInfoUsecase) CreateCheckInfo(ctx context.Context, checkInfo *CheckInfo) error {
	return uc.repo.CreateCheckInfo(ctx, checkInfo)
}

func (uc *CheckInfoUsecase) CleanCheckInfo(ctx context.Context, checkType CheckType) error {
	return uc.repo.CleanCheckInfo(ctx, checkType)
}
