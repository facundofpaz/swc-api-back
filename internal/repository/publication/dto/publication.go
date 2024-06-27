package dto

import "swc-api-back.com/internal/domain/model/publication"

type PublicationDB struct {
	ID uint64 `db:"id"`
}

func ToDomain(publicationsDB []PublicationDB) []publication.Publication {
	l := len(publicationsDB)
	resp := make([]publication.Publication, l)
	for i, p := range publicationsDB {
		resp[i].ID = publication.ID(p.ID)
	}
	return resp
}
