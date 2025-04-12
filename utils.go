package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

const (
	newsUrl     = "https://api.polygon.io/v2/reference/news"
	maxBodySize = 1 << 20
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

func structToMap[T any](v T) map[string]string {
	param := make(map[string]string)

	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	if t.Kind() != reflect.Struct {
		return param
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := val.Field(i)
		fieldKind := field.Type.Kind()

		if !fieldValue.CanInterface() {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		tagName := strings.Split(jsonTag, ",")[0]
		if tagName == "-" {
			continue
		}

		if fieldKind == reflect.Pointer {
			if !fieldValue.IsNil() {

				param[tagName] = fmt.Sprintf("%v", fieldValue.Elem().Interface())
			}
		} else if fieldKind == reflect.Struct {
			// Skip nested structs
			continue
		} else {
			// For all other types, just convert to string
			param[tagName] = fmt.Sprintf("%v", fieldValue.Interface())
		}
	}

	return param
}

func formatUrl(baseUrl string, paramMap map[string]string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}
	values := url.Values{}

	for i, j := range paramMap {
		values.Add(i, fmt.Sprintf("%v", j))

	}
	values.Add("apiKey", fmt.Sprintf("%v", os.Getenv("API_KEY")))

	return baseUrl + "?" + values.Encode()

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

func Encode[T any](w http.ResponseWriter, r *http.Request, v T) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func Decode[T any](w http.ResponseWriter, r *http.Request, v *T) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil

}
