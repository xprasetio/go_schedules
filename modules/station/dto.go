package station

type Station struct { // untuk get all station dari api jakarta mrt
	Id   string `json:"nid"`
	Name string `json:"title"`
}

type StationResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	StationId          string `json:"nid"`
	StationName        string `json:"title"`
	ScheduleBundaranHI string `json:"jadwal_hi_biasa"`
	ScheduleLebakBulus string `json:"jadwal_hb_biasa"`
}

type ScheduleResponse struct {
	StationName string `json:"station"`
	Time        string `json:"time"`
}
