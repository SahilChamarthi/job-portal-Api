package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"project/internal/middlewear"
	"project/internal/model"
	"project/internal/services"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func Test_handler_companyCreation(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices)
		expectedStatusCode int
		expectedResponse   string
	}{
		//TODO: Add test cases.
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "invalid request body",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := "invalid string request body"
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8082", strings.NewReader(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "checking validator function",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`{"key": "value"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "sucessfully adding company",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`{"company_name": "TekSystems", "company_adress":"Banglore", "domain":"develop"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8082", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := services.NewMockAllinServices(mc)
				ms.EXPECT().CompanyCreate(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"","company_adress":"","domain":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//tt.h.companyCreation(tt.args.c)
			gin.SetMode((gin.TestMode))
			c, rr, ms := tt.setup()
			h := &handler{
				r: ms,
			}
			h.companyCreation(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())

		})
	}
}

func Test_handler_getAllCompany(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error while fetching companies",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := services.NewMockAllinServices(mc)
				ms.EXPECT().GetAllCompanies().Return([]model.Company{}, errors.New("companies not found")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "sucessfully fetching companies",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq := httptest.NewRequest(http.MethodGet, "http://google.com:8082", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := services.NewMockAllinServices(mc)
				ms.EXPECT().GetAllCompanies().Return([]model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				r: ms,
			}
			h.getAllCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getCompany(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices)
		expectedStatusCode int
		expectedResponse   string
	}{
		// {
		// 	name: "missing trace id",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
		// 		c.Request = httpReq

		// 		return c, rr, nil
		// 	},
		// 	expectedStatusCode: http.StatusInternalServerError,
		// 	expectedResponse:   `{"error":"Internal Server Error"}`,
		// },
		// {
		// 	name: "Invalid companyId",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8082", nil)
		// 		ctx := httpReq.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "693")
		// 		httpReq = httpReq.WithContext(ctx)
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "one"})
		// 		c.Request = httpReq

		// 		return c, rr, nil
		// 	},
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedResponse:   `{"error":"Bad Request"}`,
		// // },
		// {
		// 	name: "error while fetching company details by companyId",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8082", nil)
		// 		ctx := httpReq.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "693")
		// 		httpReq = httpReq.WithContext(ctx)
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		// 		c.Request = httpReq
		// 		mc := gomock.NewController(t)
		// 		ms := services.NewMockAllinServices(mc)
		// 		ms.EXPECT().GetCompanyById(gomock.Any()).Return(model.Company{}, errors.New("error while fetching company")).AnyTimes()

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusInternalServerError,
		// 	expectedResponse:   `{"error":"Internal Server Error"}`,
		// },
		{
			name: "sucess while fetching company details by companyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.AllinServices) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq := httptest.NewRequest(http.MethodGet, "http://google.com:8082", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "693"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := services.NewMockAllinServices(mc)
				ms.EXPECT().GetCompanyById(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"","company_adress":"","domain":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				r: ms,
			}
			h.getCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
