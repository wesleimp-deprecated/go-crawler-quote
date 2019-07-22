package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	teams := collect()
	res, _ := json.Marshal(&teams)
	fmt.Println(string(res))
}

func collect() []*Quote {
	c := colly.NewCollector()

	var quote []*Quote = make([]*Quote, 0)

	c.OnHTML("#ticker-data", func (e *colly.HTMLElement)  {

		e.ForEach("div.item", func(i int, tr *colly.HTMLElement){
			description := tr.ChildText("strong.name")
			price := tr.ChildText("span.number")
			variation := strings.TrimSuffix(tr.ChildText("span[class*='variation']"), "%")
			quote = append(quote, &Quote{description, toF(price), toF(variation) })

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

type Quote struct {
	Description	string  `json:"description"`
	Price				float64 `json:"price"`
	Variation		float64 `json:"variation"`
}