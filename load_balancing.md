# Без балансировки

## 1
```shell
ab -n 10 -c 10 \  
  -H 'accept: application/json' \      
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc1ODMzNjgsImlkIjoxLCJvcmlnX2lhdCI6MTY5NzU3OTc2OCwicm9sZSI6ImFkbWluIn0.VH1XjZilZaJr6hlEsdRkbexKS5XVp466ADF0SO3LZm0' \
  http://localhost:4080/api/v1/stat
```
```
Server Software:        termy
Server Hostname:        localhost
Server Port:            4080

Document Path:          /api/v1/stat
Document Length:        13639 bytes

Concurrency Level:      10
Time taken for tests:   19.264 seconds
Complete requests:      10
Failed requests:        0
Total transferred:      137760 bytes
HTML transferred:       136390 bytes
Requests per second:    0.52 [#/sec] (mean)
Time per request:       19263.828 [ms] (mean)
Time per request:       1926.383 [ms] (mean, across all concurrent requests)
Transfer rate:          6.98 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       0
Processing:  1282 16207 5245.0  17898   17982
Waiting:     1281 16207 5245.2  17898   17982
Total:       1282 16207 5245.1  17899   17982

Percentage of the requests served within a certain time (ms)
  50%  17899
  66%  17921
  75%  17943
  80%  17975
  90%  17982
  95%  17982
  98%  17982
  99%  17982
 100%  17982 (longest request)
```


## 2
```shell
ab -n 50 -c 10 \  
  -H 'accept: application/json' \      
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc1ODMzNjgsImlkIjoxLCJvcmlnX2lhdCI6MTY5NzU3OTc2OCwicm9sZSI6ImFkbWluIn0.VH1XjZilZaJr6hlEsdRkbexKS5XVp466ADF0SO3LZm0' \
  http://localhost:4080/api/v1/stat
```
```
Server Software:        termy
Server Hostname:        localhost
Server Port:            4080

Document Path:          /api/v1/stat
Document Length:        13639 bytes

Concurrency Level:      10
Time taken for tests:   108.751 seconds
Complete requests:      50
Failed requests:        38
   (Connect: 0, Receive: 0, Length: 38, Exceptions: 0)
Total transferred:      688847 bytes
HTML transferred:       681997 bytes
Requests per second:    0.46 [#/sec] (mean)
Time per request:       21750.194 [ms] (mean)
Time per request:       2175.019 [ms] (mean, across all concurrent requests)
Transfer rate:          6.19 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       0
Processing:  1356 21164 3064.8  21768   24406
Waiting:     1356 21163 3064.8  21768   24405
Total:       1356 21164 3064.8  21768   24406

Percentage of the requests served within a certain time (ms)
  50%  21768
  66%  22029
  75%  22177
  80%  22241
  90%  22476
  95%  23011
  98%  24406
  99%  24406
 100%  24406 (longest request)
```



## 3
```shell
ab -n 10000 -c 100 \  
  -H 'accept: application/json' \      
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc1ODMzNjgsImlkIjoxLCJvcmlnX2lhdCI6MTY5NzU3OTc2OCwicm9sZSI6ImFkbWluIn0.VH1XjZilZaJr6hlEsdRkbexKS5XVp466ADF0SO3LZm0' \
  http://localhost:4080/api/v1/layers
```
```
Server Software:        termy
Server Hostname:        localhost
Server Port:            4080

Document Path:          /api/v1/layers
Document Length:        2 bytes

Concurrency Level:      100
Time taken for tests:   28.422 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1580000 bytes
HTML transferred:       20000 bytes
Requests per second:    351.84 [#/sec] (mean)
Time per request:       284.221 [ms] (mean)
Time per request:       2.842 [ms] (mean, across all concurrent requests)
Transfer rate:          54.29 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       6
Processing:     2  282 320.4     40    1571
Waiting:        2  282 320.4     39    1570
Total:          2  282 320.4     40    1571

Percentage of the requests served within a certain time (ms)
  50%     40
  66%    523
  75%    602
  80%    638
  90%    735
  95%    800
  98%    890
  99%    990
 100%   1571 (longest request)
```



## 4
```shell
ab -n 50000 -c 100 \  
  -H 'accept: application/json' \      
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc1ODMzNjgsImlkIjoxLCJvcmlnX2lhdCI6MTY5NzU3OTc2OCwicm9sZSI6ImFkbWluIn0.VH1XjZilZaJr6hlEsdRkbexKS5XVp466ADF0SO3LZm0' \
  http://localhost:4080/api/v1/layers
```
```
Server Software:        termy
Server Hostname:        localhost
Server Port:            4080

Document Path:          /api/v1/layers
Document Length:        2 bytes

Concurrency Level:      100
Time taken for tests:   140.852 seconds
Complete requests:      50000
Failed requests:        0
Total transferred:      7900000 bytes
HTML transferred:       100000 bytes
Requests per second:    354.98 [#/sec] (mean)
Time per request:       281.703 [ms] (mean)
Time per request:       2.817 [ms] (mean, across all concurrent requests)
Transfer rate:          54.77 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0      11
Processing:     2  279 312.8     43    3728
Waiting:        2  279 312.7     42    3727
Total:          2  279 312.8     43    3728

Percentage of the requests served within a certain time (ms)
  50%     43
  66%    513
  75%    583
  80%    620
  90%    712
  95%    779
  98%    854
  99%    912
 100%   3728 (longest request)
```



# С балансировкой

## 1
```shell
ab -n 10 -c 10 \  
  -H 'accept: application/json' \      
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc1ODMzNjgsImlkIjoxLCJvcmlnX2lhdCI6MTY5NzU3OTc2OCwicm9sZSI6ImFkbWluIn0.VH1XjZilZaJr6hlEsdRkbexKS5XVp466ADF0SO3LZm0' \
  http://localhost:4080/api/v1/stat
```
```
Server Software:        termy
Server Hostname:        localhost
Server Port:            4080

Document Path:          /api/v1/stat
Document Length:        10657 bytes

Concurrency Level:      10
Time taken for tests:   3.790 seconds
Complete requests:      10
Failed requests:        9
   (Connect: 0, Receive: 0, Length: 9, Exceptions: 0)
Total transferred:      107931 bytes
HTML transferred:       106561 bytes
Requests per second:    2.64 [#/sec] (mean)
Time per request:       3789.717 [ms] (mean)
Time per request:       378.972 [ms] (mean, across all concurrent requests)
Transfer rate:          27.81 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       0
Processing:   593 2104 923.1   2073    3197
Waiting:      592 2103 922.9   2073    3197
Total:        593 2105 923.1   2074    3197

Percentage of the requests served within a certain time (ms)
  50%   2074
  66%   2887
  75%   2966
  80%   3195
  90%   3197
  95%   3197
  98%   3197
  99%   3197
 100%   3197 (longest request)
```


## 2
```shell
ab -n 50 -c 10 \  
  -H 'accept: application/json' \      
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc1ODMzNjgsImlkIjoxLCJvcmlnX2lhdCI6MTY5NzU3OTc2OCwicm9sZSI6ImFkbWluIn0.VH1XjZilZaJr6hlEsdRkbexKS5XVp466ADF0SO3LZm0' \
  http://localhost:4080/api/v1/stat
```
```
Server Software:        termy
Server Hostname:        localhost
Server Port:            4080

Document Path:          /api/v1/stat
Document Length:        10662 bytes

Concurrency Level:      10
Time taken for tests:   23.250 seconds
Complete requests:      50
Failed requests:        48
   (Connect: 0, Receive: 0, Length: 48, Exceptions: 0)
Total transferred:      542876 bytes
HTML transferred:       536026 bytes
Requests per second:    2.15 [#/sec] (mean)
Time per request:       4650.090 [ms] (mean)
Time per request:       465.009 [ms] (mean, across all concurrent requests)
Transfer rate:          22.80 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       0
Processing:   568 4181 3432.0   4386    9628
Waiting:      568 4181 3432.3   4386    9627
Total:        568 4181 3432.0   4386    9628

Percentage of the requests served within a certain time (ms)
  50%   4386
  66%   6626
  75%   7414
  80%   8285
  90%   8877
  95%   9076
  98%   9628
  99%   9628
 100%   9628 (longest request)
```



## 3
```shell
ab -n 10000 -c 100 \  
  -H 'accept: application/json' \      
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc1ODMzNjgsImlkIjoxLCJvcmlnX2lhdCI6MTY5NzU3OTc2OCwicm9sZSI6ImFkbWluIn0.VH1XjZilZaJr6hlEsdRkbexKS5XVp466ADF0SO3LZm0' \
  http://localhost:4080/api/v1/layers
```
```
Server Software:        termy
Server Hostname:        localhost
Server Port:            4080

Document Path:          /api/v1/layers
Document Length:        2 bytes

Concurrency Level:      100
Time taken for tests:   16.638 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1580000 bytes
HTML transferred:       20000 bytes
Requests per second:    601.04 [#/sec] (mean)
Time per request:       166.378 [ms] (mean)
Time per request:       1.664 [ms] (mean, across all concurrent requests)
Transfer rate:          92.74 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       5
Processing:     1  163 283.3     12    1729
Waiting:        1  163 283.2     12    1729
Total:          2  163 283.3     12    1729

Percentage of the requests served within a certain time (ms)
  50%     12
  66%     29
  75%    115
  80%    413
  90%    679
  95%    793
  98%    909
  99%   1003
 100%   1729 (longest request)
```



## 4
```shell
ab -n 50000 -c 100 \  
  -H 'accept: application/json' \      
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc1ODMzNjgsImlkIjoxLCJvcmlnX2lhdCI6MTY5NzU3OTc2OCwicm9sZSI6ImFkbWluIn0.VH1XjZilZaJr6hlEsdRkbexKS5XVp466ADF0SO3LZm0' \
  http://localhost:4080/api/v1/layers
```
```
Server Software:        termy
Server Hostname:        localhost
Server Port:            4080

Document Path:          /api/v1/layers
Document Length:        2 bytes

Concurrency Level:      100
Time taken for tests:   95.968 seconds
Complete requests:      50000
Failed requests:        0
Total transferred:      7900000 bytes
HTML transferred:       100000 bytes
Requests per second:    521.01 [#/sec] (mean)
Time per request:       191.935 [ms] (mean)
Time per request:       1.919 [ms] (mean, across all concurrent requests)
Transfer rate:          80.39 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0      13
Processing:     1  166 322.0     14   15768
Waiting:        1  165 322.0     14   15768
Total:          2  166 322.0     14   15768

Percentage of the requests served within a certain time (ms)
  50%     14
  66%     41
  75%    177
  80%    353
  90%    602
  95%    787
  98%   1048
  99%   1235
 100%  15768 (longest request)
```
