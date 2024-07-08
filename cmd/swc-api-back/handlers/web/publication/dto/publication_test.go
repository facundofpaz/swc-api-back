package dto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"swc-api-back.com/cmd/swc-api-back/handlers/web/publication/dto"
	"swc-api-back.com/internal/domain/publication"
)

func TestToDomain(t *testing.T) {
	type scenario struct {
		name  string
		given dto.Publication
		want  publication.Publication
	}
	testCases := []scenario{
		{
			name: "happy path",
			given: dto.Publication{
				Latitude:     123.3,
				Longitude:    324.2,
				StreetName:   "st",
				StreetNumber: "342",
				Comment:      "comm",
				OfficesCount: 5,
				Name:         "Name",
				WorkingHours: map[int][]string{1: {"13:00 to 17:00"}},
				Tags:         []string{"with_cofee"},
			},
			want: publication.Publication{
				Latitude:     123.3,
				Longitude:    324.2,
				StreetName:   "st",
				StreetNumber: "342",
				Comment:      "comm",
				OfficesCount: 5,
				Name:         "Name",
				WorkingHours: map[int][]string{1: {"13:00 to 17:00"}},
				Tags:         []string{"with_cofee"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := dto.ToDomain(testCase.given)
			assert.Equal(t, testCase.want, got)
		})
	}
}
