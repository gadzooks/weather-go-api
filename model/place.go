package model

type Location struct {
	Name        string  `json:"name"`
	Region      string  `json:"region"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	SubRegion   string  `json:"subRegion"`
}

/*
snowqualmie_region:
  search_key: '04d37e830680c65b61df474e7e655d64'
  description: 'Snowqualmie Region'
*/
type Region struct {
	Name        string `json:"name"`
	SearchKey   string `json:"searchKey"`
	Description string `json:"description"`
}
