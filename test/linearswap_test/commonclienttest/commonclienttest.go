package commonclienttest

import (
	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/linearswap/restful"
	"github.com/bisonai/huobi_futures/sdk/linearswap/restful/response/common"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	GetSwapUnifiedAccountTypeAsync()
	SwapSwitchAccountTypeAsync()
}

func GetSwapUnifiedAccountTypeAsync() {
	client := new(restful.CommonClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp := make(chan common.GetSwapUnifiedAccountTypeResponse)
	go client.GetSwapUnifiedAccountTypeAsync(resp)
	x := <-resp
	log.Info("%v", x)
}

func SwapSwitchAccountTypeAsync() {
	client := new(restful.CommonClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp := make(chan common.SwapSwitchAccountTypeResponse)
	go client.SwapSwitchAccountTypeAsync(resp, 0)
	x := <-resp
	log.Info("%v", x)
}
