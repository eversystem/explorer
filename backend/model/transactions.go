package model

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gogo/protobuf/proto"
	"github.com/iost-official/explorer/backend/model/blockchain/rpcpb"
	"github.com/iost-official/explorer/backend/model/db"
	contract "github.com/iost-official/explorer/backend/model/pb"
	"github.com/iost-official/explorer/backend/util"
)

type Action struct {
	Contract string `json:"contract,omitempty"`
	// action name
	ActionName string `json:"action_name,omitempty"`
	// data
	Data       string   `json:"data,omitempty"`
}
/*
type Signer struct {
	// signature algorithm
	Algorithm int32 `json:"algorithm,omitempty"`
	// signature bytes
	Signature []byte `json:"signature,omitempty"`
	// public key
	PublicKey   []byte   `json=publicKey,proto3" json:"public_key,omitempty"`
}
*/
// transaction receipts
type Receipt struct {
	// function name
	FuncName string `json:"func_name,omitempty"`
	// content
	Content  string   `json:"content,omitempty"`
}

type TxReceipt struct {
	// transaction hash
	TxHash string `json:"tx_hash,omitempty"`
	// gas usage
	GasUsage float64 `json:"gas_usage,omitempty"`
	// ram usage
	RamUsage map[string]int64 `json:"ram_usage,omitempty"`
	// status code
	StatusCode int32 `json:"status_code,omitempty"`
	// message
	Message string `json:"message,omitempty"`
	// transaction returns
	Returns []string `json:"returns,omitempty"`
	// transaction receipts
	Receipts []Receipt `json:"receipts,omitempty"`
}

type TxAmountLimit struct {
	// token name
	Token string `json:"token,omitempty"`
	// limit value
	Value string `json:"value,omitempty"`
}

/// this struct is used as json to return
type TxnDetail struct {
	Hash          string  `json:"txHash"`
	BlockNumber   int64   `json:"blockHeight"`
	From          string  `json:"from"`
	To            string  `json:"to"`
	Amount        float64 `json:"amount"`
	GasLimit      float64 `json:"gasLimit"`
	GasPrice      float64 `json:"price"`
	Age           string  `json:"age"`
	UTCTime       string  `json:"utcTime"`
	Code          string  `json:"code"`
	StatusCode    int32   `json:"statusCode"`
	StatusMessage string  `json:"statusMessage"`
	Contract      string  `json:"contract"`
	ActionName    string  `json:"actionName"`
	Data          string  `json:"data"`
	Memo          string  `json:"memo"`
	Expiration    int64   `json:"expiration"`
	Actions       []Action `json:"actions"`
	Signers       []string `json:"signers"`
	Publisher     string   `json:"publisher,omitempty"`
	ReferredTx    string   `json:"referredTx"`
	AmountLimit   []TxAmountLimit `json:"amountLimit"`
	Receipt       TxReceipt `json:"receipt"`
}

type TxJson struct {
	Hash        string  `json:"hash"`
	BlockNumber int64   `json:"blockNumber"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	Amount      float64 `json:"amount"`
	GasLimit    float64 `json:"gasLimit"`
	GasPrice    float64 `json:"gasPrice"`
	Age         string  `json:"age"`
	UTCTime     string  `json:"utcTime"`
}

func ConvertTxJsons(txs []*db.TxStore) []*TxJson {
	var ret []*TxJson
	for _, tx := range txs {
		txnOut := &TxJson{
			Hash:        tx.Tx.Hash,
			BlockNumber: tx.BlockNumber,
			From:        tx.Tx.Publisher,
			To:          tx.Tx.Actions[0].Contract,
			GasLimit:    tx.Tx.GasLimit,
			GasPrice:    tx.Tx.GasRatio,
			Age:         util.ModifyIntToTimeStr(tx.Tx.Time / 1e9),
			UTCTime:     util.FormatUTCTime(tx.Tx.Time),
		}

		if tx.Tx.Actions[0].Contract == "token.iost" && tx.Tx.Actions[0].ActionName == "transfer" &&
			tx.Tx.TxReceipt.StatusCode == rpcpb.TxReceipt_SUCCESS {
			var params []string
			err := json.Unmarshal([]byte(tx.Tx.Actions[0].Data), &params)
			if err == nil && len(params) == 5 && params[0] == "iost" {
				txnOut.From = params[1]
				txnOut.To = params[2]
				f := getIOSTAmount(params[3])
				txnOut.Amount = f
			}
		} else if tx.Tx.Actions[0].Contract == "exchange.iost" && tx.Tx.Actions[0].ActionName == "transfer" &&
			tx.Tx.TxReceipt.StatusCode == rpcpb.TxReceipt_SUCCESS {
			var params []string
			err := json.Unmarshal([]byte(tx.Tx.Actions[0].Data), &params)
			if err == nil && len(params) == 4 && params[0] == "iost" {
				if params[1] != "" {
					txnOut.From = tx.Tx.Publisher
					txnOut.To = params[1]
					f := getIOSTAmount(params[2])
					txnOut.Amount = f
				}
			}
		}
		ret = append(ret, txnOut)
	}
	return ret
}

func ConvertTxsOutputs(tx []*db.TxStore) []*TxnDetail {
	var ret []*TxnDetail
	for _, t := range tx {
		ret = append(ret, ConvertTxOutput(t))
	}
	return ret
}

func getIOSTAmount(s string) float64 {
	point := strings.Index(s, ".")
	if point > 0 && point+9 <= len(s) {
		s = s[:point+9]
	}
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

/// convert FlatTx to TxnDetail
/*
	Status
	0: "SUCCESS",
	1: "GAS_RUN_OUT",
	2: "BALANCE_NOT_ENOUGH",
	3: "WRONG_PARAMETER",
	4: "RUNTIME_ERROR",
	5: "TIMEOUT",
	6: "WRONG_TX_FORMAT",
	7: "DUPLICATE_SET_CODE",
	8: "UNKNOWN_ERROR",

	Transaction Response Status
	0: "PENDING",
	1: "PACKED",
	2: "IRREVERSIBLE",

	Block Response Status
	0: "PENDING",
	1: "IRREVERSIBLE",

	EventTopic
	0: "CONTRACT_RECEIPT",
	1: "CONTRACT_EVENT",

*/
func ConvertTxOutput(tx *db.TxStore) *TxnDetail {
	var actions []Action
	for i := 0; i < len(tx.Tx.Actions); i++ {
		action := Action{
			Contract: tx.Tx.Actions[i].Contract,
			// action name
			ActionName: tx.Tx.Actions[i].ActionName,
			// data
			Data: tx.Tx.Actions[i].Data,
		}
		actions = append(actions, action)
	}
	/*
	var signers []Signer
	for i := 0; i < len(tx.Tx.Signers); i++ {
		signer := Signer{
			// signature algorithm
			Algorithm: tx.Tx.Signers[i].Algorithm,
			// signature bytes
			Signature: tx.Tx.Signers[i].Signature,
			// public key
			PublicKey: tx.Tx.Signers[i].PublicKey,
		}
		signers = append(signers, signer);
	}
	*/
	var receipts []Receipt
	for i := 0; i < len(tx.Tx.TxReceipt.Receipts); i++ {
		receipt := Receipt {
			// function name
			FuncName: tx.Tx.TxReceipt.Receipts[i].FuncName,
			// content
			Content: tx.Tx.TxReceipt.Receipts[i].Content,
		}
		receipts = append(receipts, receipt)
	}
	var amountLimits []TxAmountLimit
	for i := 0; i < len(tx.Tx.AmountLimit); i++ {
		receipt := TxAmountLimit {
			// function name
			Token: tx.Tx.AmountLimit[i].Token,
			// content
			Value: tx.Tx.AmountLimit[i].Value,
		}
		amountLimits = append(amountLimits, receipt)
	}
	txRecept := TxReceipt{
		// transaction hash
		TxHash: tx.Tx.TxReceipt.TxHash,
		// gas usage
		GasUsage: tx.Tx.TxReceipt.GasUsage,
		// ram usage
		RamUsage: tx.Tx.TxReceipt.RamUsage,
		// status code
		StatusCode: int32(tx.Tx.TxReceipt.StatusCode),
		// message
		Message: tx.Tx.TxReceipt.Message,
		// transaction returns
		Returns: tx.Tx.TxReceipt.Returns,
		// transaction receipts
		Receipts: receipts,
	}
	txnOut := &TxnDetail{
		Hash:          tx.Tx.Hash,
		BlockNumber:   tx.BlockNumber,
		From:          tx.Tx.Publisher,
		To:            tx.Tx.Actions[0].Contract,
		GasLimit:      tx.Tx.GasLimit,
		GasPrice:      tx.Tx.GasRatio,
		Age:           util.ModifyIntToTimeStr(tx.Tx.Time / 1e9),
		UTCTime:       util.FormatUTCTime(tx.Tx.Time),
		Code:          "",
		StatusCode:    int32(tx.Tx.TxReceipt.StatusCode),
		StatusMessage: tx.Tx.TxReceipt.Message,
		Contract:      tx.Tx.Actions[0].Contract,
		ActionName:    tx.Tx.Actions[0].ActionName,
		Data:          tx.Tx.Actions[0].Data,
		Expiration:    tx.Tx.Expiration,
		Actions:       actions,
		Signers:       tx.Tx.Signers,
		Publisher:     tx.Tx.Publisher,
		ReferredTx:    tx.Tx.ReferredTx,
		Receipt:       txRecept,
	}

	if tx.Tx.Actions[0].Contract == "token.iost" && tx.Tx.Actions[0].ActionName == "transfer" &&
		tx.Tx.TxReceipt.StatusCode == rpcpb.TxReceipt_SUCCESS {
		var params []string
		err := json.Unmarshal([]byte(tx.Tx.Actions[0].Data), &params)
		if err == nil && len(params) == 5 && params[0] == "iost" {
			txnOut.From = params[1]
			txnOut.To = params[2]
			f := getIOSTAmount(params[3])
			txnOut.Amount = f * 1e8
			txnOut.Memo = params[4]
		}
	} else if tx.Tx.Actions[0].Contract == "exchange.iost" && tx.Tx.Actions[0].ActionName == "transfer" &&
		tx.Tx.TxReceipt.StatusCode == rpcpb.TxReceipt_SUCCESS {
		var params []string
		err := json.Unmarshal([]byte(tx.Tx.Actions[0].Data), &params)
		if err == nil && len(params) == 4 && params[0] == "iost" {
			if params[1] != "" {
				txnOut.From = tx.Tx.Publisher
				txnOut.To = params[1]
				f := getIOSTAmount(params[2])
				txnOut.Amount = f * 1e8
			}
		}
	}

	if tx.Tx.Actions[0].Contract == "system.iost" && tx.Tx.Actions[0].ActionName == "setCode" {
		var params []string
		err := json.Unmarshal([]byte(tx.Tx.Actions[0].Data), &params)
		if err == nil && len(params) > 0 {
			j, e := simplejson.NewJson([]byte(params[0]))
			if e != nil {
				log.Printf("json decode setCode param failed. err=%v", err)
			} else {
				txnOut.Code, _ = j.Get("code").String()
			}

			if txnOut.Code == "" {
				bytes, e := base64.StdEncoding.DecodeString(params[0])
				if e != nil {
					log.Printf("base64 decode setCode param failed. err=%v", err)
				} else {
					var con contract.Contract
					proto.Unmarshal(bytes, &con)
					txnOut.Code = con.Code
				}
			}
		}

		// c.B64Decode(code[0])
		/* txnOut.Code = c.Code */
	}

	return txnOut
}

func GetDetailTxn(txHash string) (*TxnDetail, error) {
	tx, err := db.GetTxByHash(txHash)

	if err != nil {
		log.Printf("transaction not found. txHash:%v, err=%v", txHash, err)
		return nil, err
	}

	txnOut := ConvertTxOutput(tx)
	txnOut.Amount /= 1e8

	return txnOut, nil
}

/// get a list of transactions for a specific page using account and block
func GetFlatTxnSlicePage(page, eachPageNum, block int64) ([]*TxnDetail, error) {
	lastPageNum, err := db.GetTxTotalPageCnt(eachPageNum, block)
	if err != nil {
		return nil, err
	}

	if lastPageNum == 0 {
		return []*TxnDetail{}, nil
	}

	if page > lastPageNum {
		page = lastPageNum
	}

	start := int((page - 1) * eachPageNum)
	txnsFlat, err := db.GetTxs(block, start, int(eachPageNum))

	if err != nil {
		return nil, err
	}

	return ConvertTxsOutputs(txnsFlat), nil
}
