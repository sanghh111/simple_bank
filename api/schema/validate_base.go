package schema

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/techschool/simplebank/uti"
)

// TODO: Custom binding for gin
func (requestInfoParam *RequestInfoParam) ValidateRequestInfo() (RequestInfoValue, error) {
	var requestInfoValue RequestInfoValue

	if requestInfoParam.RequestId == "" {
		return requestInfoValue, errors.New(uti.RequestInfoEmpty)
	}

	if requestInfoParam.LangCode == "" {
		return requestInfoValue, errors.New(uti.LangCodeEmpty)
	}

	if requestInfoParam.RequestTime == "" {
		return requestInfoValue, errors.New(uti.RequestTime)
	}
	err := (&requestInfoValue).LoadInput(requestInfoParam)
	if err != nil {
		return requestInfoValue, err
	}
	return requestInfoValue, nil
}

func (requestInfoValue *RequestInfoValue) LoadInput(requestinfoParam *RequestInfoParam) error {
	var err error
	requestInfoValue.RequestId = requestinfoParam.RequestId
	requestInfoValue.LangCode = requestinfoParam.LangCode
	requestInfoValue.RequestTime, err = uti.StringToDateTime(requestinfoParam.RequestTime)
	if err != nil {
		return err
	}
	return nil
}

// parseBearerAuth parses an HTTP Bearer Authenticate string
// "Bearer 123hkladashdflhafkja.kbnlkabl return 123hkladashdflhafkja.kbnlkabl true"
func ParseBearerAuth(ctx *gin.Context) (token string, ok bool) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		return "", false
	}
	const prefix = "Bearer "
	if (len(auth) <= len(prefix)) && (auth[0:len(prefix)] == prefix) {
		return "", false
	}

	return auth[len(prefix):], true
}
