package main

type Yr struct {
	Created     string   `json:"created"`
	Update      string   `json:"update"`
	RadarIsDown bool     `json:"radarIsDown"`
	Points      []Points `json:"points"`
	Status      Status   `json:"status"`
	Links       Links    `json:"_links"`
}

type Status struct {
	Code string `json:"code"`
}

type Links struct {
	Self   Href `json:"self"`
	Parent Href `json:"parent"`
}

type Href struct {
	Href string `json:"href"`
}

type Points struct {
	Precipitation Precipitation `json:"precipitation"`
	Time          string        `json:"time"`
}

type Precipitation struct {
	Intensity float64 `json:"intensity"`
	Phase     string  `json:"phase"`
}
