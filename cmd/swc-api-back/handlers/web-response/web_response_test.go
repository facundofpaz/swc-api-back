package webresponse_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	webresponse "swc-api-back.com/cmd/swc-api-back/handlers/web-response"
	"swc-api-back.com/cmd/swc-api-back/handlers/web/publication/dto"
)

func Test_NewWebResponse(t *testing.T) {
	type scenario struct {
		name        string
		givenBody   any
		givenWriter *httptest.ResponseRecorder
		givenCode   int
	}
	testCases := []scenario{
		{
			name: "response with publication id",
			givenBody: dto.Publication{
				ID: 2,
			},
			givenWriter: httptest.NewRecorder(),
			givenCode:   http.StatusCreated,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			webresponse.NewWebResponse(tc.givenBody, tc.givenWriter, tc.givenCode)
			h := tc.givenWriter.Header().Get("Content-Type")
			assert.Equal(t, "application/json", h)
			assert.Equal(t, tc.givenCode, tc.givenWriter.Code)
			body, err := io.ReadAll(tc.givenWriter.Result().Body)
			require.NoError(t, err)
			givenBody, err := json.Marshal(tc.givenBody)
			require.NoError(t, err)
			assert.Equal(t, givenBody, body)
		})
	}
}
