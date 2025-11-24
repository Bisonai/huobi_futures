package account

type AssetModeResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data AssetMode `json:"data"`
	Ts   int64     `json:"ts"`
}

type AssetMode struct {
	AssetMode int `json:"assets_mode"`
}
