/**
  1) In L7LB, endpoint request project to L7LB, L7 keep TCP connect and L7 call to backend app and forward response form backend app to endpoint.
                                      endpoint ---->  L7LB  ----> backend app

  2) From this define, these param importance for good L7LB:
      1) number concurrency rqs
      2) ram and cpu when up request to limit LB
      3) against memory leak, cpu leak ==> stability of LB when run long time

  3) Market survey:
    LBL7 common:  HAProxy, Nginx, Traefik, ApacheLB and L7LB of cloud(aws has application LB)
    We find many result benchmark L7LB. We summarize some conclusions. Don't have LB the best, every LB has strong point difference. But we choose LB suitable.
        1) L7LB of big cloud aws, google cloud is good. It's simple for implement if you use this cloud. It's way fast and good if you choose cloud.
        2) For number concurrency request, HAProxy and Nginx have good point.
        3) For memory, HAProxy and Nginx is not good, Apache-prework, Apache-wroker is good
        4) For memory/concurrency, it's same 3.
  In one context, you have good choose. But, you remember important parameter for compare opensource LB, it's  1,2,3 for in "these param importance for good L7LB:". Opensource up and down continuous, you don't change it, you need understand L7LB work and remember important param  for compare and good select.
  We setups and benchmark some case.

  4) Detail benchmark in many case and many context, i implements detail in 1_LoadBalancing/L7Loadbalancer/LoadBalancerBenchmark/README.md
     
*/
