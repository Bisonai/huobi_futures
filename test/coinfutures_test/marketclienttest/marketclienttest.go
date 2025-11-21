package marketclienttest

import (
	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/restful"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/restful/response/market"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	GetDepthAsync()
	GetBboAsync()
	GetKlineAsync()
	GetMarkPriceKlineAsync()
	GetMergedAsync()
	GetBatchMergedAsync()
	GetTradeAsync()
	GetHistoryTradeAsync()
	GetHistoryIndexAsync()
	GetHistoryBasisAsync()
}

func GetDepthAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetDepthResponse)
	go client.GetDepthAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func GetBboAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetMarketBboResponse)
	go client.GetBboAsync(resp, "BTC-USDT")
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

func GetMarkPriceKlineAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetMarkPriceKlineResponse)
	go client.GetMarkPriceKlineAsync(resp, "BTC-USDT", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetMergedAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetDetailMergedResponse)
	go client.GetMergedAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetBatchMergedAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetdetailBatchMergedResponse)
	go client.GetBatchMergedAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetTradeAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetMarketTradeResponse)
	go client.GetTradeAsync(resp, "BTC-USDT")
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

func GetHistoryIndexAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetHistoryIndexResponse)
	go client.GetHistoryIndexAsync(resp, "BTC-USDT", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetHistoryBasisAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetHistoryBasisResponse)
	go client.GetHistoryBasisAsync(resp, "BTC-USDT", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
