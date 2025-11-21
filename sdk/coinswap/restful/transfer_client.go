package restful

import (
	"encoding/json"
	"fmt"

	"github.com/bisonai/huobi_futures/sdk/coinswap"
	"github.com/bisonai/huobi_futures/sdk/coinswap/restful/response/transfer"
	"github.com/bisonai/huobi_futures/sdk/log"
	"github.com/bisonai/huobi_futures/sdk/reqbuilder"
)

type TransferClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (tc *TransferClient) Init(accessKey string, secretKey string, host string) *TransferClient {
	if host == "" {
		host = coinswap.SPOT_DEFAULT_HOST
	}
	tc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return tc
}

// AccountTransferAsync (Transfer margin between Spot account and Coin Margined Swap account)
func (tc *TransferClient) AccountTransferAsync(data chan transfer.AccountTransferResponse, from string, to string,
	currency string, amount string) {
	if currency == "" {
		currency = "USDT"
	}

	// ulr
	url := tc.PUrlBuilder.Build(coinswap.POST_METHOD, "/v2/account/transfer", nil)

	// content
	content := fmt.Sprintf("{ \"from\":\"%s\", \"to\":\"%s\", \"currency\":\"%s\", \"amount\":\"%s\"}", from, to, currency, amount)

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
