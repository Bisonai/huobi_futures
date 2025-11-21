package commonclienttest

import (
	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/restful"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/restful/response/common"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	GetContractRiskInfoAsync()
	GetContractInsuranceFundAsync()
	GetContractAdjustFactorAsync()
	GetContractHisOpenInterestAsync()
	GetContractLadderMarginAsync()
	GetContractEliteAccountRatioAsync()
	GetContractElitePositionRatioAsync()
	GetContractLiquidationOrdersAsync()
	GetContractSettlementRecordsAsync()
	GetContractPriceLimitAsync()
	GetContractOpenInterestAsync()
	GetContractDeliveryPriceAsync()
	GetContractEstimatedSettlementPriceAsync()
	GetContractApiStateAsync()
	GetContractContractInfoAsync()
	GetContractIndexAsync()
	GetContractQueryElementsAsync()
}

func GetContractRiskInfoAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractRiskInfoResponse)
	go client.GetContractRiskInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractInsuranceFundAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractInsuranceFundResponse)
	go client.GetContractInsuranceFundAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractAdjustFactorAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractAdjustFactorResponse)
	go client.GetContractAdjustFactorAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractHisOpenInterestAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractHisOpenInterestResponse)
	go client.GetContractHisOpenInterestAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractLadderMarginAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractLadderMarginResponse)
	go client.GetContractLadderMarginAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractEliteAccountRatioAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractEliteAccountRatioResponse)
	go client.GetContractEliteAccountRatioAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractElitePositionRatioAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractElitePositionRatioResponse)
	go client.GetContractElitePositionRatioAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractLiquidationOrdersAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractLiquidationOrdersResponse)
	go client.GetContractLiquidationOrdersAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractSettlementRecordsAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractSettlementRecordsResponse)
	go client.GetContractSettlementRecordsAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractPriceLimitAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractPriceLimitResponse)
	go client.GetContractPriceLimitAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractOpenInterestAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractOpenInterestResponse)
	go client.GetContractOpenInterestAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractDeliveryPriceAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractDeliveryPriceResponse)
	go client.GetContractDeliveryPriceAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractEstimatedSettlementPriceAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractEstimatedSettlementPriceResponse)
	go client.GetContractEstimatedSettlementPriceAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractApiStateAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractApiStateResponse)
	go client.GetContractApiStateAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractContractInfoAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractContractInfoResponse)
	go client.GetContractContractInfoAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractIndexAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractIndexResponse)
	go client.GetContractIndexAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractQueryElementsAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetContractQueryElementsResponse)
	go client.GetContractQueryElementsAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}
