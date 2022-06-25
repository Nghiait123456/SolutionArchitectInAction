[======================================================================================================================================================================================================] 10s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec    106468.20   11095.04  132758.99
Latency        3.78ms     2.92ms   169.73ms
Latency Distribution
50%     3.31ms
75%     4.52ms
90%     5.80ms
95%     6.64ms
99%     8.47ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 1057615, 5xx - 0
others - 0
Throughput:   187.00MB/s
root@ip-172-31-21-150:/home/ubuntu# docker run -ti --rm --ulimit nofile=65535:65535 --network=host alpine/bombardier --http1 -c 800   -d 10s -t 1s  -l http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/
Bombarding http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/ for 10s using 800 connection(s)
[======================================================================================================================================================================================================] 10s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec    100596.06   15657.08  127293.13
Latency        7.98ms     4.09ms   218.26ms
Latency Distribution
50%     7.35ms
75%    10.91ms
90%    14.73ms
95%    17.47ms
99%    23.66ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 1001551, 5xx - 0
others - 0
Throughput:   177.07MB/s
root@ip-172-31-21-150:/home/ubuntu# docker run -ti --rm --ulimit nofile=65535:65535 --network=host alpine/bombardier --http1 -c 2000   -d 10s -t 1s  -l http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/
Bombarding http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/ for 10s using 2000 connection(s)
[======================================================================================================================================================================================================] 10s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec     93861.72   18869.76  132518.11
Latency       21.49ms    12.83ms   278.14ms
Latency Distribution
50%    21.14ms
75%    29.41ms
90%    47.16ms
95%    55.04ms
99%    80.80ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 930568, 5xx - 0
others - 0
Throughput:   164.27MB/s
root@ip-172-31-21-150:/home/ubuntu# docker run -ti --rm --ulimit nofile=65535:65535 --network=host alpine/bombardier --http1 -c 400   -d 10s -t 1s  -l http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/
Bombarding http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/ for 10s using 400 connection(s)
[======================================================================================================================================================================================================] 10s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec    105983.90   11521.03  129358.82
Latency        3.79ms     3.05ms   210.28ms
Latency Distribution
50%     3.30ms
75%     4.51ms
90%     5.81ms
95%     6.66ms
99%     8.54ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 1053295, 5xx - 0
others - 0
Throughput:   186.27MB/s
[====================================================================================================================================================================================================] 1m40s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec    107517.70    8125.00  132859.36
Latency        3.72ms   843.24us   215.95ms
Latency Distribution
50%     3.30ms
75%     4.50ms
90%     5.75ms
95%     6.57ms
99%     8.32ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 10743531, 5xx - 0
others - 0
Throughput:   190.05MB/s
[====================================================================================================================================================================================================] 3m20s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec    107919.49    7820.72  131830.94
Latency        3.71ms     0.88ms   586.09ms
Latency Distribution
50%     3.31ms
75%     4.46ms
90%     5.75ms
95%     6.56ms
99%     8.32ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 21572595, 5xx - 0
others - 0
Throughput:   190.81MB/s
root@ip-172-31-21-150:/home/ubuntu# docker run -ti --rm --ulimit nofile=65535:65535 --network=host alpine/bombardier --http1 -c 400   -d 10s -t 1s  -l http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/
Bombarding http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/ for 10s using 400 connection(s)
[======================================================================================================================================================================================================] 10s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec    105983.90   11521.03  129358.82
Latency        3.79ms     3.05ms   210.28ms
Latency Distribution
50%     3.30ms
75%     4.51ms
90%     5.81ms
95%     6.66ms
99%     8.54ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 1053295, 5xx - 0
others - 0
Throughput:   186.27MB/s
root@ip-172-31-21-150:/home/ubuntu# docker run -ti --rm --ulimit nofile=65535:65535 --network=host alpine/bombardier --http1 -c 400   -d 100s -t 1s  -l http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/
Bombarding http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/ for 1m40s using 400 connection(s)
[====================================================================================================================================================================================================] 1m40s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec    107517.70    8125.00  132859.36
Latency        3.72ms   843.24us   215.95ms
Latency Distribution
50%     3.30ms
75%     4.50ms
90%     5.75ms
95%     6.57ms
99%     8.32ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 10743531, 5xx - 0
others - 0
Throughput:   190.05MB/s
root@ip-172-31-21-150:/home/ubuntu# docker run -ti --rm --ulimit nofile=65535:65535 --network=host alpine/bombardier --http1 -c 400   -d 200s -t 1s  -l http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/
Bombarding http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/ for 3m20s using 400 connection(s)
[====================================================================================================================================================================================================] 3m20s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec    107919.49    7820.72  131830.94
Latency        3.71ms     0.88ms   586.09ms
Latency Distribution
50%     3.31ms
75%     4.46ms
90%     5.75ms
95%     6.56ms
99%     8.32ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 21572595, 5xx - 0
others - 0
Throughput:   190.81MB/s
root@ip-172-31-21-150:/home/ubuntu# docker run -ti --rm --ulimit nofile=65535:65535 --network=host alpine/bombardier --http1 -c 400   -d 600s -t 1s  -l http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/
Bombarding http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/ for 10m0s using 400 connection(s)
[====================================================================================================================================================================================================] 10m0s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec    108503.09    7588.83  135235.25
Latency        3.68ms   482.06us      1.00s
Latency Distribution
50%     3.27ms
75%     4.42ms
90%     5.71ms
95%     6.52ms
99%     8.30ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 65080275, 5xx - 0
others - 1
Errors:
Get http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/: net/http: request canceled (Client.Timeout exceeded while awaiting headers) - 1
Throughput:   191.88MB/s
