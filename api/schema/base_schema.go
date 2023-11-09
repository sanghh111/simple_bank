package api

import (
	"time"
)

type ResponseStatus struct {
	Code        string  `json:"code"`
	Type        string  `json:"type"`
	Message     string  `json:"message"`
	RequestId   string  `json:"requestId"`
	RequestTime string  `json:"requestTime"`
	ProcessTime float64 `json:"processTime"`
}

type RequestInfoParam struct {
	RequestId   string `json:"requestId"`
	RequestTime string `json:"requestTime"`
	LangCode    string `json:"langCode"`
}

type RequestInfoValue struct {
	RequestId   string `json:"requestId"`
	RequestTime time.Time
	LangCode    string `json:"langCode"`
}

type Response struct {
	Data           interface{}    `json:"data"`
	ResponseStatus ResponseStatus `json:"responseStatus"`
}
