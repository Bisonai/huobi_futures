package wsindexclienttest

import (
	"time"

	"github.com/bisonai/huobi_futures/sdk/coinfutures/ws"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/ws/response/index"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	SubMarkPriceKLine()
	ReqMarkPriceKLine()
	SubBasis()
	ReqBasis()
	SubIndexKLine()
	ReqIndexKLine()
}

func SubMarkPriceKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("")
	wsixClient.SubMarkPriceKLine("BTC-USDT", "15min", func(m *index.SubIndexKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func ReqMarkPriceKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("")
	wsixClient.ReqMarkPriceKLine("BTC-USDT", "1min", func(m *index.ReqIndexKLineResponse) {
		log.Info("%v", *m)
	}, 1612434060, 1612434260, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func SubBasis() {
	var wsixClient = new(ws.WSIndexClient).Init("")
	wsixClient.SubBasis("BTC-USDT", "15min", func(m *index.SubBasiesResponse) {
		log.Info("%v", *m)
	}, "", "")
	time.Sleep(time.Duration(20) * time.Second)
}

func ReqBasis() {
	var wsixClient = new(ws.WSIndexClient).Init("")
	wsixClient.ReqBasis("BTC-USDT", "15min", func(m *index.ReqBasisResponse) {
		log.Info("%v", *m)
	}, 1604395758, 1604396758, "", "")
	time.Sleep(time.Duration(20) * time.Second)
}

func SubIndexKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("")
	wsixClient.SubIndexKLine("BTC-USDT", "15min", func(m *index.SubIndexKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func ReqIndexKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("")
	wsixClient.ReqIndexKLine("BTC-USDT", "15min", func(m *index.ReqIndexKLineResponse) {
		log.Info("%v", *m)
	}, 1604395758, 1604396758, "")
	time.Sleep(time.Duration(20) * time.Second)
}
