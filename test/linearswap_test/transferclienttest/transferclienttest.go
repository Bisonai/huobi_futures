package transferclienttest

import (
	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/linearswap/restful"
	"github.com/bisonai/huobi_futures/sdk/linearswap/restful/response/transfer"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	TransferAsync()
}

func TransferAsync() {
	client := new(restful.TransferClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp := make(chan transfer.TransferResponse)
	go client.TransferAsync(resp, "", "", 0, "", "")
	x := <-resp
	log.Info("%v", x)
}
