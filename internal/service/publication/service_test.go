package publication_test

import (
	"context"
	"errors"
	"testing"

	publicationDomain "swc-api-back.com/internal/domain/publication"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"swc-api-back.com/internal/mocks"
	"swc-api-back.com/internal/service/publication"
)

func TestNew(t *testing.T) {
	t.Run("nil repository publication", func(t *testing.T) {
		s, err := publication.New(nil)
		assert.Nil(t, s)
		assert.Equal(t, errors.New("missing publication repository"), err)
	})
	t.Run("success", func(t *testing.T) {
		s, err := publication.New(&mocks.PublicationRepositoryMock{})
		assert.Nil(t, err)
		assert.NotNil(t, s)
	})
}

func TestGetByID(t *testing.T) {
	ctx := context.TODO()
	testCases := []struct {
		name            string
		givenID         publicationDomain.ID
		givenMocks      func(context.Context, *mocks.PublicationRepositoryMock)
		wantPublication publicationDomain.Publication
		wantError       error
	}{
		{
			name:    "fail",
			givenID: 0,
			givenMocks: func(_ context.Context, prm *mocks.PublicationRepositoryMock) {
				prm.On("GetByID", mock.Anything, mock.Anything).Return(publicationDomain.Publication{}, errors.New("error"))
			},
			wantError: errors.New("error"),
		},
		{
			name:    "success",
			givenID: 1,
			givenMocks: func(c context.Context, prm *mocks.PublicationRepositoryMock) {
				prm.On("GetByID", c, publicationDomain.ID(1)).Return(publicationDomain.Publication{Name: "a"}, nil)
			},
			wantPublication: publicationDomain.Publication{Name: "a"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			prm := mocks.NewPublicationRepositoryMock(t)
			tC.givenMocks(ctx, prm)
			s, err := publication.New(prm)
			require.NoError(t, err)
			got, gotErr := s.GetByID(ctx, tC.givenID)
			assert.Equal(t, tC.wantPublication, got)
			assert.Equal(t, tC.wantError, gotErr)
		})
	}
}

func TestCreate(t *testing.T) {
	ctx := context.TODO()
	testCases := []struct {
		name              string
		givenPublication  publicationDomain.Publication
		givenMocks        func(context.Context, *mocks.PublicationRepositoryMock)
		wantPublicationID publicationDomain.ID
		wantError         error
	}{
		{
			name:             "fail",
			givenPublication: publicationDomain.Publication{Name: "a"},
			givenMocks: func(_ context.Context, prm *mocks.PublicationRepositoryMock) {
				prm.On("Create", mock.Anything, mock.Anything).Return(publicationDomain.ID(0), errors.New("error"))
			},
			wantError: errors.New("error"),
		},
		{
			name:             "success",
			givenPublication: publicationDomain.Publication{Name: "a"},
			givenMocks: func(c context.Context, prm *mocks.PublicationRepositoryMock) {
				prm.On("Create", c, publicationDomain.Publication{Name: "a"}).Return(publicationDomain.ID(4), nil)
			},
			wantPublicationID: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			prm := mocks.NewPublicationRepositoryMock(t)
			tC.givenMocks(ctx, prm)
			s, err := publication.New(prm)
			require.NoError(t, err)
			got, gotErr := s.Create(ctx, tC.givenPublication)
			assert.Equal(t, tC.wantPublicationID, got)
			assert.Equal(t, tC.wantError, gotErr)
		})
	}
}
