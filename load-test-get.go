package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	client := &http.Client{}

	url := flag.String("url", "http://", "The url to call")
	times := flag.Int("times", 100, "the number of times the url should be called")
	flag.Parse()

	fmt.Println(*url)
	fmt.Println(*times)
	var wg sync.WaitGroup
	returningBody := make(chan *http.Response)
	wg.Add(*times)

	go func() {
		wg.Wait()
		close(returningBody)
	}()

	for i := 0; i <= *times; i++ {
		go func() {
			req, err := http.NewRequest("GET", *url, nil)
			if err != nil {
				er := fmt.Errorf("an error has occured %e", err)
				fmt.Print(er)
			}
			req.Header.Add("Accept", "application/json")
			resp, err := client.Do(req)
			returningBody <- resp
			wg.Done()
		}()
	}

	count := 0
	for range returningBody {
		count = count + 1
		fmt.Printf(" %v", count)
	}
}
