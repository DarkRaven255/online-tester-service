package responses

import "online-tests/domain/domainmodel"

type TestModel struct {
	Test *domainmodel.Test `json:"test"`
}

func NewTestModelResp(test *domainmodel.Test) *TestModel {
	return &TestModel{
		Test: test,
	}
}
