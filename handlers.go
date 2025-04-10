package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type AppHandler struct {
	db *sql.DB
}

func (h *AppHandler) FactDataTest(w http.ResponseWriter, r *http.Request) {

	testParam := factFilingParam{
		dateParam: QueryParamDate{
			Ticker: "AAPL",
		},
		tag:    "GrossProfit",
		yearly: true,
	}
	data, err := getFactData(testParam, h.db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
