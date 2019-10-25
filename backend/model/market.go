package model

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/iost-official/explorer/backend/model/cache"
	"github.com/iost-official/explorer/backend/util"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	CoinMarketCapCacheKey     = "coinMarketCapInfo"
	CoinMarketCapIOSOfBTCTUrl = "https://api.coinmarketcap.com/v2/ticker/2405?convert=BTC"
	CoinMarketCapIOSOfETHTUrl = "https://api.coinmarketcap.com/v2/ticker/2405?convert=ETH"
)

type MarketInfo struct {
	Price            string  `json:"price"`
	Volume24h        int64   `json:"volume24h"`
	PercentChange24h float64 `json:"percentChange24h"`
	MarketCap        int64   `json:"marketCap"`
	BtcPrice         string  `json:"btcPrice"`
	EthPrice         string  `json:"ethPrice"`
	LastUpdate       string  `json:"lastUpdate"`
}

