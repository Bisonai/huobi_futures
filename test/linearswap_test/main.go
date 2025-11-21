package main

import (
	"github.com/bisonai/huobi_futures/test/linearswap_test/accountclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/commonclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/marketclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/orderclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/transferclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/triggerorderclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/unifiedaccountclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/wscenternotifyclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/wsindexclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/wsmarketclienttest"
	"github.com/bisonai/huobi_futures/test/linearswap_test/wsnotifyclienttest"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	accountclienttest.RunAllExamples()
	commonclienttest.RunAllExamples()
	marketclienttest.RunAllExamples()
	orderclienttest.RunAllExamples()
	transferclienttest.RunAllExamples()
	triggerorderclienttest.RunAllExamples()
	unifiedaccountclienttest.RunAllExamples()
	wsindexclienttest.RunAllExamples()
	wsmarketclienttest.RunAllExamples()
	wsnotifyclienttest.RunAllExamples()
	wscenternotifyclienttest.RunAllExamples()
}
