package dto

import (
	"swc-api-back.com/internal/domain/publication"
)

type Publication struct {
	ID           publication.ID   `json:"id,omitempty"`
	Latitude     float32          `json:"latitude,omitempty"`
	Longitude    float32          `json:"longitud,omitempty"`
	StreetName   string           `json:"street_name,omitempty"`
	StreetNumber string           `json:"street_number,omitempty"`
	Comment      string           `json:"comment,omitempty"`
	OfficesCount int              `json:"offices_count,omitempty"`
	Name         string           `json:"name,omitempty"`
	WorkingHours map[int][]string `json:"working_hours,omitempty"`
	Tags         []string         `json:"tags,omitempty"`
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
