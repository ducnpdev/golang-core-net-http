# Practice: Server IdleTimeout

Public documentation for using **`IdleTimeout`** on Go `net/http` servers. This practice helps control keep-alive connections, resource usage, and client behavior.

---

## What is IdleTimeout?

**`IdleTimeout`** is a field on [`http.Server`](https://pkg.go.dev/net/http#Server). It is the **maximum time the server will wait for the next request** on a keep-alive connection after the previous request has finished.

- **Zero** (default): no timeout; connections can stay open indefinitely between requests.
- **Non-zero**: after a connection has been idle for this duration (no new request), the server closes it.

From the [official docs](https://pkg.go.dev/net/http#Server):

> IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled. If zero, the value of ReadTimeout is used. If both are zero, there is no timeout.

---

## Why it matters

1. **Resource control** – Idle connections consume file descriptors and memory. A bounded idle time prevents unbounded growth.
2. **Security** – Very long-lived or unlimited idle connections can contribute to exhaustion and slowloris-style issues if not combined with other timeouts.
3. **Client behavior** – When the server closes a connection due to IdleTimeout, clients that reuse that connection will see errors such as `read: connection reset by peer` or `write: broken pipe`. Clients should handle retries and use a tuned `http.Transport` (see [Transport practice](transport.md)).

---

## Recommended configuration

Use **explicit timeouts** on `http.Server` instead of relying on defaults:

| Setting            | Typical range   | Purpose                                      |
|--------------------|-----------------|----------------------------------------------|
| `ReadTimeout`      | 5–10s           | Max time for reading the full request        |
| `WriteTimeout`     | 10–30s          | Max time for writing the response            |
| **IdleTimeout**    | **30–120s**     | Max time to wait for next request on conn    |
| `ReadHeaderTimeout`| ~2s (optional)  | Max time to read request headers             |
| `MaxHeaderBytes`   | e.g. 1MB        | Limit request header size                    |

Example:

```go
srv := &http.Server{
    Addr:              ":8080",
    Handler:           mux,
    ReadTimeout:       5 * time.Second,
    WriteTimeout:      10 * time.Second,
    IdleTimeout:       60 * time.Second,  // keep-alive window
    ReadHeaderTimeout: 2 * time.Second,
    MaxHeaderBytes:    1 << 20,
}
```

For **high-throughput, short-lived requests**, a **shorter** IdleTimeout (e.g. 1–10s) reduces how long idle connections are held; clients should then use a tuned `http.Transport` and connection pooling to avoid excessive “connection reset” errors (see [Transport](transport.md)).

---

## Where it is set in this project

Server timeouts are configured in [`internal/server/server.go`](../internal/server/server.go):

```go
srv := &http.Server{
    Addr:           addr,
    Handler:        mux,
    ReadTimeout:    5 * time.Second,
    WriteTimeout:   10 * time.Second,
    IdleTimeout:    1 * time.Second,  // short for load-test demo
    MaxHeaderBytes: 1 << 20,
}
```

The 1s IdleTimeout is for **demonstration**: it makes server-side idle closure visible under load. In production you would typically use a larger value (e.g. 60–120s) unless you have a specific reason to keep it short.

---

## References

- [net/http.Server](https://pkg.go.dev/net/http#Server) – IdleTimeout, ReadTimeout, WriteTimeout
- [Make resilient Go net/http servers using timeouts](https://ieftimov.com/posts/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/) – timeouts and resilience
- [Go HTTP Server production practices](https://go.dev/doc/) – general server tuning

---

## Experiments (benchmark results)

Below are example runs showing how **IdleTimeout** affects throughput and client-visible errors when using a load generator (e.g. `hey`) against this server.

### Baseline: IdleTimeout = 60s

```bash
hey -z 30s -c 200 http://localhost:8080/health
```

- **Requests/sec:** ~51,603  
- **Errors:** Few connection resets (server keeps connections longer).

### Experiment: IdleTimeout = 1s

Same command with server configured with `IdleTimeout: 1 * time.Second`:

```bash
hey -z 30s -c 200 http://localhost:8080/health
```

- **Requests/sec:** ~41,917  
- **Errors:** Many `read: connection reset by peer` and `write: broken pipe` — the server closes idle connections after 1s while the client still tries to reuse them.

Conclusion: **short IdleTimeout** reduces idle connection lifetime and can increase client-side errors unless the **client** uses a tuned Transport (connection pool, `IdleConnTimeout` aligned with server behavior). See [Transport practice](transport.md) for the client side.

Full raw output from these runs is kept in the repo history and can be reproduced by changing `IdleTimeout` in `internal/server/server.go` and re-running the above commands.
