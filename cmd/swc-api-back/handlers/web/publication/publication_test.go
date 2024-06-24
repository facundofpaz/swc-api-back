package publication_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"swc-api-back.com/cmd/swc-api-back/handlers/web/publication"
	"swc-api-back.com/internal/mocks"
)

func TestNew(t *testing.T) {
	t.Run("nil publication service", func(t *testing.T) {
		h, err := publication.New(nil)
		assert.Nil(t, h)
		assert.Equal(t, errors.New("missing publication service"), err)
	})
	t.Run("success", func(t *testing.T) {
		h, err := publication.New(&mocks.PublicationServiceMock{})
		assert.NotNil(t, h)
		assert.NoError(t, err)
	})
}

func TestCreate(t *testing.T) {
	type scenario struct {
		name         string
		givenRequest *http.Request
		givenMock    func(context.Context, *mocks.PublicationServiceMock)
		wantBody     []byte
		wantStatus   int
	}
	testCases := []scenario{
		{
			name:         "fail invalid body",
			givenRequest: httptest.NewRequest(http.MethodPost, "/test", nil),
			givenMock:    func(context.Context, *mocks.PublicationServiceMock) {},
			wantStatus:   http.StatusBadRequest,
			wantBody:     []byte(`{"message":"error on decode creation publication body","status_code":400}`),
		},
	}
	ctx := context.TODO()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := tc.givenRequest.WithContext(ctx)
			w := httptest.NewRecorder()
			serviceMock := mocks.NewPublicationServiceMock(t)
			tc.givenMock(ctx, serviceMock)
			h, e := publication.New(serviceMock)
			require.NoError(t, e)
			h.Create(w, r)
			assert.Equal(t, tc.wantStatus, w.Code)
			body, err := io.ReadAll(w.Result().Body)
			require.NoError(t, err)
			assert.Equal(t, tc.wantBody, body)
		})
	}
}
