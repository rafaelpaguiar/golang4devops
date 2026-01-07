package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response interface {
		GetResponse() string
}

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) GetResponse() string {
	return fmt.Sprintf("%s", strings.Join(w.Words, ", "))
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (o Occurrence) GetResponse() string{
	out := []string{}
	for word, occurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return fmt.Sprintf("%s", strings.Join(out, ", "))
}

func doRequests(resquestURL string) (Response, error) {

		if _, err := url.ParseRequestURI(resquestURL); err != nil {
				return nil, fmt.Errorf("Validation error: URL is not valid: %s\n", err)
				os.Exit(1)
		}

		response, err := http.Get(resquestURL)

		if err != nil {
				return nil, fmt.Errorf("http Get error: %s\n", err)
			}

		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, fmt.Errorf("ReadAll error: %s\n", err)
		}

		if response.StatusCode != 200 {
			return nil, fmt.Errorf("Invalid output (HTTP Status Code): %d\nBody: %s\n", response.StatusCode, body)
		}

		if !json.Valid(body) {
			return nil,  RequestError{
				HTTPCode: response.StatusCode,
				Body: string(body),
				Err: fmt.Sprintf("No valid JSON returned."),
			}
		}

		var page Page

		err = json.Unmarshal(body, &page)
		if err != nil {
			return nil,  RequestError{
				HTTPCode: response.StatusCode,
				Body: string(body),
				Err: fmt.Sprintf("Page unmarshal error: %s", err),
			}
		}

		switch(page.Name){
		case "words":
			var words Words

			err = json.Unmarshal(body, &words)

			if err != nil {
				return nil,  RequestError{
					HTTPCode: response.StatusCode,
					Body: string(body),
					Err: fmt.Sprintf("Words unmarshal error: %s", err),
				}
			}

			return words, nil

		case "occurrence":

			var occurrence Occurrence

			err = json.Unmarshal(body, &occurrence)

			if err != nil {
				return nil,  RequestError{
					HTTPCode: response.StatusCode,
					Body: string(body),
					Err: fmt.Sprintf("Occurrence unmarshal error: %s", err),
				}
			}

			return occurrence, nil
	}

	return nil, nil
}

func main() {

	args := os.Args

	if len(args) < 2 {
			fmt.Printf("Usage: ./http-get <url>\n")
			os.Exit(1)
	}

	res, err := doRequests(args[1])
	if err != nil {
		if requestErr, ok := err.(RequestError); ok {
			fmt.Printf("Error: %s (HTTPCode: %d, Body: %s)\n", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
			os.Exit(1)
		}
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	if res == nil {
		fmt.Printf("No response.\n")
		os.Exit(1)
	}

	fmt.Printf("Response: %s\n", res.GetResponse())

}
