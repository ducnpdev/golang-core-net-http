package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	url         = "http://localhost:8080/health"
	totalReq    = 10000
	concurrency = 200
)

func main() {
	fmt.Println("===== DEFAULT CLIENT =====")
	runTest(defaultClient(), true)

	fmt.Println("\n===== TUNED CLIENT =====")
	runTest(tunedClient(), true)

	fmt.Println("\n===== DEFAULT CLIENT (NO READ BODY) =====")
	runTest(defaultClient(), false)
}

func defaultClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

func tunedClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:        500,
		MaxIdleConnsPerHost: 200,
		MaxConnsPerHost:     0,
		IdleConnTimeout:     30 * time.Second,
	}

	return &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
}

func runTest(client *http.Client, readBody bool) {
	start := time.Now()

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	for i := 0; i < totalReq; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			resp, err := client.Get(url)
			if err != nil {
				fmt.Println("error:", err)
				return
			}

			if readBody {
				io.Copy(io.Discard, resp.Body)
			}

			resp.Body.Close()
		}()
	}

	wg.Wait()

	elapsed := time.Since(start)
	rps := float64(totalReq) / elapsed.Seconds()

	fmt.Printf("Total time: %v\n", elapsed)
	fmt.Printf("RPS: %.2f\n", rps)
}
