package dto

import (
	"swc-api-back.com/internal/domain/model/publication"
)

type Publication struct {
	ID           publication.ID   `json:"id"`
	Latitude     float32          `json:"latitude"`
	Longitude    float32          `json:"longitud"`
	StreetName   string           `json:"street_name"`
	StreetNumber string           `json:"street_number"`
	Comment      string           `json:"comment"`
	OfficesCount int              `json:"offices_count"`
	Name         string           `json:"name"`
	WorkingHours map[int][]string `json:"working_hours"`
	Tags         []string         `json:"tags"`
}

func ToDomain(p Publication) publication.Publication {
	return publication.Publication{
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		StreetName:   p.StreetName,
		StreetNumber: p.StreetNumber,
		Comment:      p.Comment,
		OfficesCount: p.OfficesCount,
		Name:         p.Name,
		WorkingHours: p.WorkingHours,
		Tags:         p.Tags,
	}
}

func ToDto(p publication.Publication) Publication {
	return Publication{
		ID:           p.ID,
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		StreetName:   p.StreetName,
		StreetNumber: p.StreetNumber,
		Comment:      p.Comment,
		OfficesCount: p.OfficesCount,
		Name:         p.Name,
		WorkingHours: p.WorkingHours,
		Tags:         p.Tags,
	}
}
