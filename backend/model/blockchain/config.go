package blockchain

import (
	"fmt"
	"os"
	"github.com/iost-official/explorer/backend/util"
	"github.com/spf13/viper"
)

var (
	RPCAddress  = "127.0.0.1:30002"
	PasswordKey = []byte("temp")
	BPRegisterPassword = "temp"
)

func InitConfig() {
	var host = os.Getenv("HOST") //"172.31.11.131";
	RPCAddress    = "http://" + host + ":30002"
	RPCAddress = viper.GetString("rpcHost")
	BPRegisterPassword = viper.GetString("BPRegisterPassword")
	fmt.Println("RPCHost:", RPCAddress)
	bytePassword, err := util.ReadPassword("Enter Password:  ")
	if err != nil {
		fmt.Println(err)
	}
	PasswordKey = bytePassword
}
