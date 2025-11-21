package restful

import (
	"encoding/json"
	"fmt"

	"github.com/bisonai/huobi_futures/sdk/coinfutures"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/restful/response/market"

	"github.com/bisonai/huobi_futures/sdk/log"
	"github.com/bisonai/huobi_futures/sdk/reqbuilder"
)

type MarketClient struct {
	PUrlBuilder *reqbuilder.PublicUrlBuilder
}

func (mc *MarketClient) Init(host string) *MarketClient {
	if host == "" {
		host = coinfutures.COIN_FUTURES_DEFAULT_HOST
	}
	mc.PUrlBuilder = new(reqbuilder.PublicUrlBuilder).Init(host)
	return mc
}

// GetDepthAsync (Get Market Depth)
func (mc *MarketClient) GetDepthAsync(data chan market.GetDepthResponse, symbol string, fcType string) {
	// location
	location := "/market/depth"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if fcType != "" {
		option += fmt.Sprintf("&type=%s", fcType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetDepthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetDepthResponse error: %s", getErr)
	}
	data <- result
}

// GetBboAsync (Get Market BBO Data)
func (mc *MarketClient) GetBboAsync(data chan market.GetMarketBboResponse, symbol string) {
	// location
	location := "/market/bbo"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetMarketBboResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetMarketBboResponse error: %s", getErr)
	}
	data <- result
}

// GetKlineAsync (Get Kline Data)
func (mc *MarketClient) GetKlineAsync(data chan market.GetKlineResponse, symbol string, period string, size string,
	from string, to string) {
	// location
	location := "/market/history/kline"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if period != "" {
		option += fmt.Sprintf("&period=%s", period)
	}
	if size != "" {
		option += fmt.Sprintf("&size=%s", size)
	}
	if from != "" {
		option += fmt.Sprintf("&from=%s", from)
	}
	if to != "" {
		option += fmt.Sprintf("&to=%s", to)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetKlineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetKlineResponse error: %s", getErr)
	}
	data <- result
}

// GetMarkPriceKlineAsync (Get Kline Data of Mark Price)
func (mc *MarketClient) GetMarkPriceKlineAsync(data chan market.GetMarkPriceKlineResponse, symbol string, period string, size string) {
	// location
	location := "/index/market/history/mark_price_kline"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if period != "" {
		option += fmt.Sprintf("&period=%s", period)
	}
	if size != "" {
		option += fmt.Sprintf("&size=%s", size)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetMarkPriceKlineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetMarkPriceKlineResponse error: %s", getErr)
	}
	data <- result
}

// GetMergedAsync (Get Market Data Overview)
func (mc *MarketClient) GetMergedAsync(data chan market.GetDetailMergedResponse, symbol string) {
	// location
	location := "/market/detail/merged"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetDetailMergedResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetDetailMergedResponse error: %s", getErr)
	}
	data <- result
}

// GetBatchMergedAsync (Get a Batch of Market Data Overview(V2))
func (mc *MarketClient) GetBatchMergedAsync(data chan market.GetdetailBatchMergedResponse, symbol string) {
	// location
	location := "/v2/market/detail/batch_merged"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetdetailBatchMergedResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetdetailBatchMergedResponse error: %s", getErr)
	}
	data <- result
}

// GetTradeAsync (Query The Last Trade of a Contract)
func (mc *MarketClient) GetTradeAsync(data chan market.GetMarketTradeResponse, symbol string) {
	// location
	location := "/market/trade"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetMarketTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetMarketTradeResponse error: %s", getErr)
	}
	data <- result
}

// GetHistoryTradeAsync (Query a Batch of Trade Records of a Contract)
func (mc *MarketClient) GetHistoryTradeAsync(data chan market.GetHistoryTradeResponse, symbol string, size string) {
	// location
	location := "/market/history/trade"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if size != "" {
		option += fmt.Sprintf("&size=%s", size)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHistoryTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHistoryTradeResponse error: %s", getErr)
	}
	data <- result
}

// GetHistoryIndexAsync (Query Index Kline Data)
func (mc *MarketClient) GetHistoryIndexAsync(data chan market.GetHistoryIndexResponse, symbol string, period string, size string) {
	// location
	location := "/index/market/history/index"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if period != "" {
		option += fmt.Sprintf("&period=%s", period)
	}
	if size != "" {
		option += fmt.Sprintf("&size=%s", size)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHistoryIndexResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHistoryIndexResponse error: %s", getErr)
	}
	data <- result
}

// GetHistoryBasisAsync (Query Basis Data)
func (mc *MarketClient) GetHistoryBasisAsync(data chan market.GetHistoryBasisResponse, symbol string, period string,
	basisPriceType string, size string) {
	// location
	location := "/index/market/history/basis"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if period != "" {
		option += fmt.Sprintf("&period=%s", period)
	}
	if basisPriceType != "" {
		option += fmt.Sprintf("&basis_price_type=%s", basisPriceType)
	}
	if size != "" {
		option += fmt.Sprintf("&size=%s", size)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHistoryBasisResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHistoryBasisResponse error: %s", getErr)
	}
	data <- result
}
