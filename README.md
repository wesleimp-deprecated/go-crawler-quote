# Go-crawler-quote

Scraping to get a current quote from [Valor Data](http://www.valor.com.br/valor-data/)

## Usage

Clone the repo

```bash
$ git clone https://github.com/wesleimp/go-crawler-quote.git
$ cd go-crawler-quote
```

Generate BIN and execute


```bash
$ go build
$ ./go-crawler-quote
```

## Output

```json
[{
	"description": "Dólar Comercial",
	"price": 3.7377,
	"variation": -0.21
}, {
	"description": "Dólar Ptax",
	"price": 3.74,
	"variation": -0.02
}, {
	"description": "Dólar Turismo",
	"price": 3.89,
	"variation": -0.21
}, {
	"description": "Euro Comercial",
	"price": 4.1897,
	"variation": -0.3
}, {
	"description": "Euro x Dólar",
	"price": 1.1215,
	"variation": -0.06
}]
```

## Licence

[MIT](https://github.com/wesleimp/go-crawler-quote/blob/master/LICENSE)
