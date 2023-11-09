package api

import (
	"time"

	"github.com/techschool/simplebank/uti"
)

func GetResponse(data interface{}, requestId string, requestTime time.Time) *Response {

	now := time.Now()
	processTime := (now.Sub(requestTime)).Seconds()

	responseStatus := ResponseStatus{
		Code:        "00",
		Type:        "success",
		Message:     "success",
		RequestId:   requestId,
		RequestTime: requestTime.Format(uti.DateTimeLayout),
		ProcessTime: processTime,
	}

	return &Response{ResponseStatus: responseStatus,
		Data: data}
}

func GetResponseError(err error, requestId string, requestTime time.Time) *ResponseStatus {
	now := time.Now()
	processTime := (now.Sub(requestTime)).Seconds()
	var code string = "99"
	mess, ok := uti.MessInputError[err.Error()]
	if !ok {
		mess = err.Error()
	} else {
		code = err.Error()
	}

	responseStatus := ResponseStatus{
		Code:        code,
		Type:        "fail",
		Message:     mess,
		RequestId:   requestId,
		RequestTime: requestTime.Format(uti.DateTimeLayout),
		ProcessTime: processTime,
	}

	return &responseStatus
}