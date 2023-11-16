package schema

import (
	"errors"

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
