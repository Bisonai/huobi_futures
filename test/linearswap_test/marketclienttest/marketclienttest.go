package marketclienttest

import (
	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/linearswap/restful"
	"github.com/bisonai/huobi_futures/sdk/linearswap/restful/response/market"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	GetContractInfoAsync()
	GetIndexAsync()
	GetPriceLimitAsync()
	GetOpenInterestAsync()
	GetDepthAsync()
	GetBboAsync()
	GetKLineAsync()
	GetMarkPriceKLineAsync()
	GetMergedAsync()
	GetTradeAsync()
	GetBatchMergedAsync()
	GetHisTradeAsync()
	GetRiskInfoAsync()
	GetInsuranceFundAsync()
	IsolatedGetAdjustFactorFundAsync()
	CrossGetAdjustFactorFundAsync()
	GetHisOpenInterestAsync()
	IsolatedGetLadderMarginAsync()
	CrossGetLadderMarginAsync()
	GetEliteAccountRatioAsync()
	GetElitePositionRatioAsync()
	IsolatedGetApiStateAsync()
	CrossGetTransferStateAsync()
	CrossGetTradeStateAsync()
	GetFundingRateAsync()
	GetBatchFundingRateAsync()
	GetHisFundingRateAsync()
	GetLiquidationOrdersAsync()
	GetPremiumIndexKLineAsync()
	GetEstimatedRateKLineAsync()
	GetBasisAsync()
	GetEstimatedSettlementPriceAsync()
	GetSwapLiquidationOrdersAsync()
	GetSwapSettlementRecordsAsync()
	GetSwapQueryElementsAsync()
	GetTimestampAsync()
	GetHeartbeatAsync()
	GetSummaryAsync()
	GetBatchMergedV2Async()
}

func GetContractInfoAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetContractInfoResponse)
	go client.GetContractInfoAsync(resp, "BTC-USDT", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetIndexAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetIndexResponse)
	go client.GetIndexAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetPriceLimitAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetPriceLimitResponse)
	go client.GetPriceLimitAsync(resp, "BTC-USDT", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetOpenInterestAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetOpenInterestResponse)
	go client.GetOpenInterestAsync(resp, "BTC-USDT", "", "", "")
	x := <-resp
	log.Info("%v", x)
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
	resp := make(chan market.GetBboResponse)
	go client.GetBboAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func GetKLineAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetKLineResponse)
	go client.GetKLineAsync(resp, "BTC-USDT", "", 100, 0, 100)
	x := <-resp
	log.Info("%v", x)
}

func GetMarkPriceKLineAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetStrKLineResponse)
	go client.GetMarkPriceKLineAsync(resp, "BTC-USDT", "", 100)
	x := <-resp
	log.Info("%v", x)
}

func GetMergedAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetMergedResponse)
	go client.GetMergedAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetBatchMergedAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetBatchMergedResponse)
	go client.GetBatchMergedAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func GetTradeAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetTradeResponse)
	go client.GetTradeAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func GetHisTradeAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetHisTradeResponse)
	go client.GetHisTradeAsync(resp, "BTC-USDT", 100)
	x := <-resp
	log.Info("%v", x)
}

func GetRiskInfoAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetRiskInfoResponse)
	go client.GetRiskInfoAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func GetInsuranceFundAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetInsuranceFundResponse)
	go client.GetInsuranceFundAsync(resp, "BTC-USDT", 0, 100)
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetAdjustFactorFundAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetAdjustFactorFundResponse)
	go client.IsolatedGetAdjustFactorFundAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetAdjustFactorFundAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetAdjustFactorFundResponse)
	go client.CrossGetAdjustFactorFundAsync(resp, "BTC-USDT", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetHisOpenInterestAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetHisOpenInterestResponse)
	go client.GetHisOpenInterestAsync(resp, "BTC-USDT", "", 1, 100, "", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetLadderMarginAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetLadderMarginResponse)
	go client.IsolatedGetLadderMarginAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetLadderMarginAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetLadderMarginResponse)
	go client.CrossGetLadderMarginAsync(resp, "BTC-USDT", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetEliteAccountRatioAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetEliteRatioResponse)
	go client.GetEliteAccountRatioAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func GetElitePositionRatioAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetEliteRatioResponse)
	go client.GetElitePositionRatioAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetApiStateAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetApiStateResponse)
	go client.IsolatedGetApiStateAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetTransferStateAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetTransferStateResponse)
	go client.CrossGetTransferStateAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetTradeStateAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetTradeStateResponse)
	go client.CrossGetTradeStateAsync(resp, "BTC-USDT", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetFundingRateAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetFundingRateResponse)
	go client.GetFundingRateAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetBatchFundingRateAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetBatchFundingRateResponse)
	go client.GetBatchFundingRateAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetHisFundingRateAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetHisFundingRateResponse)
	go client.GetHisFundingRateAsync(resp, "BTC-USDT", 0, 100)
	x := <-resp
	log.Info("%v", x)
}

func GetLiquidationOrdersAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetLiquidationOrdersResponse)
	go client.GetLiquidationOrdersAsync(resp, "BTC-USDT", 0, 100, 0, 100)
	x := <-resp
	log.Info("%v", x)
}

func GetPremiumIndexKLineAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetStrKLineResponse)
	go client.GetPremiumIndexKLineAsync(resp, "BTC-USDT", "", 100)
	x := <-resp
	log.Info("%v", x)
}

func GetEstimatedRateKLineAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetStrKLineResponse)
	go client.GetEstimatedRateKLineAsync(resp, "BTC-USDT", "", 100)
	x := <-resp
	log.Info("%v", x)
}

func GetBasisAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetBasisResponse)
	go client.GetBasisAsync(resp, "BTC-USDT", "", 100, "")
	x := <-resp
	log.Info("%v", x)
}

func GetEstimatedSettlementPriceAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetEstimatedSettlementPriceResponse)
	go client.GetEstimatedSettlementPriceAsync(resp, "BTC-USDT", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapLiquidationOrdersAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetSwapLiquidationOrdersResponse)
	go client.GetSwapLiquidationOrdersAsync(resp, "BTC-USDT", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapSettlementRecordsAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetSwapSettlementRecordsResponse)
	go client.GetSwapSettlementRecordsAsync(resp, "BTC-USDT", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapQueryElementsAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetSwapQueryElementsResponse)
	go client.GetSwapQueryElementsAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func GetTimestampAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetTimestampResponse)
	go client.GetTimestampAsync(resp)
	x := <-resp
	log.Info("%v", x)
}

func GetHeartbeatAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetHeartbeatResponse)
	go client.GetHeartbeatAsync(resp)
	x := <-resp
	log.Info("%v", x)
}

func GetSummaryAsync() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetSummaryResponse)
	go client.GetSummaryAsync(resp)
	x := <-resp
	log.Info("%v", x)
}

func GetBatchMergedV2Async() {
	client := new(restful.MarketClient).Init(config.Host)
	resp := make(chan market.GetBatchMergedV2Response)
	go client.GetBatchMergedV2Async(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}
