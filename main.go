package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/iancoleman/strcase"
)

type Quote struct {
	Author   string `Author:"author"`
	Category string `Category:"category"`
	Text     string `Text:"text"`
}

var quote [1]Quote

func main() {
	url := "https://famous-quotes4.p.rapidapi.com/random?category=all&count=2"

	req, errRequesting := http.NewRequest("GET", url, nil)
	if errRequesting != nil {
		log.Println("Failed to make a request")
	}

	req.Header.Add("x-rapidapi-key", "8d6ceaae3fmsh77482db1c42818ep137cb8jsn6e55c2d0bf2e")
	req.Header.Add("x-rapidapi-host", "famous-quotes4.p.rapidapi.com")

	res, errReceiving := http.DefaultClient.Do(req)
	if errReceiving != nil {
		log.Println("Failed to receive a response")
	}

	defer res.Body.Close()

	response, errReading := io.ReadAll(res.Body)
	if errReading != nil {
		log.Println("Failed to read response")
	}

	errUnmarshaling := json.Unmarshal(response, &quote)
	if errUnmarshaling != nil {
		log.Println("Failed to unmarshal")
	}

	for _, output := range quote {
		fmt.Printf("\nAuthor: %v\nCategory: %v\nQuote: %v\n", output.Author, strcase.ToCamel(output.Category), output.Text)
	}
}
