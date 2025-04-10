package main

import (
	"log"
)

func main() {

	db, err := CreateConn()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queryparam := factFilingParam{
		dateParam: QueryParamDate{},
		tag:       "GrossProfit",
		yearly:    true,
	}
	condFacts, err := getFactData(queryparam, db)
	if err != nil {
		log.Fatal(err)
	}
	printFactData(condFacts)

}
