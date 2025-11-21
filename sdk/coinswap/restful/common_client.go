package restful

import (
	"encoding/json"
	"fmt"

	"github.com/bisonai/huobi_futures/sdk/coinswap"
	"github.com/bisonai/huobi_futures/sdk/coinswap/restful/response/common"
	"github.com/bisonai/huobi_futures/sdk/log"
	"github.com/bisonai/huobi_futures/sdk/reqbuilder"
)

type CommonClient struct {
	PUrlBuilder *reqbuilder.PublicUrlBuilder
}

func (ac *CommonClient) Init(host string) *CommonClient {
	if host == "" {
		host = coinswap.COIN_SWAP_DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PublicUrlBuilder).Init(host)
	return ac
}

// GetSwapRiskInfoAsync (Query information on contract insurance fund balance and estimated clawback rate)
func (ac *CommonClient) GetSwapRiskInfoAsync(data chan common.GetSwapRiskInfoResponse, contractCode string) {
	// location
	location := "/swap-api/v1/swap_risk_info"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapRiskInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapRiskInfoResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapInsuranceFundAsync (Query history records of insurance fund balance)
func (ac *CommonClient) GetSwapInsuranceFundAsync(data chan common.GetSwapInsuranceFundResponse, contractCode string,
	pageIndex string, pageSize string) {
	// location
	location := "/swap-api/v1/swap_insurance_fund"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if pageIndex != "" {
		option += fmt.Sprintf("&page_index=%s", pageIndex)
	}
	if pageSize != "" {
		option += fmt.Sprintf("&page_size=%s", pageSize)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapInsuranceFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapInsuranceFundResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapAdjustFactorAsync (Query information on Tiered Adjustment Factor)
func (ac *CommonClient) GetSwapAdjustFactorAsync(data chan common.GetSwapAdjustFactorResponse, contractCode string) {
	// location
	location := "/swap-api/v1/swap_adjustfactor"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapAdjustFactorResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapAdjustFactorResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapHisOpenInterestAsync (Query information on open interest)
func (ac *CommonClient) GetSwapHisOpenInterestAsync(data chan common.GetSwapHisOpenInterestResponse, contractCode string,
	period string, size string, amountType string) {
	// location
	location := "/swap-api/v1/swap_his_open_interest"

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
	if amountType != "" {
		option += fmt.Sprintf("&amount_type=%s", amountType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapHisOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapHisOpenInterestResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapLadderMarginAsync (Query information on Tiered Margin interface)
func (ac *CommonClient) GetSwapLadderMarginAsync(data chan common.GetSwapLadderMarginResponse, contractCode string) {
	// location
	location := "/swap-api/v1/swap_ladder_margin"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapLadderMarginResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapLadderMarginResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapEliteAccountRatioAsync (Query Top Trader Sentiment Index Function-Account)
func (ac *CommonClient) GetSwapEliteAccountRatioAsync(data chan common.GetSwapEliteAccountRatioResponse, contractCode string,
	period string) {
	// location
	location := "/swap-api/v1/swap_elite_account_ratio"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if period != "" {
		option += fmt.Sprintf("&period=%s", period)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapEliteAccountRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapEliteAccountRatioResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapElitePositionRatioAsync (Query Top Trader Sentiment Index Function-Position)
func (ac *CommonClient) GetSwapElitePositionRatioAsync(data chan common.GetSwapElitePositionRatioResponse, contractCode string,
	period string) {
	// location
	location := "/swap-api/v1/swap_elite_position_ratio"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if period != "" {
		option += fmt.Sprintf("&period=%s", period)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapElitePositionRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapElitePositionRatioResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapEstimatedSettlementPriceAsync (Get the estimated settlement price)
func (ac *CommonClient) GetSwapEstimatedSettlementPriceAsync(data chan common.GetSwapEstimatedSettlementPriceResponse,
	contractCode string) {
	// location
	location := "/swap-api/v1/swap_estimated_settlement_price"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapEstimatedSettlementPriceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapEstimatedSettlementPriceResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapApiStateAsync (Query information on system status)
func (ac *CommonClient) GetSwapApiStateAsync(data chan common.GetSwapApiStateResponse,
	contractCode string) {
	// location
	location := "/swap-api/v1/swap_api_state"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapApiStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapApiStateResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapFundingRateAsync (Query funding rate)
func (ac *CommonClient) GetSwapFundingRateAsync(data chan common.GetSwapFundingRateResponse,
	contractCode string) {
	// location
	location := "/swap-api/v1/swap_funding_rate"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapFundingRateResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapBatchFundingRateAsync (Query a Batch of Funding Rate)
func (ac *CommonClient) GetSwapBatchFundingRateAsync(data chan common.GetSwapBatchFundingRateResponse,
	contractCode string) {
	// location
	location := "/swap-api/v1/swap_batch_funding_rate"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapBatchFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapBatchFundingRateResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapHistoricalFundingRateAsync (Query historical funding rate)
func (ac *CommonClient) GetSwapHistoricalFundingRateAsync(data chan common.GetSwapHistoricalFundingRateResponse,
	contractCode string, pageIndex string, pageSize string) {
	// location
	location := "/swap-api/v1/swap_historical_funding_rate"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if pageIndex != "" {
		option += fmt.Sprintf("&page_index=%s", pageIndex)
	}
	if pageSize != "" {
		option += fmt.Sprintf("&page_size=%s", pageSize)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapHistoricalFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapHistoricalFundingRateResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapLiquidationOrdersAsync (Query Liquidation Orders(New))
func (ac *CommonClient) GetSwapLiquidationOrdersAsync(data chan common.GetSwapLiquidationOrdersResponse,
	startTime string, endTime string, direct string, fromId string, contract string, tradeType string) {
	// location
	location := "/swap-api/v3/swap_liquidation_orders"

	// option
	option := ""
	if startTime != "" {
		option += fmt.Sprintf("start_time=%s", startTime)
	}
	if endTime != "" {
		option += fmt.Sprintf("&end_time=%s", endTime)
	}
	if direct != "" {
		option += fmt.Sprintf("&direct=%s", direct)
	}
	if fromId != "" {
		option += fmt.Sprintf("&from_id=%s", fromId)
	}
	if contract != "" {
		option += fmt.Sprintf("&contract=%s", contract)
	}
	if tradeType != "" {
		option += fmt.Sprintf("&trade_type=%s", tradeType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapLiquidationOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapLiquidationOrdersResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapSettlementRecordsAsync (Query historical settlement records of the platform interface)
func (ac *CommonClient) GetSwapSettlementRecordsAsync(data chan common.GetSwapSettlementRecordsResponse, contractCode string,
	startTime string, endTime string, pageIndex string, pageSize string) {
	// location
	location := "/swap-api/v1/swap_settlement_records"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if startTime != "" {
		option += fmt.Sprintf("&start_time=%s", startTime)
	}
	if endTime != "" {
		option += fmt.Sprintf("&end_time=%s", endTime)
	}
	if pageIndex != "" {
		option += fmt.Sprintf("&page_index=%s", pageIndex)
	}
	if pageSize != "" {
		option += fmt.Sprintf("&page_size=%s", pageSize)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapSettlementRecordsResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapContractInfoAsync (Query Swap Info)
func (ac *CommonClient) GetSwapContractInfoAsync(data chan common.GetSwapContractInfoResponse, contractCode string) {
	// location
	location := "/swap-api/v1/swap_contract_info"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapContractInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapContractInfoResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapIndexAsync (Query Swap Index Price Information)
func (ac *CommonClient) GetSwapIndexAsync(data chan common.GetSwapIndexResponse, contractCode string) {
	// location
	location := "/swap-api/v1/swap_index"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapIndexResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapIndexResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapQueryElementsAsync (Contract Elements)
func (ac *CommonClient) GetSwapQueryElementsAsync(data chan common.GetSwapQueryElementsResponse, contractCode string) {
	// location
	location := "/swap-api/v1/swap_query_elements"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapQueryElementsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapQueryElementsResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapPriceLimitAsync (Query Swap Price Limitation)
func (ac *CommonClient) GetSwapPriceLimitAsync(data chan common.GetSwapPriceLimitResponse, contractCode string) {
	// location
	location := "/swap-api/v1/swap_price_limit"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapPriceLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapPriceLimitResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapOpenInterestAsync (Get Swap Open Interest Information)
func (ac *CommonClient) GetSwapOpenInterestAsync(data chan common.GetSwapOpenInterestResponse, contractCode string) {
	// location
	location := "/swap-api/v1/swap_open_interest"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapOpenInterestResponse error: %s", getErr)
	}
	data <- result
}
