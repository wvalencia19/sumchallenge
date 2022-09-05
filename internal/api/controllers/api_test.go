package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sum/internal/api"
	"sum/internal/api/middlewares"
	"sum/internal/models"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var secretKey = "secret"

func TestAuthentication(t *testing.T) {
	tests := []struct {
		name               string
		params             models.User
		wantHTTPStatusCode int
	}{
		{
			name: "valid data",
			params: models.User{
				Username: "wilson",
				Password: "val",
			},
			wantHTTPStatusCode: http.StatusOK,
		},
		{
			name: "missing username",
			params: models.User{
				Username: "",
				Password: "val",
			},
			wantHTTPStatusCode: http.StatusUnauthorized,
		},
		{
			name: "missing password",
			params: models.User{
				Username: "wilson",
				Password: "",
			},
			wantHTTPStatusCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			w := httptest.NewRecorder()
			r.POST("/auth", Auth(1*time.Hour, secretKey))
			payload := getAuthPayload(tt.params)

			req, err := http.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(payload))
			if err != nil {
				t.Fail()
			}

			req.Header.Add("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			if w.Code != tt.wantHTTPStatusCode {
				t.Fail()
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name               string
		payload            []byte
		token              string
		wantHTTPStatusCode int
		result             string
	}{
		{
			name:               "with valid token",
			payload:            []byte(`{"data": [1,2,3,4]}`),
			token:              createToken(10 * time.Minute),
			wantHTTPStatusCode: http.StatusOK,
			result:             "4a44dc15364204a80fe80e9039455cc1608281820fe2b24f1e5233ade6af1dd5",
		},
		{
			name:               "with token expired",
			payload:            []byte(`{"data": [1,2,3,4]}`),
			token:              createToken(1 * time.Millisecond),
			wantHTTPStatusCode: http.StatusUnauthorized,
			result:             "",
		},
		{
			name:               "with invalid json",
			payload:            []byte(`{"data": [1,2,3,4]`),
			token:              createToken(10 * time.Minute),
			wantHTTPStatusCode: http.StatusBadRequest,
			result:             "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			r := gin.Default()

			r.POST("/sum", middlewares.JwtAuthMiddleware(secretKey), Sum)
			time.Sleep(1 * time.Second)
			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodPost, "/sum", bytes.NewBuffer(tt.payload))
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+tt.token)

			r.ServeHTTP(w, req)

			response := struct {
				Result string `json:"result"`
			}{}
			_ = json.NewDecoder(w.Body).Decode(&response)
			if w.Code != tt.wantHTTPStatusCode {
				t.Fail()
			}

			assert.Equal(t, response.Result, tt.result)
		})
	}
}

func getAuthPayload(user models.User) []byte {
	b, _ := json.Marshal(user)
	return b
}

func createToken(ttl time.Duration) string {
	token, _ := api.GenerateToken("usernamne", secretKey, ttl)
	return token
}
