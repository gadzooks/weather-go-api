package model

type Location struct {
	Name        string
	Region      string
	Description string
	Latitude    float64
	Longitude   float64
	SubRegion   string
}

/*
snowqualmie_region:
  search_key: '04d37e830680c65b61df474e7e655d64'
  description: 'Snowqualmie Region'
*/
type Region struct {
	Name        string
	SearchKey   string
	Description string
}
