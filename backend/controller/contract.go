package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"os"

	"github.com/iost-official/explorer/backend/model/blockchain"

	"github.com/iost-official/explorer/backend/model/db"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

/*
const (
	chainInfoUrl    = "http://" + HOST + ":30001/getChainInfo"
	chainStorageUrl = "http://" + HOST + ":30001/getContractStorage"
)
*/

var (
	bpClient           *http.Client
	errContracttNotExist = errors.New("contract not exists")
)

func init() {
	bpClient = &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			MaxIdleConns: 5,
		},
	}
}

func getContractDetail(id string) (*rpcpb.Contract, error) {
	return blockchain.GetContract(id, false)
}
