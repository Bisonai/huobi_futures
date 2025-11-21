package main

import (
	"github.com/bisonai/huobi_futures/test/coinswap_test/accountclienttest"
	"github.com/bisonai/huobi_futures/test/coinswap_test/commonclienttest"
	"github.com/bisonai/huobi_futures/test/coinswap_test/marketclienttest"
	"github.com/bisonai/huobi_futures/test/coinswap_test/orderclienttest"
	"github.com/bisonai/huobi_futures/test/coinswap_test/transferclienttest"
	"github.com/bisonai/huobi_futures/test/coinswap_test/triggerorderclienttest"
	"github.com/bisonai/huobi_futures/test/coinswap_test/wscenternotifyclienttest"
	"github.com/bisonai/huobi_futures/test/coinswap_test/wsindexclienttest"
	"github.com/bisonai/huobi_futures/test/coinswap_test/wsmarketclienttest"
	"github.com/bisonai/huobi_futures/test/coinswap_test/wsnotifyclienttest"
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
	wsindexclienttest.RunAllExamples()
	wsmarketclienttest.RunAllExamples()
	wsnotifyclienttest.RunAllExamples()
	wscenternotifyclienttest.RunAllExamples()
}
