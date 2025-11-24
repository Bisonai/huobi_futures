package account

type GetAccountPositionResponse struct {
	Code int `json:"code"`

	Message string `json:"message"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data GetAccountPositionResponseData `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetAccountPositionResponseData struct {
	State                 string                             `json:"state"`
	Equity                string                             `json:"equity"`
	InitialMargin         string                             `json:"initial_margin"`
	MaintenanceMargin     string                             `json:"maintenance_margin"`
	MaintenanceMarginRate string                             `json:"maintenance_margin_rate"`
	ProfitUnreal          string                             `json:"profit_unreal"`
	AvailableMargin       string                             `json:"available_margin"`
	VoucherValue          string                             `json:"voucher_value"`
	CreatedTime           int64                              `json:"created_time"`
	UpdatedTime           int64                              `json:"updated_time"`
	Details               []GetAccountPositionResponseDetail `json:"details,omitempty"`
}

type GetAccountPositionResponseDetail struct {
	Symbol string `json:"symbol"`

	ContractCode string `json:"contract_code"`

	Volume float32 `json:"volume"`

	Available float32 `json:"available"`

	Frozen float32 `json:"frozen"`

	CostOpen float32 `json:"cost_open"`

	CostHold float32 `json:"cost_hold"`

	ProfitUnreal float32 `json:"profit_unreal"`

	ProfitRate float32 `json:"profit_rate"`

	Profit float32 `json:"profit"`

	MarginAsset string `json:"margin_asset"`

	PositionMargin float32 `json:"position_margin"`

	LeverRate int `json:"lever_rate"`

	Direction string `json:"direction"`

	LastPrice float32 `json:"last_price"`

	MarginMode string `json:"margin_mode"`

	MarginAccount string `json:"margin_account"`

	AdlRiskPercent string `json:"adl_risk_percent"`

	WithdrawAvailable string `json:"withdraw_available"`

	ContractType string `json:"contract_type"`

	Pair string `json:"pair"`

	BusinessType string `json:"business_type"`

	PositionMode string `json:"position_mode"`

	LiquidationPrice string `json:"liquidation_price"`
}
