package publication

type (
	ID          uint64
	Publication struct {
		ID           ID
		Description  string
		Latitude     float32
		Longitude    float32
		StreetName   string
		StreetNumber string
		Comment      string
		OfficesCount int
		Name         string
		WorkingHours map[int][]string
		Tags         []string
	}
)
