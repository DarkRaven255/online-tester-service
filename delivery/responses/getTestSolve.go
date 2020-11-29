package responses

import (
	"online-tests/delivery/models/testsolvemodel"
	"online-tests/domain/domainmodel"
)

type TestSolveModel struct {
	Test *testsolvemodel.Test `json:"test"`
}

func NewTestSolveModelResp(test *domainmodel.Test) *TestSolveModel {
	return &TestSolveModel{
		Test: testsolvemodel.NewTestSolveModel(test),
	}
}
