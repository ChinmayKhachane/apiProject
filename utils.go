package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"log"
	"net/url"
)

const (
	newsUrl = "https://api.polygon.io/v2/reference/news"
)

func printStockPrices(stockPrices []stockPrice) {
	for _, stockP := range stockPrices {
		fmt.Printf("Ticker: %v, Date: %v, Volume: %v, Open: %v, High: %v, Low: %v, Transactions: %v\n",
			stockP.Ticker, stockP.Date, stockP.Volume, stockP.Open, stockP.High, stockP.Low, stockP.Transactions)
	}
}

func printCompany(comp company) {
	fmt.Printf("Ticker: %s, CIK: %d, Title: %s", comp.Ticker, comp.Cik, comp.Title)
}

func printFactData(facts []condensedFacts) {
	for _, fact := range facts {
		fmt.Printf("Tag: %s, Date: %s, Value: %v, Qtrs: %d \n", fact.Tag, fact.Date, fact.Value, fact.Qtrs)
	}

}

func formatted_url(baseUrl string, params map[string]string) string {

	_, ok := params["ticker"]
	if !ok {
		log.Fatal(ok)
	}

	_, ok = params["limit"]
	if !ok {
		log.Fatal(ok)
	}

	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func newsRequest(url string) []articleInfo {
	client := resty.New()

	resp, err := client.R().Get(url)
	if err != nil {
		log.Fatal(err)
	}
	jsonResponse := resp.String()

	results := gjson.Get(jsonResponse, "results")

	articles := []articleInfo{}

	for _, result := range results.Array() {
		candArticle := articleInfo{
			PublisherName: result.Get("publisher.name").String(),
			Title:         result.Get("title").String(),
			Author:        result.Get("author").String(),
			ArticleURL:    result.Get("article_url").String(),
		}
		articles = append(articles, candArticle)

	}
	return articles

}
