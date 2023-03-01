- [Cache Cluster is require for high load](#cache_cluster_is_require_for_high_load)
- [Benchmark Redis](#benchmark_redis)

## Cache Cluster is require for high load  <a name="cache_cluster_is_require_for_high_load"></a>

Redis cluster, memcache cluster are famous open source used by world leading products. To support high load and limit
i/o bottlenecks, cache cluster is almost a must-have solution. </br>

## Benchmark Redis <a name="benchmark_redis"></a>

Redis cluster has extremely good scalability and extremely high performance. I have a benchmark, it easily booked 2.5
million ops with 8 instances of EC2:cache.r6g.large, have 16VCPU ==> ~ 2.5 M ops/sec. Redis directly benchmark, it
reaches
200M ops with 40 nodes, has 2800 VCPU.
My benchmark
details: https://github.com/Nghiait123456/InfraSREDevopsBackendForBigProject/blob/master/InfraSREDevops/3_Cache/result_benmark_redis_cluster.txt </br>
200M ops bencmark details: https://redis.com/blog/redis-enterprise-extensions-linear-scalability-200m-ops-sec/ </br>

The redis extension also accommodates most major projects. Aws default allows you to have 100 clusters in one region,
and this number is completely scalable. </br>