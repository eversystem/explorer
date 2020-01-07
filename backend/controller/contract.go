package controller

import (
	"errors"
	"net/http"

	"github.com/iost-official/explorer/backend/model/blockchain"
	"github.com/iost-official/explorer/backend/model/blockchain/rpcpb"
	"github.com/labstack/echo"
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

func GetContractDetail(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("ContractId required")
	}
	contractInfo, err := blockchain.GetContract(id, false)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, FormatResponse(struct {
		*rpcpb.Contract
	}{
		contractInfo,
	}))
}
