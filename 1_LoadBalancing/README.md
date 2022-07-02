# Table of Contents
- [What is Load Balancer?](#WhatIsLoadBalancer)
- [Why you need Load Balancer?](#WhyYouNeedLoadbalancer)
- [How many type Load Balancer?](#HowManyTypeLoadbalancer)
- [Dissect load balancer?](#DissectLoadBalancer)
    - [L7 LoadBalancer](#L7LoadBalancer)
    

## What is Load Balancer? <a name="WhatIsLoadBalancer"></a>
  It's a system navigating and split load. Ex: You have a billion rqs, you want 50% request to North American's region, 30% request to Asian, 20 % request to Europe. This is job of load balance.

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

