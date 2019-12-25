package db

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	MongoLink     string
	MongoUser     = ""
	MongoPassWord = ""
	Db            string
)

const (
	CollectionBlocks                    = "blocks"
	CollectionBP                        = "blockproducer"
	CollectionTxs                       = "txs"
	CollectionFlatTx                    = "flatxs"
	CollectionAccount                   = "accounts"
	CollectionAccountTx                 = "accountTx"
	CollectionAccountPubkey             = "accountPubkey"
	CollectionContract                  = "contracts"
	CollectionContractTx                = "contractTx"
	CollectionTaskCursor                = "taskCursors"
	CollectionBlockPay                  = "blockPays"
	CollectionApplyIOST                 = "applyIOST"
	CollectionVoteTx                    = "voteTx"
	CollectionProducerAward             = "producerAward"
	CollectionUserAward                 = "userAward"
	CollectionProducerContributionAward = "producerContributionAward"
	CollectionUserContributionAward     = "userContributionAward"
	CollectionFailedAward               = "failedUserAward"
	CollectionAwardInfo                 = "awardInfo"
	CollectionProducerLevelInfo         = "producerLevel"
)

func InitConfig() {
	dbConfig := viper.GetStringMapString("mongodb")
	Db = dbConfig["db"]
	MongoUser = dbConfig["user"]
	MongoPassWord = dbConfig["password"]
	//MongoLink = fmt.Sprintf("%s:%s", dbConfig["host"], dbConfig["port"])
	MongoLink = dbConfig["mongolink"]
	fmt.Println("mongolink", Db, MongoLink)
}
