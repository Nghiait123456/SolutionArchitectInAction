/**
Base code from link: https://github.com/gaplo917/load-balancer-benchmark
I have many problem when i custom form this base code

1) Load test local to local:
   1) context: same to author, you create  16 nodejs instance for backend load test, you create other tool L7LB : apache-event, apache-prework, apache-worker, haproxy, nginx, traefik in same docker-compose. All instance run in same docker-network.
   2) follow request: tool loadtest ==> L7LB  => nodejs.
   3) this is perfect context, every thing in local network, every is very fast and tiny. We will return {"200:OK"},  you need go to README_AUTHOR.md and run line 53 to 71.
   4) result: in context local ==> local or local => endpoint google, we get result same 1_LoadBalancing/L7Loadbalancer/load-balancer-benchmark/result easy with cpu i5 8 core 8GB ram
            : in context product, this is point we care important. 
 
*/