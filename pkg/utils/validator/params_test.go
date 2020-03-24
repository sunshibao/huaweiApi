package validator_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"huaweiApi/pkg/utils/h"
	"huaweiApi/pkg/utils/validator"
)

type TestParamsStructure struct {
	Key string `json:"key"`
}

func (s *TestParamsStructure) Validate() error {

	return validator.NewWrapper(
		validator.ValidateString(s.Key, "key", validator.ItemNotEmptyLimit, validator.ItemNoLimit),
	).Validate()
}

func TestParams(t *testing.T) {

	tests := []struct {
		Input string
		Want  TestParamsStructure
		Error error
	}{
		{
			Input: `{"key":"a"}`,
			Want:  TestParamsStructure{Key: "a"},
			Error: nil,
		},
		{
			Input: `"sss"`,
			Want:  TestParamsStructure{},
			Error: h.EBindBodyError.WithMessage("json: cannot unmarshal string into Go value of type validator_test.TestParamsStructure"),
		},
		{
			Input: `{"key":""}`,
			Want:  TestParamsStructure{},
			Error: h.EParamsError.WithMessage("\"key\" '' is too short"),
		},
	}

	for _, test := range tests {

		var err error
		resp := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(resp)
		bodyReader := strings.NewReader(test.Input)
		ctx.Request = httptest.NewRequest("POST", "/test", bodyReader)

		data := &TestParamsStructure{}
		err = validator.Params(ctx, data)

		assert.Equal(t, test.Error, err)
		assert.Equal(t, test.Want, *data)
	}
}
