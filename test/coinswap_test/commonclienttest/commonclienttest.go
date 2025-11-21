package commonclienttest

import (
	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/coinswap/restful"
	"github.com/bisonai/huobi_futures/sdk/coinswap/restful/response/common"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	GetSwapRiskInfoAsync()
	GetSwapInsuranceFundAsync()
	GetSwapAdjustFactorAsync()
	GetSwapHisOpenInterestAsync()
	GetSwapLadderMarginAsync()
	GetSwapEliteAccountRatioAsync()
	GetSwapElitePositionRatioAsync()
	GetSwapEstimatedSettlementPriceAsync()
	GetSwapApiStateAsync()
	GetSwapFundingRateAsync()
	GetSwapBatchFundingRateAsync()
	GetSwapHistoricalFundingRateAsync()
	GetSwapLiquidationOrdersAsync()
	GetSwapSettlementRecordsAsync()
	GetSwapContractInfoAsync()
	GetSwapIndexAsync()
	GetSwapQueryElementsAsync()
	GetSwapPriceLimitAsync()
	GetSwapOpenInterestAsync()
}

func GetSwapRiskInfoAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapRiskInfoResponse)
	go client.GetSwapRiskInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapInsuranceFundAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapInsuranceFundResponse)
	go client.GetSwapInsuranceFundAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapAdjustFactorAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapAdjustFactorResponse)
	go client.GetSwapAdjustFactorAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapHisOpenInterestAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapHisOpenInterestResponse)
	go client.GetSwapHisOpenInterestAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapLadderMarginAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapLadderMarginResponse)
	go client.GetSwapLadderMarginAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapEliteAccountRatioAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapEliteAccountRatioResponse)
	go client.GetSwapEliteAccountRatioAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapElitePositionRatioAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapElitePositionRatioResponse)
	go client.GetSwapElitePositionRatioAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapEstimatedSettlementPriceAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapEstimatedSettlementPriceResponse)
	go client.GetSwapEstimatedSettlementPriceAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapApiStateAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapApiStateResponse)
	go client.GetSwapApiStateAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapFundingRateAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapFundingRateResponse)
	go client.GetSwapFundingRateAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapBatchFundingRateAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapBatchFundingRateResponse)
	go client.GetSwapBatchFundingRateAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapHistoricalFundingRateAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapHistoricalFundingRateResponse)
	go client.GetSwapHistoricalFundingRateAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapLiquidationOrdersAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapLiquidationOrdersResponse)
	go client.GetSwapLiquidationOrdersAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapSettlementRecordsAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapSettlementRecordsResponse)
	go client.GetSwapSettlementRecordsAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapContractInfoAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapContractInfoResponse)
	go client.GetSwapContractInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapIndexAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapIndexResponse)
	go client.GetSwapIndexAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapQueryElementsAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapQueryElementsResponse)
	go client.GetSwapQueryElementsAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapPriceLimitAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapPriceLimitResponse)
	go client.GetSwapPriceLimitAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapOpenInterestAsync() {
	client := new(restful.CommonClient).Init(config.Host)
	resp := make(chan common.GetSwapOpenInterestResponse)
	go client.GetSwapOpenInterestAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}
