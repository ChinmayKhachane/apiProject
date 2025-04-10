package main

import "fmt"

func printStockPrices(stockPrices []stockPrice) {
	for _, stockP := range stockPrices {
		fmt.Printf("Ticker: %v, Date: %v, Volume: %v, Open: %v, High: %v, Low: %v, Transactions: %v\n",
			stockP.Ticker, stockP.Date, stockP.Volume, stockP.Open, stockP.High, stockP.Low, stockP.Transactions)
	}
}

func printCompany(comp company) {
	fmt.Printf("Ticker: %s, CIK: %d, Title: %s", comp.Ticker, comp.cik, comp.Title)
}

func printFactData(facts []condensedFacts) {
	for _, fact := range facts {
		fmt.Printf("Tag: %s, Date: %s, Value: %v, Qtrs: %d \n", fact.tag, fact.date, fact.value, fact.qtrs)
	}

}
