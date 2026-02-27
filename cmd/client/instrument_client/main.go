package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

const (
	url         = "http://localhost:8080/health"
	totalReq    = 20
	concurrency = 5
)

func main() {
	client := newInstrumentedClient()

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	for i := 0; i < totalReq; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()

			req, _ := http.NewRequest("GET", url, nil)

			trace := &httptrace.ClientTrace{
				GetConn: func(hostPort string) {
					fmt.Println("→ GetConn:", hostPort)
				},
				GotConn: func(info httptrace.GotConnInfo) {
					fmt.Println("✓ GotConn reused:", info.Reused)
				},
				PutIdleConn: func(err error) {
					fmt.Println("↩ PutIdleConn:", err)
				},
				ConnectStart: func(network, addr string) {
					fmt.Println("Dial start:", addr)
				},
				ConnectDone: func(network, addr string, err error) {
					fmt.Println("Dial done:", addr, "err:", err)
				},
			}

			req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("request error:", err)
				return
			}

			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()

		}(i)
	}

	wg.Wait()
}

func newInstrumentedClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 5,
		IdleConnTimeout:     30 * time.Second,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			fmt.Println("🔥 NEW TCP DIAL:", addr)
			dialer := &net.Dialer{}
			conn, err := dialer.DialContext(ctx, network, addr)
			if err != nil {
				return nil, err
			}
			return &loggedConn{Conn: conn}, nil
		},
	}

	return &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
}

type loggedConn struct {
	net.Conn
}

func (c *loggedConn) Close() error {
	fmt.Println("❌ TCP CLOSED")
	return c.Conn.Close()
}
