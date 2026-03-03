# HTTP/2 Deep Technical Notes

## 1. Tổng quan

HTTP/2 là phiên bản nâng cấp của HTTP/1.1, tập trung vào: -
Multiplexing - Giảm latency - Header compression - Efficient binary
framing

HTTP/2 không thay đổi semantics của HTTP (GET/POST vẫn như cũ), chỉ thay
đổi cách truyền.

------------------------------------------------------------------------

## 2. Vấn đề của HTTP/1.1

### Head-of-line Blocking (Application level)

Trong HTTP/1.1: - 1 connection chỉ xử lý 1 request tại một thời điểm -
Muốn concurrent → phải mở nhiều TCP connection

Gây: - TCP handshake overhead - Slow start nhiều lần - Resource tốn kém

------------------------------------------------------------------------

## 3. HTTP/2 Kiến trúc

### Binary Framing Layer

HTTP/1.1 là text-based. HTTP/2 dùng binary frames.

------------------------------------------------------------------------

## 4. Các khái niệm chính

### Stream

Stream = 1 request/response độc lập trong 1 TCP connection.

### Frame

Frame là đơn vị nhỏ nhất truyền đi. Các loại frame: - HEADERS - DATA -
SETTINGS - WINDOW_UPDATE - PING - RST_STREAM

### Multiplexing

Nhiều stream có thể chạy song song trong cùng 1 TCP connection.

------------------------------------------------------------------------

## 5. Flow Control

HTTP/2 có flow control: - Per stream - Per connection

------------------------------------------------------------------------

## 6. Header Compression (HPACK)

Giảm lặp header bằng: - Static table - Dynamic table

------------------------------------------------------------------------

## 7. TLS và ALPN

HTTP/2 qua HTTPS dùng: - TLS handshake - ALPN negotiate protocol: h2

------------------------------------------------------------------------

## 8. TCP Head-of-Line Blocking

HTTP/2 giải quyết HOL ở HTTP layer. Nhưng TCP vẫn yêu cầu byte đúng thứ
tự. Packet mất → toàn connection bị block.

------------------------------------------------------------------------

## 9. So sánh HTTP/1.1 vs HTTP/2

  Feature               HTTP/1.1   HTTP/2
  --------------------- ---------- ---------------------
  Framing               Text       Binary
  Multiplex             No         Yes
  Header compression    No         HPACK
  Multiple TCP needed   Yes        No
  TLS required          No         Thực tế gần như Yes

------------------------------------------------------------------------

## 10. Trong Kubernetes

  Scenario               HTTP version
  ---------------------- ---------------
  Pod → Pod (http://)    HTTP/1.1
  Pod → Pod (https://)   HTTP/2
  gRPC                   HTTP/2
  Istio mesh             Thường HTTP/2

------------------------------------------------------------------------

## 11. Khi nào HTTP/1.1 nhanh hơn?

-   Low traffic
-   Single request
-   CPU yếu

------------------------------------------------------------------------

## 12. Khi nào nên dùng HTTP/2?

-   Microservice concurrent cao
-   gRPC
-   Service mesh
-   High latency network

------------------------------------------------------------------------

## 13. HTTP/3 khác gì?

HTTP/3 dùng QUIC (UDP): - Không bị TCP HOL - Stream độc lập - Faster
handshake

------------------------------------------------------------------------

# Summary

HTTP/2 cải thiện: - Concurrency - Latency under load - Header efficiency

Nhưng không loại bỏ TCP limitation.
