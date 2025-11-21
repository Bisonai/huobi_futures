package transferclienttest

import (
	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/coinswap/restful"
	"github.com/bisonai/huobi_futures/sdk/coinswap/restful/response/transfer"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	AccountTransferAsync()
}

func AccountTransferAsync() {
	client := new(restful.TransferClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp := make(chan transfer.AccountTransferResponse)
	go client.AccountTransferAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
