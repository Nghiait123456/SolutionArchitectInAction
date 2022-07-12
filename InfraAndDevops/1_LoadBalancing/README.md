# Table of Contents
- [What is Load Balancer?](#WhatIsLoadBalancer)
- [Why you need Load Balancer?](#WhyYouNeedLoadbalancer)
- [How many type Load Balancer?](#HowManyTypeLoadbalancer)
- [Dissect load balancer?](#DissectLoadBalancer)
- [L7 LoadBalancer](#L7LoadBalancer)
  - [Define](#Define)
  - [Life circle](#LifeCircle)
  - [How to good L7LB](#HowToGoodL7LB)
  - [Market Review](#MarketReview)
  - [Source code loadbalancer](#SourceCodeLoadbalancer])
  - [L7 Local to local](#L7LocalToLocal)
      - [Load test local to local](#LoadTestLocalToLocal)
        - [Context L7LB local](#ContextL7LBLocal)
        - [Process request](#ProccessRequest)
        - [Result](#ResultL7Local)
      - [Load test product](#LoadTestProduct)
        - [What is problem?](#WhatIsProblem)
        - [What is solution?](#WhatIsSolution)
        - [Result in product L7LB?](#ResultInProductL7LB)
          - [Haproxy L7LB result?](#HaproxyL7LBResult)
              - [100K rps with EC2_C5_4X_larger](#100KRpsWithEC2_C5_4X_larger)
              - [140K rps with EC2_C5_4X_larger](#100KRpsWithEC2_C5_4X_larger)
                - [Why is 140K rps ](#WhyIs140KRps)
                - [Detail parameter with 140K rps](#DetailParameterWith140KRps)
                - [Why i don't up to up rps?](#WhyIDontUpToUpRps)
              - [todo wait aws and test with larger Cpu](#100KRpsWithEC2_C5_4X_larger)
      - [2M connect TCP HAPROXY](#2M connect TCP HAPROXY)
      - [Auto config haproxy ](#autoConfigHaroxy)
      - [Point root domain to haproxy](#PointRootDomainToHaproxy)
      - [Auto handle haproxy down](#AutoHandleHaproxyDown)
      - [Auto handle and alert backend down](#AutoHandleAndAlertBackendDown)
      - [What ssl in Haproxy](#WhatSslInHaproxy)

- [L4 LoadBalancer](#L4 LoadBalancer)
  - [Define](#L4Define) 
  - [Why you need L4LB?](#WhyYouNeedL4LB)
  - [How to L4LB work?](#HowToL4LBWork?)
  - [Why L4LB faster and many times faster than L7LB?](#WhyL4LBFasterAndManyTimesFasterThanL7LB?)
  - [Implement in local?](#ImplementL4LBInLocal?)
  - [Implement in product?](#ImplementL4LBInProduct?)
  - [Result in product?](#ResultInProduct?)
    - [TodoWithManyOption?](#TodoWithManyOption?)
  - [Load balancer instance with L4Lb?](#LoadBalancerInstanceWithL4Lb?)
  - [What happen when L4LB down?](#WhatHappenWhenL4LBDown?)
    - [What solution?](#WhatSolutionL4LbDown?)
    - [Real time cluster wide tracking?](#RealTimeClusterWideTrackingL4LB?)
  - [Same solution L4LB in AWS?](#SameSolutionL4LBInAWS?)
  - [Point root domain to L4 Haproxy](#PointRootDomainToL4Haproxy)
  - [Simple load balancer regison base DNS and L4LB](#PointRootDomainToL4Haproxy)
      
                
## What is Load Balancer? <a name="WhatIsLoadBalancer"></a>
  It's a system navigating and split load. Ex: You have a billion rps, you want 50% request to North American's region, 30% request to Asian, 20 % request to Europe. This is job of load balance.

## Why you need Load Balancer?  <a name="WhyYouNeedLoadbalancer"></a>
  One server don't have scale for everything. When request up, you need more server. At time, you need one system navigating and split load to this server. You need load balancer.

## How many type Load Balancer?  <a name="HowManyTypeLoadbalancer"></a>
  For Hardware, have 2 type: LoadBalance software and LoadBalance Hardware. </br>
  For layout internet, have many type: </br>
     1) L7 LoadBalance </br>
     2) L4 LoadBalance </br>
     3) L3 or L3 and L4 loadBalance </br>
     4) Dynamic routing loadBalance + smart traffic </br>

  =================> Maybe have many type LB, but in this doc, we dissect for 4 type before. Type 1, 2, 3, 4 is order by ascending form small project to big project. Stemming from a problem: there is no server that can expand forever for needs, LB still run on server, LB must also increase gradually according to types 1, 2, 3, 4. We will explain it in specify type.

## Dissect load balancer? <a name="DissectLoadBalancer"></a>
## L7 LoadBalancer <a name="L7LoadBalancer"></a>
## Define <a name="Define"></a>

![](../img/1_LoadBalancing/L7LBDefine.png)

Simple define: L7LB is load balancer work on layer 7: Http, Ftp, Irc, Ssh, Dns,... It's directly interacting with the application backend system.

## Life circle <a name="LifeCircle"></a>
In L7LB, endpoint request project to L7LB, L7 keep TCP connect and L7 call to backend app and forward response form backend app to endpoint.  </br>
============================== endpoint ---->  L7LB  ----> backend app ===========================

## How to good L7LB <a name="HowToGoodL7LB"></a>
From this define, these param importance for good L7LB:
1) Number concurrency rps is good
2) Ram and cpu when up request to limit LB is small and has a stable operating threshold.
3) Against memory leak, cpu leak ==> stability of LB when run long time
4) Low latency

## Market Review <a name="MarketReview"></a> 
LBL7 common:  HAProxy, Nginx, Traefik, ApacheLB and L7LB of cloud(aws has application LB)
We find many result benchmark L7LB. We summarize some conclusions. Don't have LB the best, every LB has strong point difference. But we choose LB suitable. </br>
1) L7LB of big cloud aws, google cloud is good. It's simple for implement if you use this cloud. It's way fast and good if you choose cloud.
2) For number concurrency request, HAProxy and Nginx have good point.
3) For memory, HAProxy and Nginx is not good, Apache-prework, Apache-wroker is good
4) For memory/concurrency, it's same 3. </br>
In one context, you have good choose. But, you remember important parameter for compare opensource LB, it's  1,2,3,4  for in "How to good L7LB". </br>
Opensource up and down continuous, you don't change it, you need understand L7LB work and remember param for compare and good select. </br>

## Source code load balancer <a name="SourceCodeLoadbalancer"></a>
Base code from link: https://github.com/gaplo917/load-balancer-benchmark </br>
I have many problem then i needs custom for this form this base code. Link source is : https://github.com/Nghiait123456/InfraForBigProject/tree/master/1_LoadBalancing/L7Loadbalancer/LoadBalancerBenchmark </br>

## L7 local to local <a name="L7LocalToLocal"></a> 
## Load test local to local <a name="LoadTestLocalToLocal"></a> 
## Context L7LB local <a name="ContextL7LBLocal"></a>
Same to author, you create 3 nodejs instance for backend load test, you create other tool L7LB : apache-event, apache-prework, apache-worker, haproxy, nginx, traefik in same docker-compose. All instance run in same docker-network.
## Process request <a name="ProccessRequest"></a>
 ================================  tool loadtest ==> L7LB  => nodejs backend =================================

## Result <a name="ResultL7Local"></a>
1) In context local ==> local, we get result /L7Loadbalancer/LoadBalancerBenchmark/result easy with cpu i5 8 core 8GB ram
2) You are very fast have quick test in local, but in product, it's not simple. <br/>

## Load test product <a name="LoadTestProduct"></a>
## What is problem? <a name="WhatIsProblem"></a>
1) 4 tuple TCP: TCP uses 4-tuple (source IP, source port, destination IP, destination port)
   1 tcp connect define form 4 parameter source IP + source port + destination IP + destination port,combined 4 parameter is unique. It's required for defined one connect
2) ulimit in Linux: in Linux, everything is file, ulimit is parameter define maximum number of files open in one time
3) design LB : endpoint ==> LB  ==> backend, test high performance, i need many backend.

## What is solution?  <a name="WhatIsSolution"></a>

Solution and ideal from 3 theory :
1) From theory 1, i use one Ip for load test, but if it goes to limited, i will use multi ip load test
2) From theory 2, i set up ulimit to maximum in linux.
3) From theory 3, in load test, we have data not simple is {"ok"} and response from backend very fast. But in rps high, you pay a lot of money and pay many times for implements it. I have simple solution, i use big project for load test. EX: gooogle.com:80 ==> return 404, and response 60 character in 100 ms, it's perfect for load test. </br>
I call google.com (backend) from aws (L7LB), it's very good and keep status response is 404.

## Result in product L7LB  <a name="ResultInProductL7LB"></a>
## Haproxy L7LB result? <a name="HaproxyL7LBResult"></a>
## 100K rps with EC2_C5_4X_larger <a name="100KRpsWithEC2_C5_4X_larger"></a>
Context test in product <br/>
1) For LB: we build one instance LoadBalancerBenchmark/docker-compose-haproxy-local-to-gg.yml in this. (docker-compose -f docker-compose-haproxy-local-to-gg.yml up)<br/>
2) For Load Test, we build  one instance EC2_C5_4X_larger, build docker and use bombardier for load test. It's same : docker run -ti --rm --ulimit nofile=65535:65535 --network=host alpine/bombardier --http1 -c 400   -d 600s -t 1s  -l http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/ <br/>
Result:  <br/>
1) We test with rps from 1000 to 100.000 rps, time test form 10 s to 10 min,  total request from 10 000 to 65080275, cpu <= 60 %, ram is very good  <br/>
2) When we stop test, cpu is very fast to ~ 0 or 1 %, ram is very fast free => don't have leak ram, leak cpu  <br/>
3) Detail result in link /L7Loadbalancer/LoadBalancerBenchmark/result/HaproxyProduct/c5_4x_large.md  <br/>
=>>>>>>>>>>>>>>>>>>>>  Haproxy is very good for rps, with one instance EC2_C5_4X_larger, we pass and run stable 100.000 rps, we run stable in longtime, we have 100 M request continuous in 10 to 15 min and cpu and ram is good.  <br/>


## 140K rps with EC2_C5_4X_larger?  <a name="140KRpsWithEC2_C5_4X_larger"></a>
yup, 140 k rps with one instance c5_4x_large
This is big number for one instance LB, 140k rps is enough for most mid-range web sites, about 20.30 M users online. Of course this depends on the application, here I consider social networks. The average number of users interacting in 1s is 8 to 10% of online users. This number is also quite similar to facebook in terms of ratio. In 2010, facebook had 700 M user and rps is 13 M.

## Why is 140K rps <a name="WhyIs140KRps"></a>
Review case study 100 K rps, i find some problem: </br>
1) cpu in range 50 to 60 % ==> can completely push to stable threshold at 65 to 75 % </br>

2) There is a lot of RAM left over, only using about 15 % ==> Obviously with the problem of reponse api json, LB Haproxy does not use much ram, it uses CPU ==> choose devices with high cpu/ram will get better performance with context fast response json api. </br>

3) Bandwith, total connect Tcp is small but loadtest don't increment rps ==> i use more LB for loadtest to increase number of requests. </br>


## Detail parameter with 140K rps?  <a name="DetailParameterWith140KRps"></a>
I use 2 instances of C5_4x_larger for loadtest and 1 instance of C5_4x_larger for LB. </br>
With test in 60s, every Loadtest return71112.85 rps and 4232563 total request. ==> total, i have 143 k rps and ~ 8.5 m request total. The cpu is stable at 70% and released to ~0 % as soon as the loadtest is stopped. </br>

With test in 600s, every Loadtest return 69817.63 rps and 41625417 total rquest ==> total, i have ~ 140 K rps and 83 M request total. This is a huge number for mid-range products. The cpu is stable at 70% and released to ~0 % as soon as the loadtest is stopped. </br>

detail result in link: 1_LoadBalancing/L7Loadbalancer/LoadBalancerBenchmark/result/HaproxyProduct/140K_rps_c5_4x_large.md </br>



## Why i don't up to up rps?  <a name="WhyIDontUpToUpRps"></a>
I was configed process number is 10. All cpu run is ~ 70% cpu. With me, and with many other documents and experiences, this is the safe max when running the product. Instead of pushing the CPU to 70 to 85% to increase the load, increase the LB server configuration or scale the LB horizontally.


