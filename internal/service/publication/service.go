package publication

import (
	"context"
	"errors"

	publicationRepo "swc-api-back.com/internal/repository/publication"

	"swc-api-back.com/internal/domain/publication"
)

type service struct {
	publicationRepository publicationRepo.Publication
}

func New(publicationRepository publicationRepo.Publication) (Publication, error) {
	if publicationRepository == nil {
		return nil, errors.New("missing publication repository")
	}
	return service{
		publicationRepository: publicationRepository,
	}, nil
}

func (s service) GetByID(ctx context.Context, id publication.ID) (publication.Publication, error) {
	return s.publicationRepository.GetByID(ctx, id)
}

func (s service) Create(ctx context.Context, p publication.Publication) (publication.ID, error) {
	return s.publicationRepository.Create(ctx, p)
}
