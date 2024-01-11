package repository

import (
	"context"

	"swc-api-back.com/internal/domain/model/publication"
)

type Publication interface {
	GetByID(context.Context, publication.ID) (publication.Publication, error)
}
