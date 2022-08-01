package models

type Fail struct {
	Code int `json:"code" example:"500"`
	Msg string `json:"msg" example:"error"`
}

type Success struct {
	Code int `json:"code" example:"200"`
	Msg string `json:"msg" example:"ok"`
}

type Empty struct {
}