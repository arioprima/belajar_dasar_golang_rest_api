package web

type WebResponse struct {
	Code   int64       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
