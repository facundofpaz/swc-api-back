package publication

import (
	"context"

	"swc-api-back.com/internal/domain/model/publication"
)

type repository struct {
}

func New() (Publication, error) {
	return repository{}, nil

}

func (r repository) GetByID(ctx context.Context, id publication.ID) (publication.Publication, error) {
	return publication.Publication{
		ID:          45,
		Description: "is a big description",
	}, nil
}
