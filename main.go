package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/quote", getQuote)

	log.Fatal(http.ListenAndServe(":3333", router))
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	quotes := collect()
	json.NewEncoder(w).Encode(&quotes)
}

func collect() []*Quote {
	c := colly.NewCollector()

	quote := make([]*Quote, 0)

	c.OnHTML("#ticker-data", func(e *colly.HTMLElement) {

		e.ForEach("div.item", func(i int, tr *colly.HTMLElement) {
			description := tr.ChildText("strong.name")
			price := tr.ChildText("span.number")
			variation := strings.TrimSuffix(tr.ChildText("span[class*='variation']"), "%")
			quote = append(quote, &Quote{description, toF(price), toF(variation)})

		})

	})

	c.Visit("http://www.valor.com.br/valor-data/")

	return quote
}

func toF(s string) float64 {
	s = strings.Replace(s, ",", ".", 1)
	f, _ := strconv.ParseFloat(s, 64)

	return f
}

// Quote structure
type Quote struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Variation   float64 `json:"variation"`
}
