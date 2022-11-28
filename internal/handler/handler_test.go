package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSumHandler(t *testing.T) {
	type args struct {
		body numbersBody
	}
	tests := []struct {
		name     string
		args     args
		wantHttp int
		wantResp float64
	}{
		{
			name: "expected",
			args: args{
				body: numbersBody{
					X: 9,
					Y: 10,
				},
			},
			wantHttp: http.StatusOK,
			wantResp: float64(19),
		},
		{
			name: "internal server error",
			args: args{
				body: numbersBody{
					X: 0,
					Y: 0,
				},
			},
			wantHttp: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			var jsonBody []byte
			if tt.args.body.X == 0 && tt.args.body.Y == 0 {
				jsonBody = []byte{}
			} else {
				jsonBody, _ = json.Marshal(tt.args.body)
			}
			ctx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))

			SumHandler(ctx)

			if w.Code != tt.wantHttp {
				t.Fatalf("status code is not %d but got %d", tt.wantHttp, w.Code)
			}

			if w.Code == http.StatusOK {
				bodyResponse := make(map[string]interface{})
				err := json.Unmarshal(w.Body.Bytes(), &bodyResponse)
				if err != nil {
					t.Fatal("Cannot parse response body", err)
				}

				if bodyResponse["result"].(float64) != tt.wantResp {
					t.Fatalf("Result is not %f, but got %f", tt.wantResp, bodyResponse["result"])
				}
			}
		})
	}
}

func TestMultiplicationHandler(t *testing.T) {
	type args struct {
		body numbersBody
	}
	tests := []struct {
		name     string
		args     args
		wantHttp int
		wantResp float64
	}{
		{
			name: "expected",
			args: args{
				body: numbersBody{
					X: 9,
					Y: 10,
				},
			},
			wantHttp: http.StatusOK,
			wantResp: float64(90),
		},
		{
			name: "internal server error",
			args: args{
				body: numbersBody{
					X: 0,
					Y: 0,
				},
			},
			wantHttp: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			var jsonBody []byte
			if tt.args.body.X == 0 && tt.args.body.Y == 0 {
				jsonBody = []byte{}
			} else {
				jsonBody, _ = json.Marshal(tt.args.body)
			}
			ctx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))

			MultiplicationHandler(ctx)

			if w.Code != tt.wantHttp {
				t.Fatalf("status code is not %d but got %d", tt.wantHttp, w.Code)
			}

			if w.Code == http.StatusOK {
				bodyResponse := make(map[string]interface{})
				err := json.Unmarshal(w.Body.Bytes(), &bodyResponse)
				if err != nil {
					t.Fatal("Cannot parse response body", err)
				}

				if bodyResponse["result"].(float64) != tt.wantResp {
					t.Fatalf("Result is not %f, but got %f", tt.wantResp, bodyResponse["result"])
				}
			}
		})
	}
}

func TestDivideHandler(t *testing.T) {
	type args struct {
		body numbersBody
	}
	tests := []struct {
		name     string
		args     args
		wantHttp int
		wantResp float64
	}{
		{
			name: "expected",
			args: args{
				body: numbersBody{
					X: 10,
					Y: 2,
				},
			},
			wantHttp: http.StatusOK,
			wantResp: float64(5),
		},
		{
			name: "internal server error",
			args: args{
				body: numbersBody{
					X: 0,
					Y: 0,
				},
			},
			wantHttp: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			var jsonBody []byte
			if tt.args.body.X == 0 && tt.args.body.Y == 0 {
				jsonBody = []byte{}
			} else {
				jsonBody, _ = json.Marshal(tt.args.body)
			}
			ctx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonBody))

			DivideHandler(ctx)

			if w.Code != tt.wantHttp {
				t.Fatalf("status code is not %d but got %d", tt.wantHttp, w.Code)
			}

			if w.Code == http.StatusOK {
				bodyResponse := make(map[string]interface{})
				err := json.Unmarshal(w.Body.Bytes(), &bodyResponse)
				if err != nil {
					t.Fatal("Cannot parse response body", err)
				}

				if bodyResponse["result"].(float64) != tt.wantResp {
					t.Fatalf("Result is not %f, but got %f", tt.wantResp, bodyResponse["result"])
				}
			}
		})
	}
}
