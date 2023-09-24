package data

import (
	"context"

	"github.com/ginkgo1981/nft-syncer/internal/config"
	"github.com/ginkgo1981/nft-syncer/internal/logger"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
)

type CkbNodeClient struct {
	Rpc rpc.Client
}

func NewCkbNodeClient(conf *config.CkbNode, logger *logger.Logger) (*CkbNodeClient, error) {
	client, err := rpc.Dial(conf.RpcUrl)
	if err != nil {
		logger.Errorf(context.TODO(), "failed to connect to the ckb node")
		return nil, err
	}
	return &CkbNodeClient{
		Rpc: client,
	}, nil
}
