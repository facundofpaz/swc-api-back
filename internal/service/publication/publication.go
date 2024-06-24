package publication

import (
	"context"

	"swc-api-back.com/internal/domain/model/publication"
)

//go:generate mockery --name=Publication --structname=PublicationServiceMock --filename=publication_service.go --output=../../mocks

type Publication interface {
	GetByID(context.Context, publication.ID) (publication.Publication, error)
	Create(context.Context, publication.Publication) (publication.ID, error)
}
