package transferclienttest

import (
	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/restful"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/restful/response/transfer"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	FuturesTransferAsync()
	AccountTransferAsync()
}

func FuturesTransferAsync() {
	client := new(restful.TransferClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp := make(chan transfer.FuturesTransferResponse)
	go client.FuturesTransferAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func AccountTransferAsync() {
	client := new(restful.TransferClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp := make(chan transfer.AccountTransferResponse)
	go client.AccountTransferAsync(resp, "", "", 0, "", "")
	x := <-resp
	log.Info("%v", x)
}
