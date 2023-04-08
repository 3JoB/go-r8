package model

type View struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}
