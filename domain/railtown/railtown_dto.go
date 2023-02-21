package railtown

type RailTownREQ struct {
	Query string `json:"query"`
}

type RailTownRES struct {
	Id        int64               `json:"id"`
	Location  RailTownRESLocation `json:"location"`
	Current   RailTownRESCurrent  `json:"current"`
	CreatedAt string              `json:"created_at"`
}

func (rt *RailTownRES) FakeIT() {
	//
}

type RailTownRESLocation struct {
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type RailTownRESCurrent struct {
	TempC float32 `json:"temp_c"`
}
