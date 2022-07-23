package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		wantHttp int
		wantResp string
	}{
		{
			name:     "expected",
			wantHttp: http.StatusOK,
			wantResp: "hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = httptest.NewRequest("GET", "/", nil)

			HelloWorld(ctx)

			if w.Code != tt.wantHttp {
				t.Fatalf("status code is not %d but got %d", tt.wantHttp, w.Code)
			}

			if w.Code == http.StatusOK {
				bodyResponse := make(map[string]interface{})
				err := json.Unmarshal(w.Body.Bytes(), &bodyResponse)
				if err != nil {
					t.Fatal("Cannot parse response body", err)
				}

				if bodyResponse["message"] != tt.wantResp {
					t.Fatalf("Result is not %s, but got %s", tt.wantResp, bodyResponse["message"])
				}
			}
		})
	}
}
