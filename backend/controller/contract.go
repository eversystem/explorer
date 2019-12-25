package controller

import (
	"errors"

	"github.com/iost-official/explorer/backend/model/blockchain"
	"github.com/iost-official/explorer/backend/model/blockchain/rpcpb"
)

/*
const (
	chainInfoUrl    = "http://" + HOST + ":30001/getChainInfo"
	chainStorageUrl = "http://" + HOST + ":30001/getContractStorage"
)
*/

var (
	errContracttNotExist = errors.New("contract not exists")
)

func init() {
}

func getContractDetail(id string) (*rpcpb.Contract, error) {
	return blockchain.GetContract(id, false)
}
