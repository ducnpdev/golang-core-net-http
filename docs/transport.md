# Practice: Client Transport

Public documentation for using **`http.Transport`** in Go HTTP clients. Configuring Transport is essential for connection reuse, performance, and compatibility with servers that use short **IdleTimeout** (see [IdleTimeout practice](idletimeout.md)).

---

## What is Transport?

**`http.Transport`** implements [`http.RoundTripper`](https://pkg.go.dev/net/http#RoundTripper). It manages:

- **Connection pooling** – reuse of TCP connections across requests
- **Dial and TLS** – how connections are established
- **Timeouts** – dial, TLS handshake, response headers, idle connection lifetime

When you use `http.Client` without setting `Transport`, the client uses a default Transport with conservative limits (e.g. `DefaultMaxIdleConnsPerHost = 2`). For high concurrency or when talking to servers with short keep-alive (IdleTimeout), a **custom Transport** is recommended.

Request flow:

```text
Client.Do(req) → Transport.RoundTrip(req) → getConn → dial (or reuse idle conn)
```

---

## Deep dive: observing Transport (instrumented client)

The project includes an **instrumented client** that logs Transport behavior so you can see connection reuse, new dials, and when connections are closed. Code: [`cmd/client/instrument_client/main.go`](../cmd/client/instrument_client/main.go).

### How it works

1. **`httptrace.ClientTrace`** – Hooks into the Transport lifecycle for each request:
   - **GetConn(hostPort)** – Transport is about to get a connection (from pool or new dial).
   - **GotConn(info)** – A connection was obtained; `info.Reused` is `true` if it came from the idle pool, `false` if newly dialed.
   - **PutIdleConn(err)** – After the response is done, Transport puts the connection back to the idle pool; if `err != nil`, the connection was not reused (e.g. server closed it).
   - **ConnectStart** / **ConnectDone** – Actual TCP connect start and finish (only when a new connection is created).

2. **Custom `DialContext`** – Wraps the real dialer and logs **`NEW TCP DIAL`** only when Transport actually opens a new TCP connection (no idle conn available).

3. **`loggedConn`** – Wraps `net.Conn` and logs **`TCP CLOSED`** when the connection is closed (by server, by Transport after use, or when the pool discards it).

### Flow mapped to Transport

```text
Request starts
    → GetConn(hostPort)           "I need a conn for this host"
    → [Transport tries idle pool]
        → If idle conn found:
            → GotConn(Reused: true)
        → If not:
            → ConnectStart        (TCP dial starts)
            → NEW TCP DIAL        (our DialContext)
            → ConnectDone
            → GotConn(Reused: false)
    → RoundTrip runs (request/response)
    → Response body closed
    → PutIdleConn(err)            "Put conn back to pool" (err=nil) or "could not reuse" (err≠nil)
    → If conn actually closed:    TCP CLOSED (our wrapper)
```

So: **GetConn** = “need conn”; **GotConn(reused)** = “got one (from pool or new)”; **PutIdleConn** = “done with conn, back to pool or closed”; **NEW TCP DIAL** / **TCP CLOSED** = real TCP open/close.

### Transport config in the instrumented client

```go
tr := &http.Transport{
    MaxIdleConns:        100,
    MaxIdleConnsPerHost: 5,        // small so you see both reuse and new dials
    IdleConnTimeout:     30 * time.Second,
    DialContext:         func(...) (net.Conn, error) { ... },  // logs NEW TCP DIAL
}
```

With **concurrency 5** and **MaxIdleConnsPerHost: 5**, the first 5 requests cause new dials; later requests reuse those connections (GotConn reused: true, no NEW TCP DIAL). If the server has a short IdleTimeout, you may see **PutIdleConn** with `err != nil` and **TCP CLOSED** when the server closes the connection.

### How to run and interpret

1. Start the server (short IdleTimeout for demo):
   ```bash
   go run cmd/server/main.go
   ```
2. In another terminal, run the instrumented client:
   ```bash
   go run cmd/client/instrument_client/main.go
   ```

Look for:

- **GotConn reused: false** + **NEW TCP DIAL** → new connection.
- **GotConn reused: true** (no dial) → connection from pool.
- **PutIdleConn: &lt;nil&gt;** → connection returned to pool.
- **PutIdleConn: &lt;err&gt;** or **TCP CLOSED** → connection not reused (e.g. server closed it or pool full).

This gives a concrete view of **getConn → idle vs dial → use → put back** inside Transport.

---

## Why configure Transport?

1. **Connection reuse** – Default client keeps only 2 idle connections per host. Under concurrency this leads to many new connections and more “connection reset” / “broken pipe” when the server closes idle connections (e.g. short IdleTimeout).
2. **Performance** – More idle connections per host and sensible timeouts reduce dials and improve RPS when the server allows keep-alive.
3. **Alignment with server** – If the server uses a short `IdleTimeout`, the client should use a larger pool and an `IdleConnTimeout` that matches or is lower than the server’s idle window, so the client does not hold connections the server has already closed.

---

## Key Transport fields

| Field | Purpose | Typical / default |
|-------|---------|--------------------|
| `MaxIdleConns` | Max idle connections **total** (all hosts) | 100 |
| `MaxIdleConnsPerHost` | Max idle connections **per host** | 2 (default); raise for high concurrency (e.g. 50–200) |
| `MaxConnsPerHost` | Max **total** connections per host (0 = unlimited) | 0 or cap (e.g. 100) |
| `IdleConnTimeout` | How long idle connections stay in the pool | 90s default; can match or be &lt; server IdleTimeout |
| `DialContext` | How TCP (and optionally TLS) connections are created | Set for dial timeout (e.g. 30s) |
| `TLSHandshakeTimeout` | Max time for TLS handshake | 10s |
| `ResponseHeaderTimeout` | Max time to wait for response headers | 10s |
| `ExpectContinueTimeout` | Max time for 100-Continue | 1s |

Official reference: [net/http.Transport](https://pkg.go.dev/net/http#Transport).

---

## Example: tuned client

This project uses a tuned client in [`cmd/client/main.go`](../cmd/client/main.go) to cope with a server that has **short IdleTimeout** (1s) and high concurrency:

```go
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
```

- **MaxIdleConnsPerHost: 200** – allows many idle connections per host so that under 200 concurrent goroutines, connections can be reused instead of opening new ones.
- **IdleConnTimeout: 30s** – client drops idle connections after 30s; if server IdleTimeout is 1s, the server will often close first, but the larger pool still reduces the impact of those closures.
- **Client.Timeout** – overall per-request timeout (dial + request + response body).

Always **read and close** the response body (`io.Copy(io.Discard, resp.Body)` then `resp.Body.Close()`) so the connection can be returned to the pool.

---

## Best practices

1. **Reuse one Transport (and Client)** – Create a single `Transport` and reuse it; it is safe for concurrent use and holds the connection pool.
2. **Close response bodies** – Call `resp.Body.Close()` (or drain the body) so connections can be reused; otherwise you get leaks and “connection reset” when the server closes the connection.
3. **Set both Transport timeouts and Client.Timeout** – Use Transport for phase-specific limits (dial, TLS, headers) and `Client.Timeout` as an overall cap.
4. **Tune per host** – For a small set of hosts and high concurrency, increase `MaxIdleConnsPerHost` (and optionally `MaxConnsPerHost`) to match your concurrency and server IdleTimeout.

---

## Relation to server IdleTimeout

When the **server** uses a short **IdleTimeout** (e.g. 1s):

- The server closes connections that have been idle for that duration.
- A **default** client (only 2 idle conns per host) opens many new connections and often hits connections that the server has already closed → `read: connection reset by peer`, `write: broken pipe`.
- A **tuned** client with a larger `MaxIdleConnsPerHost` and appropriate `IdleConnTimeout` reuses connections more effectively and finishes more requests successfully, improving RPS and reducing errors.

See [IdleTimeout practice](idletimeout.md) for the server side and the benchmark notes there.

---

## References

- [net/http.Transport](https://pkg.go.dev/net/http#Transport)
- [net/http.Client](https://pkg.go.dev/net/http#Client)
- [net/http/httptrace](https://pkg.go.dev/net/http/httptrace) – ClientTrace for observing Transport (GetConn, GotConn, PutIdleConn, etc.)
- [DefaultMaxIdleConnsPerHost](https://pkg.go.dev/net/http#DefaultMaxIdleConnsPerHost) (value 2)
- [IdleTimeout practice](idletimeout.md) – server-side keep-alive and experiments

---

## Test / benchmark (this project)

From the project root:

1. Start the server (uses short IdleTimeout for demo):
   ```bash
   go run cmd/server/main.go
   ```
2. Run the client (default vs tuned):
   ```bash
   go run cmd/client/main.go
   ```

You should see:

- **Default client** – more errors (`connection reset by peer`, `broken pipe`) and lower RPS.
- **Tuned client** – fewer errors and higher RPS (e.g. ~27k vs ~18k in the included runs).

The client code compares default `http.Client{}` and `tunedClient()` against the same server to illustrate the impact of Transport configuration when the server has a short IdleTimeout.

**Observe Transport lifecycle (deep dive):** Run the instrumented client to see GetConn, GotConn (reused or not), PutIdleConn, and when TCP dial/close happen:

```bash
go run cmd/client/instrument_client/main.go
```

See [Deep dive: observing Transport](#deep-dive-observing-transport-instrumented-client) above for how to interpret the output.
