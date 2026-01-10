package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/rafaelpaguiar/golang4devops/http-login-packaged/pkg/api"
)

func main() {

	var (
		requestURL string
		password   string
		parsedURL  *url.URL
		err        error
	)

	flag.StringVar(&requestURL, "url", "", "url to access.")
	flag.StringVar(&password, "password", "", "use a password to access our api.")

	flag.Parse()

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("Validation error: URL is not valid: %s\n", err)
		flag.Usage()
		os.Exit(1)
	}

	apiInstance := api.New(api.Options{
		Password: password,
		LoginURL: parsedURL.Scheme + "://" + parsedURL.Host + "/login",
	})

	res, err := apiInstance.DoGetRequest(parsedURL.String())
	if err != nil {
		if requestErr, ok := err.(api.RequestError); ok {
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
