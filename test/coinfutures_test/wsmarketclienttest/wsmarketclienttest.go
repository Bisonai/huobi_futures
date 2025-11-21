package wsmarketclienttest

import (
	"time"

	"github.com/bisonai/huobi_futures/sdk/coinfutures/ws"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/ws/response/market"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	SubKLine()
	ReqKLine()
	SubDepth()
	SubIncrementalDepth()
	SubDetail()
	SubBBO()
	SubTradeDetail()
	ReqTradeDetail()
}
func SubKLine() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com")
	wsmkClient.SubKLine("BTC-USDT", "15min", func(m *market.SubKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(10) * time.Second)
}

func ReqKLine() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com")
	wsmkClient.ReqKLine("BTC-USDT", "1min", func(m *market.ReqKLineResponse) {
		log.Info("%v", *m)
	}, 1604395758, 1604396758, "")
	time.Sleep(time.Duration(10) * time.Second)
}

func SubDepth() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com")
	wsmkClient.SubDepth("BTC-USDT", "step0", func(m *market.SubDepthResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func SubIncrementalDepth() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com")
	wsmkClient.SubIncrementalDepth("BTC-USDT", "20", func(m *market.SubDepthResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func SubDetail() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com")
	wsmkClient.SubDetail("BTC-USDT", func(m *market.SubKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func SubBBO() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com")
	wsmkClient.SubBBO("BTC-USDT", func(m *market.SubBBOResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func SubTradeDetail() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com")
	wsmkClient.SubTradeDetail("BTC-USDT", func(m *market.SubTradeDetailResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func ReqTradeDetail() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com")
	wsmkClient.ReqTradeDetail("BTC-USDT", func(m *market.ReqTradeDetailResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}
