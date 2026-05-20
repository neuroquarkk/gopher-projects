package models

type ApiResponse struct {
	Ip      string  `json:"query"`
	Country string  `json:"country"`
	Region  string  `json:"regionName"`
	City    string  `json:"city"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Isp     string  `json:"isp"`
}
