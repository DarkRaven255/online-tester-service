package response

import "online-tests/domain/domainmodel"

type GetTestResp struct {
	Test *domainmodel.Test `json:"test"`
}

func NewGetTestResponse(test *domainmodel.Test) *GetTestResp {
	return &GetTestResp{
		Test: test,
	}
}
