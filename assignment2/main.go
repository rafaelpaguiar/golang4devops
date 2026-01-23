package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/time/rate"
)

func main() {
	finished := make(chan struct{})
	limiter := rate.NewLimiter(5, 5)

	for i := 0; i < 10; i++ {
		go func(id int) {
			limiter.Wait(context.Background())

			resp, err := http.Get("http://localhost:8080/ratelimit")
			if err != nil {
				fmt.Printf("worker %d error: %v\n", id, err)
				finished <- struct{}{}
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("worker %d error: %v\n", id, err)
				finished <- struct{}{}
				return
			}

			fmt.Printf("worker %d: %s\n", id, body)
			finished <- struct{}{}
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-finished
	}
}
