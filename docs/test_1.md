# run:
1. run cli
- IdleTimeout: 60s
```
hey -z 30s -c 200 http://localhost:8080/health
```
- output
```
Summary:
  Total:        30.0033 secs
  Slowest:      0.1493 secs
  Fastest:      0.0000 secs
  Average:      0.0060 secs
  Requests/sec: 51603.8488
  
  Total data:   157746028 bytes
  Size/request: 157 bytes

Response time histogram:
  0.000 [1]     |
  0.015 [984342]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.030 [11455] |
  0.045 [2316]  |
  0.060 [1005]  |
  0.075 [235]   |
  0.090 [272]   |
  0.105 [70]    |
  0.119 [99]    |
  0.134 [21]    |
  0.149 [184]   |


Latency distribution:
  10%% in 0.0010 secs
  25%% in 0.0022 secs
  50%% in 0.0032 secs
  75%% in 0.0044 secs
  90%% in 0.0065 secs
  95%% in 0.0090 secs
  99%% in 0.0182 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0163 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0065 secs
  req write:    0.0000 secs, 0.0000 secs, 0.1041 secs
  resp wait:    0.0057 secs, 0.0000 secs, 0.1491 secs
  resp read:    0.0002 secs, 0.0000 secs, 0.0938 secs

Status code distribution:
  [200] 1000000 responses

Error distribution:
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60098->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60099->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60100->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60101->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60102->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60103->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60104->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60105->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60106->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60107->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60108->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60109->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60110->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60111->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60112->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60113->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60114->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60115->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60116->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60117->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60118->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60119->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60120->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60121->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60122->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60123->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60124->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60125->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60126->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60127->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60128->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60129->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60130->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60131->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60132->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60133->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60134->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60135->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60136->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60137->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60138->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60139->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60140->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60141->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60142->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60143->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60144->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60145->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60146->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60147->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60148->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60149->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60150->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60151->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60152->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60153->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60154->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60155->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60156->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60157->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60158->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60159->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60160->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60161->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60162->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60163->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60164->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60165->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60166->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60167->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60168->[::1]:8080: read: connection reset by peer
```

1. run cli. Experiment 1 – Giảm IdleTimeout
- IdleTimeout: 1s
```
hey -z 30s -c 200 http://localhost:8080/health
```

- output
```
Summary:
  Total:        30.0079 secs
  Slowest:      0.3660 secs
  Fastest:      0.0000 secs
  Average:      0.0060 secs
  Requests/sec: 41917.2899
  
  Total data:   126784664 bytes
  Size/request: 126 bytes

Response time histogram:
  0.000 [1]     |
  0.037 [992517]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.073 [3809]  |
  0.110 [1641]  |
  0.146 [520]   |
  0.183 [367]   |
  0.220 [751]   |
  0.256 [301]   |
  0.293 [56]    |
  0.329 [6]     |
  0.366 [31]    |


Latency distribution:
  10%% in 0.0010 secs
  25%% in 0.0024 secs
  50%% in 0.0034 secs
  75%% in 0.0047 secs
  90%% in 0.0073 secs
  95%% in 0.0108 secs
  99%% in 0.0291 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0195 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0087 secs
  req write:    0.0000 secs, 0.0000 secs, 0.1321 secs
  resp wait:    0.0057 secs, 0.0000 secs, 0.3660 secs
  resp read:    0.0002 secs, 0.0000 secs, 0.1416 secs

Status code distribution:
  [200] 1000000 responses

Error distribution:
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60874->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60875->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60876->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60877->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60878->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60879->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60880->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60881->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60882->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60883->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60884->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60885->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60886->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60887->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60888->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60889->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60890->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60891->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60892->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60893->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60894->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60895->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60896->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60897->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60898->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60899->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60900->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60901->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60902->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60903->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60904->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60905->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60906->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60907->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60908->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60909->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60910->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60911->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60912->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60913->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60914->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60916->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60917->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60918->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60920->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60922->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60923->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60924->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60925->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60928->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60929->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60930->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60931->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60932->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60933->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60934->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60935->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60936->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60937->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60938->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60940->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60941->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60942->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": read tcp [::1]:60943->[::1]:8080: read: connection reset by peer
  [1]   Get "http://localhost:8080/health": write tcp [::1]:60919->[::1]:8080: write: broken pipe
  [1]   Get "http://localhost:8080/health": write tcp [::1]:60921->[::1]:8080: write: broken pipe
  [1]   Get "http://localhost:8080/health": write tcp [::1]:60926->[::1]:8080: write: broken pipe
  [1]   Get "http://localhost:8080/health": write tcp [::1]:60927->[::1]:8080: write: broken pipe
  [1]   Get "http://localhost:8080/health": write tcp [::1]:60939->[::1]:8080: write: broken pipe
```