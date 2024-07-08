package publication

import (
	"context"

	"swc-api-back.com/internal/domain/publication"
)

//go:generate mockery --name=Publication --structname=PublicationRepositoryMock --filename=publication_repository.go --output=../../mocks
type Publication interface {
	GetByID(context.Context, publication.ID) (publication.Publication, error)
	Create(context.Context, publication.Publication) (publication.ID, error)
}
