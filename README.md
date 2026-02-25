# golang-core-net-http

STEP 1 – Usage Mastery (2 tuần)
🎯 Mục tiêu:

Dùng net/http ở mức production engineer.

Task 1 – Build Production HTTP Server

Phải có:

ReadTimeout

WriteTimeout

IdleTimeout

MaxHeaderBytes

Graceful shutdown (SIGTERM)

Context cancellation

Task 2 – Viết Custom RoundTripper

Intercept:

latency

headers

body size

retry logic

Hiểu rõ:

Client → Transport → roundTrip → getConn → dial
Task 3 – Stress Test

Dùng:

wrk

hey

Benchmark:

10k concurrent

60s

Ghi lại:

RPS

p99 latency

memory

goroutine count

🟡 STEP 2 – Đọc net/http Source (2–3 tuần)

Đọc theo thứ tự này (đừng random):

1️⃣ server.go

Hiểu flow:

ListenAndServe
→ Serve
→ newConn
→ conn.serve()
→ readRequest()
→ handler.ServeHTTP()

🎯 Task:

Vẽ goroutine diagram

Hiểu tại sao mỗi connection có goroutine riêng

2️⃣ request.go

Hiểu:

Header parsing

Body reader

ContentLength handling

3️⃣ transport.go (Quan trọng nhất)

Hiểu:

roundTrip
getConn
idleConn
wantConn
dialConn

Phải hiểu:

Connection pooling

idle connection reuse

cancel context

race handling

🔵 STEP 3 – Deep Understanding (1–2 tuần)
Task 4 – Debug Connection Reuse

Viết demo:

100 request cùng host

Log connection pointer

Quan sát reuse

Task 5 – Intentionally Break Things

Không close Body

Set timeout sai

Force context cancel

Quan sát:

Goroutine leak

Idle conn stuck

Memory growth

Task 6 – Add Log vào Source

Fork Go.

Build local Go:

cd src
./make.bash

Add log vào:

conn.serve()

roundTrip()

Chạy lại test.

🔴 STEP 4 – Contribution Preparation (1 tuần)

Vào issue:

Filter:

label: net/http
label: NeedsFix

Chỉ chọn:

doc improvement

test improvement

small bug