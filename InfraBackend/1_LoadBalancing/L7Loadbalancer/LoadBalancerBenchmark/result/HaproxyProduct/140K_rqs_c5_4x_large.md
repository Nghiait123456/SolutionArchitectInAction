docker run -ti --rm --ulimit nofile=655355:655355  --network=host alpine/bombardier --http1 -c 2000 -d 60s -t 1s -l http://ec2-3-0-101-88.ap-southeast-1.compute.amazonaws.com:8080/
Bombarding http://ec2-3-0-101-88.ap-southeast-1.compute.amazonaws.com:8080/ for 1m0s using 2000 connection(s)
[=====================================================================================================================================================================================================] 1m0s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec     71112.85   12866.52  141519.52
Latency       28.35ms     5.89ms   355.49ms
Latency Distribution
50%    25.00ms
75%    29.27ms
90%    45.19ms
95%    58.82ms
99%    65.20ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 4232563, 5xx - 0
others - 0
Throughput:   124.62MB/s


root@ip-172-31-17-41:/home/ubuntu/InfraForBigProject/1_LoadBalancing/L7Loadbalancer# docker run -ti --rm --ulimit nofile=655355:655355  --network=host alpine/bombardier --http1 -c 2000 -d 600s -t 1s -l http://ec2-3-0-101-88.ap-southeast-1.compute.amazonaws.com:8080/
Bombarding http://ec2-3-0-101-88.ap-southeast-1.compute.amazonaws.com:8080/ for 10m0s using 2000 connection(s)
[====================================================================================================================================================================================================] 10m0s
Done!
Statistics        Avg      Stdev        Max
Reqs/sec     69817.63   10984.83  131801.22
Latency       28.83ms     6.58ms      1.01s
Latency Distribution
50%    25.43ms
75%    29.24ms
90%    44.18ms
95%    58.99ms
99%    63.63ms
HTTP codes:
1xx - 0, 2xx - 0, 3xx - 0, 4xx - 41625417, 5xx - 0
others - 580
Errors:
Get http://ec2-3-0-101-88.ap-southeast-1.compute.amazonaws.com:8080/: net/http: request canceled (Client.Timeout exceeded while awaiting headers) - 580
net/http: request canceled (Client.Timeout exceeded while reading body) - 214
Throughput:   122.59MB/s
root@ip-172-31-17-41:/home/ubuntu/InfraForBigProject/1_LoadBalancing/L7Loadbalancer# Connection to ec2-52-77-249-144.ap-southeast-1.compute.amazonaws.com closed by remote host.
Connection to ec2-52-77-249-144.ap-southeast-1.compute.amazonaws.com closed.
