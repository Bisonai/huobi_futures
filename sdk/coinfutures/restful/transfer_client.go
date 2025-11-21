package restful

import (
	"encoding/json"
	"fmt"

	"github.com/bisonai/huobi_futures/sdk/coinfutures"
	"github.com/bisonai/huobi_futures/sdk/coinfutures/restful/response/transfer"
	"github.com/bisonai/huobi_futures/sdk/log"
	"github.com/bisonai/huobi_futures/sdk/reqbuilder"
)

type TransferClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (tc *TransferClient) Init(accessKey string, secretKey string, host string) *TransferClient {
	if host == "" {
		host = coinfutures.SPOT_DEFAULT_HOST
	}
	tc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return tc
}

// FuturesTransferAsync (Transfer margin between Spot account and Future account)
func (tc *TransferClient) FuturesTransferAsync(data chan transfer.FuturesTransferResponse, amount string, currency string, fcType string) {
	if currency == "" {
		currency = "USDT"
	}

	// ulr
	url := tc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/v1/futures/transfer", nil)

	// content
	content := fmt.Sprintf("{ \"amount\":\"%s\", \"type\":\"%s\", \"currency\":\"%s\"}", amount, currency, fcType)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := transfer.FuturesTransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to FuturesTransferResponse error: %s", getErr)
	}
	data <- result
}

// AccountTransferAsync ([General] Spot - transfer funds between contract accounts)
func (tc *TransferClient) AccountTransferAsync(data chan transfer.AccountTransferResponse, from string, to string, amount float32, marginAccount string, currency string) {
	if currency == "" {
		currency = "USDT"
	}

	// ulr
	url := tc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/v2/account/transfer", nil)

	// content
	content := fmt.Sprintf("{ \"from\":\"%s\", \"to\":\"%s\", \"currency\":\"%s\", \"amount\":%f, \"margin-account\":\"%s\" }", from, to, currency, amount, marginAccount)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := transfer.AccountTransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to AccountTransferResponse error: %s", getErr)
	}
	data <- result
}
