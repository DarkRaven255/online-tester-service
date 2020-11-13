package response

import "online-tests/delivery/model"

type GetTestResp struct {
	Test model.Test `json:"test"`
}
