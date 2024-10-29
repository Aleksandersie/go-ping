package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	urlArr := []string{"https://github.com/", "https://www.yaplakal.com/", "https://www.google.ru/", "htts://ya.ru/"}

	resCh := make(chan int)
	errCh := make(chan error)

	var wg sync.WaitGroup

	for _, url := range urlArr {
		wg.Add(1)

		go func() {
			ping(url, resCh, errCh)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for range urlArr {
		select {
		case err := <-errCh:
			fmt.Println(err)
		case res := <-resCh:
			fmt.Printf("%d\n", res)
		}

	}

}

func ping(url string, resCh chan int, errCh chan error) {
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}

	resCh <- resp.StatusCode
}
