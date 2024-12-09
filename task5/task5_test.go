package task5

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {

	testsCases := []struct {
		Name             string
		InputBody        string
		ExpectStatusCode int
		ExpectResponse   string
	}{
		{
			"Valid_input=Rustam",
			"Rustam",
			200,
			"Hello, Rustam!",
		},
		{
			"Empty_input",
			"",
			400,
			"",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {
			r := strings.NewReader(tt.InputBody)
			req := httptest.NewRequest("GET", "/", r)
			w := httptest.NewRecorder()

			helloHandler(w, req)

			resp := w.Result()
			assert.Equal(t, tt.ExpectStatusCode, resp.StatusCode, "expected status code")

			readBodyRaw, _ := io.ReadAll(resp.Body)

			assert.Equal(t, tt.ExpectResponse, string(readBodyRaw), "expected body")
		})
	}
}
