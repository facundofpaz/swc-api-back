package publication

import (
	"context"
	"errors"

	"swc-api-back.com/internal/domain/model/publication"
	db "swc-api-back.com/internal/infraestructure/db"
	"swc-api-back.com/internal/repository/publication/dto"
)

type repository struct {
	dbClient db.Client
}

func New(dbClient db.Client) (Publication, error) {
	if dbClient == nil {
		return nil, errors.New("missing db client")
	}
	return repository{dbClient: dbClient}, nil

}

func (r repository) GetByID(ctx context.Context, id publication.ID) (publication.Publication, error) {
	resp := make([]dto.PublicationDB, 0)
	query := "SELECT * FROM publication WHERE id=?"

	if err := r.dbClient.Select(ctx, &resp, query, id); err != nil {
		return publication.Publication{}, nil
	}
	publications := dto.ToDomain(resp)
	return publications[0], nil
}
