package restful

import (
	"encoding/json"
	"fmt"

	"github.com/bisonai/huobi_futures/sdk/coinfutures"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/restful/response/common"
	"github.com/bisonai/huobi_futures/sdk/log"
	"github.com/bisonai/huobi_futures/sdk/reqbuilder"
)

type CommonClient struct {
	PUrlBuilder *reqbuilder.PublicUrlBuilder
}

func (ac *CommonClient) Init(host string) *CommonClient {
	if host == "" {
		host = coinfutures.COIN_FUTURES_DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PublicUrlBuilder).Init(host)
	return ac
}

// GetContractRiskInfoAsync (Query information on contract insurance fund balance and estimated clawback rate)
func (ac *CommonClient) GetContractRiskInfoAsync(data chan common.GetContractRiskInfoResponse, symbol string) {
	// location
	location := "/api/v1/contract_risk_info"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractRiskInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractRiskInfoResponse error: %s", getErr)
	}
	data <- result
}

// GetContractInsuranceFundAsync (Query history records of insurance fund balance)
func (ac *CommonClient) GetContractInsuranceFundAsync(data chan common.GetContractInsuranceFundResponse, symbol string) {
	// location
	location := "/api/v1/contract_insurance_fund"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractInsuranceFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractInsuranceFundResponse error: %s", getErr)
	}
	data <- result
}

// GetContractAdjustFactorAsync (Query information on Tiered Adjustment Factor)
func (ac *CommonClient) GetContractAdjustFactorAsync(data chan common.GetContractAdjustFactorResponse, symbol string) {
	// location
	location := "/api/v1/contract_adjustfactor"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractAdjustFactorResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractAdjustFactorResponse error: %s", getErr)
	}
	data <- result
}

// GetContractHisOpenInterestAsync (Query information on open interest)
func (ac *CommonClient) GetContractHisOpenInterestAsync(data chan common.GetContractHisOpenInterestResponse, symbol string,
	contractType string, period string, size string, amountType string) {
	// location
	location := "/api/v1/contract_his_open_interest"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
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
	result := common.GetContractHisOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractHisOpenInterestResponse error: %s", getErr)
	}
	data <- result
}

// GetContractLadderMarginAsync (Query information on Tiered Margin)
func (ac *CommonClient) GetContractLadderMarginAsync(data chan common.GetContractLadderMarginResponse, symbol string) {
	// location
	location := "/api/v1/contract_ladder_margin"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractLadderMarginResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractLadderMarginResponse error: %s", getErr)
	}
	data <- result
}

// GetContractEliteAccountRatioAsync (Query Top Trader Sentiment Index Function-Account)
func (ac *CommonClient) GetContractEliteAccountRatioAsync(data chan common.GetContractEliteAccountRatioResponse, symbol string,
	period string) {
	// location
	location := "/api/v1/contract_elite_account_ratio"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
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
	result := common.GetContractEliteAccountRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractEliteAccountRatioResponse error: %s", getErr)
	}
	data <- result
}

// GetContractElitePositionRatioAsync (Query Top Trader Sentiment Index Function-Position)
func (ac *CommonClient) GetContractElitePositionRatioAsync(data chan common.GetContractElitePositionRatioResponse, symbol string,
	period string) {
	// location
	location := "/api/v1/contract_elite_position_ratio"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
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
	result := common.GetContractElitePositionRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractElitePositionRatioResponse error: %s", getErr)
	}
	data <- result
}

// GetContractLiquidationOrdersAsync (Query Liquidation Order Information(New))
func (ac *CommonClient) GetContractLiquidationOrdersAsync(data chan common.GetContractLiquidationOrdersResponse, symbol string,
	tradeType string, startTime string, endTime string, direct string, fromID string) {
	// location
	location := "/api/v3/contract_liquidation_orders"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if tradeType != "" {
		option += fmt.Sprintf("&trade_type=%s", tradeType)
	}
	if startTime != "" {
		option += fmt.Sprintf("&start_time=%s", startTime)
	}
	if endTime != "" {
		option += fmt.Sprintf("&end_time=%s", endTime)
	}
	if direct != "" {
		option += fmt.Sprintf("&direct=%s", direct)
	}
	if fromID != "" {
		option += fmt.Sprintf("&from_id=%s", fromID)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractLiquidationOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractLiquidationOrdersResponse error: %s", getErr)
	}
	data <- result
}

// GetContractSettlementRecordsAsync (Query historical settlement records of the platform interface)
func (ac *CommonClient) GetContractSettlementRecordsAsync(data chan common.GetContractSettlementRecordsResponse, symbol string,
	startTime string, endTime string, pageIndex string, pageSize string) {
	// location
	location := "/api/v1/contract_settlement_records"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
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
	result := common.GetContractSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractSettlementRecordsResponse error: %s", getErr)
	}
	data <- result
}

// GetContractPriceLimitAsync (Get Contract Price Limitation)
func (ac *CommonClient) GetContractPriceLimitAsync(data chan common.GetContractPriceLimitResponse, symbol string,
	contractType string, contractCode string) {
	// location
	location := "/api/v1/contract_price_limit"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if contractCode != "" {
		option += fmt.Sprintf("&contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractPriceLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractPriceLimitResponse error: %s", getErr)
	}
	data <- result
}

// GetContractOpenInterestAsync (Get Contract Open Interest Information)
func (ac *CommonClient) GetContractOpenInterestAsync(data chan common.GetContractOpenInterestResponse, symbol string,
	contractType string, contractCode string) {
	// location
	location := "/api/v1/contract_open_interest"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if contractCode != "" {
		option += fmt.Sprintf("&contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractOpenInterestResponse error: %s", getErr)
	}
	data <- result
}

// GetContractDeliveryPriceAsync (Get the estimated delivery price)
func (ac *CommonClient) GetContractDeliveryPriceAsync(data chan common.GetContractDeliveryPriceResponse, symbol string) {
	// location
	location := "/api/v1/contract_delivery_price"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractDeliveryPriceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractDeliveryPriceResponse error: %s", getErr)
	}
	data <- result
}

// GetContractEstimatedSettlementPriceAsync (Get The Estimated Settlement Price)
func (ac *CommonClient) GetContractEstimatedSettlementPriceAsync(data chan common.GetContractEstimatedSettlementPriceResponse, symbol string) {
	// location
	location := "/api/v1/contract_estimated_settlement_price"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractEstimatedSettlementPriceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractEstimatedSettlementPriceResponse error: %s", getErr)
	}
	data <- result
}

// GetContractApiStateAsync (Query information on system status)
func (ac *CommonClient) GetContractApiStateAsync(data chan common.GetContractApiStateResponse, symbol string) {
	// location
	location := "/api/v1/contract_api_state"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractApiStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractApiStateResponse error: %s", getErr)
	}
	data <- result
}

// GetContractContractInfoAsync (Get Contract Info)
func (ac *CommonClient) GetContractContractInfoAsync(data chan common.GetContractContractInfoResponse, symbol string, contractType string,
	contractCode string) {
	// location
	location := "/api/v1/contract_contract_info"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if contractCode != "" {
		option += fmt.Sprintf("&contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractContractInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractContractInfoResponse error: %s", getErr)
	}
	data <- result
}

// GetContractIndexAsync (Get Contract Index Price Information)
func (ac *CommonClient) GetContractIndexAsync(data chan common.GetContractIndexResponse, symbol string) {
	// location
	location := "/api/v1/contract_index"

	// option
	option := ""
	if symbol != "" {
		option += fmt.Sprintf("symbol=%s", symbol)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := ac.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetContractIndexResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractIndexResponse error: %s", getErr)
	}
	data <- result
}

// GetContractQueryElementsAsync (Contract Elements)
func (ac *CommonClient) GetContractQueryElementsAsync(data chan common.GetContractQueryElementsResponse, contractCode string) {
	// location
	location := "/api/v1/contract_query_elements"

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
	result := common.GetContractQueryElementsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractQueryElementsResponse error: %s", getErr)
	}
	data <- result
}
