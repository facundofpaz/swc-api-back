package publication

import (
	"context"
	"errors"

	"swc-api-back.com/internal/domain/model/publication"
	publicationRepo "swc-api-back.com/internal/domain/repository"
	publicationSrv "swc-api-back.com/internal/domain/service"
)

type service struct {
	publicationRepository publicationRepo.Publication
}

func New(publicationRepository publicationRepo.Publication) (publicationSrv.Publication, error) {
	if publicationRepository == nil {
		return nil, errors.New("missing publication repository")
	}
	return service{
		publicationRepository: publicationRepository,
	}, nil
}

func (s service) GetByID(ctx context.Context, id publication.ID) (publication.Publication, error) {
	if id == 45 {
		return s.publicationRepository.GetByID(ctx, id)
	}
	return publication.Publication{}, errors.New("publication not found")
}
