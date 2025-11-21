package wscenternotifyclienttest

import (
	"time"

	"github.com/bisonai/huobi_futures/config"
	"github.com/bisonai/huobi_futures/sdk/linearswap/ws"
	"github.com/bisonai/huobi_futures/sdk/linearswap/ws/response/centernotify"
	"github.com/bisonai/huobi_futures/sdk/log"
)

func RunAllExamples() {
	Heartbeat()
}
func Heartbeat() {
	var wscnfClient = new(ws.WSCenterNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wscnfClient.SubHeartbeat("*", func(m *centernotify.SubHeartbeatResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(500) * time.Second)
	wscnfClient.UnsubHeartbeat("*", "")
	time.Sleep(time.Duration(500) * time.Second)
}
