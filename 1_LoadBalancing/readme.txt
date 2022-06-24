/**
  1) What is Load Balancer?
    It's a system navigating and split load. Ex: You have a billion rqs, you want 50% request to North American's region, 30% request to Asian, 20 % request to Europe. This is job of load balance.

  2) Why you need Load Balancer?
    One server don't have scale for everything. When request up, you need more server. At time, you need one system navigating and split load to this server. You need load balancer.

  3) How many type Load Balancer?
    For Hardware, have 2 type: LoadBalance software and LoadBalance Hardware.
    For layout internet, have many type:
        1) L7 LoadBalance
        2) L4 LoadBalance
        3) L3 or L3 and L4 loadBalance
        4) Dynamic routing loadBalance + smart traffic

    Maybe have many type LB, but in this doc, we dissect for 4 type before. Type 1, 2, 3, 4 is order by ascending form small project to big project. Stemming from a problem: there is no server that can expand forever for needs, LB still run on server, LB must also increase gradually according to types 1, 2, 3, 4. We will explain it in special case.
*/