package main

import (
	"github.com/iost-official/explorer/backend/config"
	"github.com/iost-official/explorer/backend/controller"
	"github.com/iost-official/explorer/backend/middleware"
	"github.com/labstack/echo"
	echoMiddle "github.com/labstack/echo/middleware"
)

func main() {
	config.ReadConfig("./config")
	controller.SyncBP()

	e := echo.New()
	e.Debug = true
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler
	e.Use(middleware.CorsHeader)
	e.Use(echoMiddle.Recover())
	e.Use(echoMiddle.Logger())

	// index
	e.GET("/api/indexBlocks", controller.GetIndexBlocks)
	e.GET("/api/indexTxns", controller.GetIndexTxns)

	/* // blocks */
	e.GET("/api/blocks", controller.GetBlocks)
	e.GET("/api/block/:id", controller.GetBlockDetail)

	// transactions
	e.GET("/api/txs", controller.GetTxs)
	e.GET("/api/tx/:id", controller.GetTxnDetail)

	// accounts
	e.GET("/api/accounts", controller.GetAccounts)
	e.GET("/api/account/:id", controller.GetAccountDetail)
	e.GET("/api/account/:id/txs", controller.GetAccountTxs)

	// search
	e.GET("/api/search/:id", controller.GetSearch)

	// mail
	e.POST("/api/feedback", controller.SendMail)

	// applyIOST
	e.POST("/api/applyIOST", controller.ApplyIOST)

	// bp list
	e.GET("/api/BPList", controller.GetBPList)
	e.GET("/api/BPLastProducer", controller.GetBPLastProducer)

	// register BP
	e.GET("/api/BPRegister", controller.RegistBP)

	// voteAward
	e.POST("/api/SetVoteAwardInfo", controller.SetVoteAwardInfo)
	e.POST("/api/SetProducerContributions", controller.SetProducerContributions)
	e.GET("/api/GetUserAward", controller.GetUserAward)
	e.GET("/api/GetUserContributionAward", controller.GetUserContributionAward)
	e.GET("/api/GetProducerContributionAward", controller.GetProducerContributionAward)
	e.GET("/api/GetProducerAward", controller.GetProducerAward)
	//e.GET("/api/CalculateAward", controller.CalculateAward)
	e.GET("/api/GetVoteAwardList", controller.GetVoteAwardList)
	e.GET("/api/GetVoteAwardInfo", controller.GetVoteAwardInfo)

	e.Logger.Fatal(e.Start(":8088"))
}
