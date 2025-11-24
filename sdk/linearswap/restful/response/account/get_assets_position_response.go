package account

type GetAssetsPositionResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data GetAssetsPositionResponseData `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}

type GetAssetsPositionResponseData struct {
	State                 string                            `json:"state"`
	Equity                string                            `json:"equity"`
	Details               []GetAssetsPositionResponseDetail `json:"details,omitempty"`
	InitialMargin         string                            `json:"initial_margin"`
	MaintenanceMargin     string                            `json:"maintenance_margin"`
	MaintenanceMarginRate string                            `json:"maintenance_margin_rate"`
	ProfitUnreal          string                            `json:"profit_unreal"`
	AvailableMargin       string                            `json:"available_margin"`
	VoucherValue          string                            `json:"voucher_value"`
	CreatedTime           int64                             `json:"created_time"`
	UpdatedTime           int64                             `json:"updated_time"`
}

type GetAssetsPositionResponseDetail struct {
	Symbol string `json:"symbol,omitempty"`

	ContractCode string `json:"contract_code,omitempty"`

	MarginAsset string `json:"margin_asset"`

	MarginBalance float32 `json:"margin_balance"`

	MarginStatic float32 `json:"margin_static"`

	MarginPosition float32 `json:"margin_position"`

	MarginFrozen float32 `json:"margin_frozen"`

	MarginAvailable float32 `json:"margin_available,omitempty"`

	ProfitReal float32 `json:"profit_real"`

	ProfitUnreal float32 `json:"profit_unreal"`

	WithdrawAvailable float32 `json:"withdraw_available,omitempty"`

	RiskRate float32 `json:"risk_rate,omitempty"`

	LiquidationPrice float32 `json:"liquidation_price,omitempty"`

	LeverRate float32 `json:"lever_rate,omitempty"`

	AdjustFactor float32 `json:"adjust_factor,omitempty"`

	MarginMode string `json:"margin_mode"`

	MarginAccount string `json:"margin_account"`

	NewRiskRate string `json:"new_risk_rate"`

	PositionMode string `json:"position_mode"`

	TradePartition string `json:"trade_partition"`

	AdlRiskPercent string `json:"adl_risk_percent"`

	ContractDetail []GetAssetPositionResponseContractDetail `json:"contract_detail,omitempty"`

	Positions []GetAssetPositionResponsePositions `json:"positions,omitempty"`
}

type GetAssetPositionResponseContractDetail struct {
	Symbol string `json:"symbol"`

	ContractCode string `json:"contract_code"`

	MarginPosition float32 `json:"margin_position"`

	MarginFrozen float32 `json:"margin_frozen"`

	MarginAvailable float32 `json:"margin_available"`

	ProfitUnreal float32 `json:"profit_unreal"`

	LiquidationPrice float32 `json:"liquidation_price"`

	LeverRate float32 `json:"lever_rate"`

	AdjustFactor float32 `json:"adjust_factor"`
}

type GetAssetPositionResponsePositions struct {
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

	NewRiskRate string `json:"new_risk_rate"`

	TradePartition string `json:"trade_partition"`

	PositionMode string `json:"position_mode"`

	AdlRiskPercent string `json:"adl_risk_percent"`
}
