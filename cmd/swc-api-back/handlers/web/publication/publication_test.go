package publication_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"swc-api-back.com/cmd/swc-api-back/handlers/web/publication"
	publicationDomain "swc-api-back.com/internal/domain/model/publication"
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
		{
			name:         "fail publication service error",
			givenRequest: httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(getValidBody())),
			givenMock: func(_ context.Context, serviceMock *mocks.PublicationServiceMock) {
				serviceMock.On("Create", mock.Anything, mock.Anything).Return(publicationDomain.ID(0), errors.New("fail create"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   []byte(`{"message":"fail create","status_code":500}`),
		},
		{
			name:         "success",
			givenRequest: httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(getValidBody())),
			givenMock: func(ctx context.Context, serviceMock *mocks.PublicationServiceMock) {
				serviceMock.On("Create", ctx, getValidPublication()).Return(publicationDomain.ID(1), nil)
			},
			wantStatus: http.StatusCreated,
			wantBody:   []byte(`{"id":1}`),
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

func getValidPublication() publicationDomain.Publication {
	return publicationDomain.Publication{
		Latitude:     35.4,
		Longitude:    24.1,
		StreetName:   "avda siempre viva",
		StreetNumber: "1548",
		Comment:      "simpsons family",
		OfficesCount: 5,
		Name:         "The Simpsons",
		WorkingHours: map[int][]string{
			1: {"8 a 12", "16 a 20"},
			2: {"9 a 13", "17 a 21"},
		},
		Tags: []string{"coffe", "parking"},
	}
}

func getValidBody() string {
	return `{
				"latitude": 35.4,
				"longitud": 24.1,
				"street_name": "avda siempre viva",
				"street_number": "1548",
				"comment": "simpsons family",
				"offices_count": 5,
				"name": "The Simpsons",
				"working_hours": {
					"1": [
						"8 a 12",
						"16 a 20"
					],
					"2": [
						"9 a 13",
						"17 a 21"
					]
				},
				"tags": [
					"coffe",
					"parking"
				]
			}`
}
