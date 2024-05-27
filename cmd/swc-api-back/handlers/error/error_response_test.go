package errorweb_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	errorweb "swc-api-back.com/cmd/swc-api-back/handlers/error"
)

func TestBadRequestError(t *testing.T) {
	err := errorweb.BadRequestErr{
		Message: "dummy msg",
	}
	assert.Equal(t, "dummy msg", err.Error())
}

func TestNotFoundError(t *testing.T) {
	err := errorweb.NotFoundErr{
		Message: "dummy msg",
	}
	assert.Equal(t, "dummy msg", err.Error())
}

func TestNewWebError(t *testing.T) {
	type scenario struct {
		name        string
		givenError  error
		givenWriter *httptest.ResponseRecorder
		wantStatus  int
		wantBody    []byte
	}
	testCases := []scenario{
		{
			name: "bad request",
			givenError: errorweb.BadRequestErr{
				Message: "dummy msg",
			},
			givenWriter: httptest.NewRecorder(),
			wantStatus:  http.StatusBadRequest,
			wantBody:    []byte(`{"message":"dummy msg","status_code":400}`),
		},
		{
			name: "not found",
			givenError: errorweb.NotFoundErr{
				Message: "dummy msg 2",
			},
			givenWriter: httptest.NewRecorder(),
			wantStatus:  http.StatusNotFound,
			wantBody:    []byte(`{"message":"dummy msg 2","status_code":404}`),
		},
		{
			name:        "internal server error",
			givenError:  errors.New("dummy msg 3"),
			givenWriter: httptest.NewRecorder(),
			wantStatus:  http.StatusInternalServerError,
			wantBody:    []byte(`{"message":"dummy msg 3","status_code":500}`),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errorweb.NewWebError(tc.givenError, tc.givenWriter)
			assert.Equal(t, tc.wantStatus, tc.givenWriter.Code)
			assert.Equal(t, "application/json", tc.givenWriter.Result().Header.Get("Content-Type"))
			body, err := io.ReadAll(tc.givenWriter.Result().Body)
			require.NoError(t, err)
			assert.Equal(t, tc.wantBody, body)
		})
	}
}
