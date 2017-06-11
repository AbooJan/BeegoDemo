package models

// BaseResponse is All Response base
type BaseResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}
