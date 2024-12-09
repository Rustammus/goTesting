package task10

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMuxer(t *testing.T) {
	testsCases := []struct {
		Name         string
		Path         string
		ExpectStatus int
	}{
		{
			Name:         "Ok",
			Path:         "/ok",
			ExpectStatus: http.StatusOK,
		},
		{
			Name:         "ServerErr",
			Path:         "/error",
			ExpectStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {
			ts := httptest.NewServer(NewMuxer())

			defer ts.Close()

			resp, err := http.Get(ts.URL + tt.Path)
			assert.NoError(t, err, "response err")
			assert.Equal(t, tt.ExpectStatus, resp.StatusCode, "status code mismatch")
		})
	}
}
