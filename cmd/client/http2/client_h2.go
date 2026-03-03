package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

const (
	url         = "https://localhost:8443/slow"
	totalReq    = 100
	concurrency = 5
)

func main() {
	client := newH2Client()

	var wg sync.WaitGroup

	for i := 0; i < totalReq; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			req, _ := http.NewRequest("GET", url, nil)

			trace := &httptrace.ClientTrace{
				GetConn: func(hostPort string) {
					fmt.Println("→ GetConn:", hostPort)
				},
				GotConn: func(info httptrace.GotConnInfo) {
					fmt.Println("✓ GotConn reused:", info.Reused)
				},
				ConnectStart: func(network, addr string) {
					fmt.Println("Dial start:", addr)
				},
				ConnectDone: func(network, addr string, err error) {
					fmt.Println("Dial done:", addr)
				},
			}

			req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

			start := time.Now()
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("error:", err)
				return
			}

			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()

			fmt.Println("Request", i, "done in", time.Since(start))

		}(i)
	}

	wg.Wait()
}

func newH2Client() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			fmt.Println("🔥 NEW TCP DIAL:", addr)
			dialer := &net.Dialer{}
			return dialer.DialContext(ctx, network, addr)
		},
	}

	return &http.Client{
		Transport: tr,
	}
}
