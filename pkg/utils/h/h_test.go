// Copyright 2019 Shanghai JingDuo Information Technology co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package h

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponseData_Error(t *testing.T) {

	tests := []struct {
		Input *ResponseData
		Want  string
	}{
		{
			Input: &ResponseData{
				Code:    http.StatusBadRequest,
				Message: "BadRequest",
			},
			Want: `{"code":400,"msg":"BadRequest"}`,
		},
		{
			Input: &ResponseData{
				Code:    http.StatusBadRequest,
				Message: "Parameters Error",
			},
			Want: `{"code":400,"msg":"Parameters Error"}`,
		},
		{
			Input: &ResponseData{Data: make(chan int)},
			Want:  `{"code":500,"msg":"Unknown"}`,
		},
	}

	for _, item := range tests {
		assert.Equal(t, item.Want, item.Input.Error())
	}
}

func TestResponseData_String(t *testing.T) {

	tests := []struct {
		Input *ResponseData
		Want  string
	}{
		{
			Input: &ResponseData{
				Code:    http.StatusBadRequest,
				Message: "BadRequest",
			},
			Want: `{"code":400,"msg":"BadRequest"}`,
		},
		{
			Input: &ResponseData{
				Code:    http.StatusBadRequest,
				Message: "Parameters Error",
			},
			Want: `{"code":400,"msg":"Parameters Error"}`,
		},
		{
			Input: &ResponseData{Data: make(chan int)},
			Want:  ``,
		},
	}

	for _, item := range tests {
		assert.Equal(t, item.Want, item.Input.String())
	}
}

func TestResponseData_WithCode(t *testing.T) {

	tests := []struct {
		Input     ResponseData
		InputCode int
		Want      *ResponseData
	}{
		{
			Input: ResponseData{
				Code:    http.StatusBadRequest,
				Message: "BadRequest",
			},
			InputCode: 500,
			Want: &ResponseData{
				Code:    500,
				Message: "BadRequest",
			},
		},
	}

	for _, item := range tests {
		assert.Equal(t, item.Want, item.Input.WithCode(item.InputCode))
	}
}

func TestResponseData_WithError(t *testing.T) {

	tests := []struct {
		Input      ResponseData
		InputError error
		Want       *ResponseData
	}{
		{
			Input: ResponseData{
				Code:    http.StatusBadRequest,
				Message: "BadRequest",
			},
			InputError: errors.New("BadGateway"),
			Want: &ResponseData{
				Code:    http.StatusBadRequest,
				Message: "BadGateway",
			},
		},
	}

	for _, item := range tests {
		assert.Equal(t, item.Want, item.Input.WithError(item.InputError))
	}
}

func TestE(t *testing.T) {

	tests := []struct {
		Input error
		Want  struct {
			Code           int
			ResponseString string
		}
	}{
		{
			Input: ENotFound,
			Want: struct {
				Code           int
				ResponseString string
			}{Code: ENotFound.Status, ResponseString: `{"code":404,"msg":"NotFound"}` + "\n"},
		},
		{
			Input: EParamsError,
			Want: struct {
				Code           int
				ResponseString string
			}{Code: EParamsError.Status, ResponseString: `{"code":400,"msg":"ParamsError"}` + "\n"},
		},
		{
			Input: errors.New("BadRequest"),
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusInternalServerError, ResponseString: `{"code":500,"msg":"BadRequest"}` + "\n"},
		},
	}

	for _, item := range tests {

		resp := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(resp)
		E(ctx, http.StatusInternalServerError, item.Want.Code, item.Want.ResponseString)
		resp.Flush()
	}
}

func TestData(t *testing.T) {

	tests := []struct {
		Input struct {
			Body       interface{}
			HTTPMethod string
		}
		Want struct {
			Code           int
			ResponseString string
		}
	}{
		{
			Input: struct {
				Body       interface{}
				HTTPMethod string
			}{
				Body:       struct{ Hello string }{Hello: "World"},
				HTTPMethod: http.MethodGet,
			},
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusOK, ResponseString: `{"code":0,"data":{"Hello":"World"}}` + "\n"},
		},
		{
			Input: struct {
				Body       interface{}
				HTTPMethod string
			}{
				Body:       struct{ Hello string }{Hello: "World"},
				HTTPMethod: http.MethodPost,
			},
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusCreated, ResponseString: `{"code":0,"data":{"Hello":"World"}}` + "\n"},
		},
		{
			Input: struct {
				Body       interface{}
				HTTPMethod string
			}{
				Body:       struct{ Hello string }{Hello: "World"},
				HTTPMethod: http.MethodDelete,
			},
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusNoContent, ResponseString: ``},
		},
	}

	for _, item := range tests {

		resp := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(resp)
		ctx.Request = &http.Request{Method: item.Input.HTTPMethod}
		Data(ctx, item.Input.Body)
		resp.Flush()

		assert.Equal(t, item.Want.Code, resp.Code)
		assert.Equal(t, item.Want.ResponseString, resp.Body.String())
	}
}

func TestR(t *testing.T) {

	tests := []struct {
		Input struct {
			Body       interface{}
			HTTPMethod string
		}
		Want struct {
			Code           int
			ResponseString string
		}
	}{
		{
			Input: struct {
				Body       interface{}
				HTTPMethod string
			}{
				Body:       struct{ Hello string }{Hello: "World"},
				HTTPMethod: http.MethodGet,
			},
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusOK, ResponseString: `{"Hello":"World"}` + "\n"},
		},
		{
			Input: struct {
				Body       interface{}
				HTTPMethod string
			}{
				Body:       struct{ Hello string }{Hello: "World"},
				HTTPMethod: http.MethodPost,
			},
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusCreated, ResponseString: `{"Hello":"World"}` + "\n"},
		},
		{
			Input: struct {
				Body       interface{}
				HTTPMethod string
			}{
				Body:       struct{ Hello string }{Hello: "World"},
				HTTPMethod: http.MethodDelete,
			},
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusNoContent, ResponseString: ``},
		},
	}

	for _, item := range tests {

		resp := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(resp)
		ctx.Request = &http.Request{Method: item.Input.HTTPMethod}
		R(ctx, item.Input.Body)
		resp.Flush()

		assert.Equal(t, item.Want.Code, resp.Code)
		assert.Equal(t, item.Want.ResponseString, resp.Body.String())
	}
}

func TestRJsonP(t *testing.T) {

	tests := []struct {
		Input struct {
			Body       interface{}
			URL        string
			HTTPMethod string
		}
		Want struct {
			Code           int
			ResponseString string
		}
	}{
		{
			Input: struct {
				Body       interface{}
				URL        string
				HTTPMethod string
			}{
				Body:       struct{ Hello string }{Hello: "World"},
				URL:        "http://localhost/?callback=c",
				HTTPMethod: http.MethodGet,
			},
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusOK, ResponseString: `c({"Hello":"World"});`},
		},
		{
			Input: struct {
				Body       interface{}
				URL        string
				HTTPMethod string
			}{
				Body:       struct{ Hello string }{Hello: "World"},
				URL:        "http://localhost/?callback=c",
				HTTPMethod: http.MethodPost,
			},
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusCreated, ResponseString: `c({"Hello":"World"});`},
		},
		{
			Input: struct {
				Body       interface{}
				URL        string
				HTTPMethod string
			}{
				Body:       struct{ Hello string }{Hello: "World"},
				URL:        "http://localhost/?callback=c",
				HTTPMethod: http.MethodDelete,
			},
			Want: struct {
				Code           int
				ResponseString string
			}{Code: http.StatusNoContent, ResponseString: ``},
		},
	}

	for _, item := range tests {

		resp := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(resp)
		requestURL, err := url.Parse(item.Input.URL)
		assert.Nil(t, err)
		ctx.Request = &http.Request{Method: item.Input.HTTPMethod, URL: requestURL}
		RJsonP(ctx, item.Input.Body)
		resp.Flush()

		assert.Equal(t, item.Want.Code, resp.Code)
		assert.Equal(t, item.Want.ResponseString, resp.Body.String())
	}
}
