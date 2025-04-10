package main

type stockPrice struct {
	Ticker       string  `json:"ticker"`
	Date         string  `json:"date"`
	Volume       int64   `json:"volume"`
	Open         float64 `json:"open"`
	Close        float64 `json:"close"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Transactions int64   `json:"transactions"`
}

type company struct {
	cik    int64  `json:"cik"`
	Ticker string `json:"ticker"`
	Title  string `json:"title"`
}

type QueryParamDate struct {
	Ticker     string `json:"ticker"`
	End_date   string `json:"end_date"`
	Start_date string `json:"start_date"`
}

type factFilingRow struct {
	adsh     string `json:"adsh"`
	tag      string `json:"tag"`
	version  string `json:"version"`
	date     string `json:"date"`
	qtrs     int8   `json:"qtrs"`
	uom      string `json:"uom"`
	segments string `json:"segments"`
	value    int64  `json:"value"`
}

func (q QueryParamDate) getTicker() string {
	return q.Ticker
}
