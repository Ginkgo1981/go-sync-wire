package data

import (
	"context"
	"time"

	"github.com/ginkgo1981/nft-syncer/internal/biz"
	"github.com/ginkgo1981/nft-syncer/internal/logger"
)

var _ biz.CheckInfoRepo = (*checkInfoRepo)(nil)

type CheckInfo struct {
	// gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	BlockNumber uint64
	BlockHash   string
	CheckType   biz.CheckType
}

type checkInfoRepo struct {
	data   *Data
	logger *logger.Logger
}

// CleanCheckInfo implements biz.CheckInfoRepo
func (rp *checkInfoRepo) CleanCheckInfo(context.Context, biz.CheckType) error {
	var checkInfos []CheckInfo

	if err := rp.data.db.Debug().Where("check_type = ?", 0).Order("id").Limit(100).Find(&checkInfos).Error; err != nil {
		return err
	}

	if len(checkInfos) < 10 {
		return nil
	}

	lastInfo := checkInfos[len(checkInfos)-1]

	if err := rp.data.db.Debug().Where("id < ?", lastInfo.ID).Order("id").Limit(10).Delete(&CheckInfo{}).Error; err != nil {
		return err
	}

	return nil
}

// CreateCheckInfo implements biz.CheckInfoRepo
func (rp *checkInfoRepo) CreateCheckInfo(ctx context.Context, info *biz.CheckInfo) error {
	if err := rp.data.db.Debug().WithContext(ctx).Create(&CheckInfo{
		BlockNumber: info.BlockNumber,
		BlockHash:   info.BlockHash,
		CheckType:   info.CheckType,
	}).Error; err != nil {
		return err
	}

	return nil
}

// FindLastCheckInfo implements biz.CheckInfoRepo
func (rp *checkInfoRepo) FindLastCheckInfo(ctx context.Context) (*biz.CheckInfo, error) {
	info := &CheckInfo{}
	if err := rp.data.db.Debug().WithContext(ctx).Where("check_type = ?", 0).Order("block_number desc").Find(&info).Error; err != nil {
		return nil, err
	}

	return &biz.CheckInfo{
		Id:          uint64(info.ID),
		BlockNumber: info.BlockNumber,
		BlockHash:   info.BlockHash,
	}, nil
}

func NewCheckInfoRepo(data *Data, logger *logger.Logger) biz.CheckInfoRepo {
	return &checkInfoRepo{
		data:   data,
		logger: logger,
	}
}

func (rp checkInfoRepo) FindOrCreateCheckInfo(ctx context.Context, info *biz.CheckInfo) error {
	if err := rp.data.db.WithContext(ctx).FirstOrCreate(info, CheckInfo{BlockNumber: info.BlockNumber, CheckType: biz.CheckType(0)}).Error; err != nil {
		return err
	}
	return nil
}

func (rp checkInfoRepo) UpdateCheckInfo(ctx context.Context, info biz.CheckInfo) error {
	if err := rp.data.db.WithContext(ctx).Save(info).Error; err != nil {
		return err
	}
	return nil
}
