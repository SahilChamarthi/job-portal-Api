package handlers

import (
	"bytes"
	"context"
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
		// TODO: Add test cases.
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
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", strings.NewReader(requestBody))
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
				requestBody := []byte(`{"CompanyName": "TekSystems", "Adress":"Banglore", "Domain":"develop"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := services.NewMockAllinServices(mc)
				ms.EXPECT().CompanyCreate(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

				return c, rr, nil
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"CompanyName":"","Adress":"","Domain":""}`,
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
