package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/iancoleman/strcase"
)

func main() {
	QuoteAPI()
}

func QuoteAPI() {
	url := "https://famous-quotes4.p.rapidapi.com/random?category=all&count=2"

	req, errRequestingQuote := http.NewRequest("GET", url, nil)
	if errRequestingQuote != nil {
		log.Println("Failed to make a request")
	}

	req.Header.Add("x-rapidapi-key", "8d6ceaae3fmsh77482db1c42818ep137cb8jsn6e55c2d0bf2e")
	req.Header.Add("x-rapidapi-host", "famous-quotes4.p.rapidapi.com")

	res, errReceivingQuote := http.DefaultClient.Do(req)
	if errReceivingQuote != nil {
		log.Println("Failed to receive a response")
	}

	defer res.Body.Close()

	response, errReadingQuote := io.ReadAll(res.Body)
	if errReadingQuote != nil {
		log.Println("Failed to read response")
	}

	QuoteJSON(&response)
}

type Quote struct {
	Author   string `Author:"author"`
	Category string `Category:"category"`
	Text     string `Text:"text"`
}

var quote [1]Quote

func QuoteJSON(response *[]byte) {
	errUnmarshalingQuote := json.Unmarshal(*response, &quote)
	if errUnmarshalingQuote != nil {
		log.Println("Failed to unmarshal")
	}

	for _, output := range quote {
		fmt.Printf("\nAuthor: %v\nCategory: %v\nQuote: %v\n", output.Author, strcase.ToCamel(output.Category), output.Text)
	}
}
