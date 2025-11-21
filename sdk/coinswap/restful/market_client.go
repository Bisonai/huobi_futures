package restful

import (
	"encoding/json"
	"fmt"

	"github.com/bisonai/huobi_futures/sdk/coinswap"
	"github.com/bisonai/huobi_futures/sdk/coinswap/restful/response/market"
	"github.com/bisonai/huobi_futures/sdk/log"
	"github.com/bisonai/huobi_futures/sdk/reqbuilder"
)

type MarketClient struct {
	PUrlBuilder *reqbuilder.PublicUrlBuilder
}

func (mc *MarketClient) Init(host string) *MarketClient {
	if host == "" {
		host = coinswap.COIN_SWAP_DEFAULT_HOST
	}
	mc.PUrlBuilder = new(reqbuilder.PublicUrlBuilder).Init(host)
	return mc
}

// GetMarketDepthAsync (Get Market Depth)
func (mc *MarketClient) GetMarketDepthAsync(data chan market.GetMarketDepthResponse, symbol string, fcType string) {
	// location
	location := "/swap-ex/market/depth"

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
	result := market.GetMarketDepthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetMarketDepthResponse error: %s", getErr)
	}
	data <- result
}

// GetMarketBboAsync (Get Market BBO Data)
func (mc *MarketClient) GetMarketBboAsync(data chan market.GetMarketBboResponse, contractCode string) {
	// location
	location := "/swap-ex/market/bbo"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
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

// GetKlineAsync (Get KLine Data)
func (mc *MarketClient) GetKlineAsync(data chan market.GetKlineResponse, contractCode string, period string,
	size string, from string, to string) {
	// location
	location := "/swap-ex/market/history/kline"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
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

// GetSwapMarkPriceKlineAsync (Get Kline Data of Mark Price)
func (mc *MarketClient) GetSwapMarkPriceKlineAsync(data chan market.GetSwapMarkPriceKlineResponse, contractCode string, period string,
	size string) {
	// location
	location := "/index/market/history/swap_mark_price_kline"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
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
	result := market.GetSwapMarkPriceKlineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapMarkPriceKlineResponse error: %s", getErr)
	}
	data <- result
}

// GetDetailMergedAsync (Get Market Data Overview)
func (mc *MarketClient) GetDetailMergedAsync(data chan market.GetDetailMergedResponse, contractCode string) {
	// location
	location := "/swap-ex/market/detail/merged"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
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

// GetDetailBatchMergedAsync (Get a Batch of Market Data Overview(V2))
func (mc *MarketClient) GetDetailBatchMergedAsync(data chan market.GetDetailBatchMergedResponse, contractCode string) {
	// location
	location := "/v2/swap-ex/market/detail/batch_merged"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetDetailBatchMergedResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetDetailBatchMergedResponse error: %s", getErr)
	}
	data <- result
}

// GetMarketTradeAsync (Query The Last Trade of a Contract)
func (mc *MarketClient) GetMarketTradeAsync(data chan market.GetMarketTradeResponse, contractCode string) {
	// location
	location := "/swap-ex/market/trade"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
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
func (mc *MarketClient) GetHistoryTradeAsync(data chan market.GetHistoryTradeResponse, contractCode string, size string) {
	// location
	location := "/swap-ex/market/history/trade"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
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

// GetSwapPremiumIndexKlineAsync (Query Premium Index Kline Data)
func (mc *MarketClient) GetSwapPremiumIndexKlineAsync(data chan market.GetSwapPremiumIndexKlineResponse, contractCode string,
	period string, size string) {
	// location
	location := "/index/market/history/swap_premium_index_kline"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
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
	result := market.GetSwapPremiumIndexKlineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapPremiumIndexKlineResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapEstimatedRateKlineAsync (Query Estimated Funding Rate Kline Data)
func (mc *MarketClient) GetSwapEstimatedRateKlineAsync(data chan market.GetSwapEstimatedRateKlineResponse, contractCode string,
	period string, size string) {
	// location
	location := "/index/market/history/swap_estimated_rate_kline"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
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
	result := market.GetSwapEstimatedRateKlineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapEstimatedRateKlineResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapBasisAsync (Query Basis Data)
func (mc *MarketClient) GetSwapBasisAsync(data chan market.GetSwapBasisResponse, contractCode string,
	period string, basisPriceType string, size string) {
	// location
	location := "/index/market/history/swap_basis"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
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
	result := market.GetSwapBasisResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapBasisResponse error: %s", getErr)
	}
	data <- result
}
