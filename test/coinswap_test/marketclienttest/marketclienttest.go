package marketclienttest

import (
	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/coinswap/restful"
	"github.com/bisonai/huobi_futures/sdk/coinswap/restful/response/market"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	GetMarketDepthAsync()
	GetMarketBboAsync()
	GetKlineAsync()
	GetSwapMarkPriceKlineAsync()
	GetDetailMergedAsync()
	GetDetailBatchMergedAsync()
	GetMarketTradeAsync()
	GetHistoryTradeAsync()
	GetSwapPremiumIndexKlineAsync()
	GetSwapEstimatedRateKlineAsync()
	GetSwapBasisAsync()
}

func GetMarketDepthAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetMarketDepthResponse)
	go client.GetMarketDepthAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func GetMarketBboAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetMarketBboResponse)
	go client.GetMarketBboAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetKlineAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetKlineResponse)
	go client.GetKlineAsync(resp, "BTC-USDT", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapMarkPriceKlineAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetSwapMarkPriceKlineResponse)
	go client.GetSwapMarkPriceKlineAsync(resp, "BTC-USDT", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetDetailMergedAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetDetailMergedResponse)
	go client.GetDetailMergedAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetDetailBatchMergedAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetDetailBatchMergedResponse)
	go client.GetDetailBatchMergedAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetMarketTradeAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetMarketTradeResponse)
	go client.GetMarketTradeAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetHistoryTradeAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetHistoryTradeResponse)
	go client.GetHistoryTradeAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapPremiumIndexKlineAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetSwapPremiumIndexKlineResponse)
	go client.GetSwapPremiumIndexKlineAsync(resp, "BTC-USDT", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapEstimatedRateKlineAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetSwapEstimatedRateKlineResponse)
	go client.GetSwapEstimatedRateKlineAsync(resp, "BTC-USDT", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapBasisAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetSwapBasisResponse)
	go client.GetSwapBasisAsync(resp, "BTC-USDT", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
