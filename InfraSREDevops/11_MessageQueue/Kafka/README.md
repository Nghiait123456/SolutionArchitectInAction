- [Benchmark Kafka](#benchmark_kafka)

## Benchmark Kafka? <a name="benchmark_kafka"></a>

Cassandra is a nosql database for extremely high performance and extremely low latency, used with very large databases
such as social networks. </br>
https://www.scylladb.com/2021/08/24/apache-cassandra-4-0-vs-scylla-4-4-comparing-performance/

## Why kafka is high throw output? <a name="why_kafka_is_high_throw_output"></a>

In Kafka, producers and consumers are fully decoupled and agnostic of each other, which is a key design element to
achieve the high scalability that Kafka is known for. For example, producers never need to wait for consumers. Kafka
provides various guarantees such as the ability to process events exactly-once. </br>

From the above example, it can be seen that the scale out production and consummer can happen completely
independently. </br>

## Persistent data storage in kafka affects performance?  <a name="data_storage_in_kafka_affects_performance"></a>

Events in a topic can be read as often as needed—unlike traditional messaging systems, events are not deleted after
consumption. Instead, you define for how long Kafka should retain your events through a per-topic configuration setting,
after which old events will be discarded. Kafka's performance is effectively constant with respect to data size, so
storing data for a long time is perfectly fine. </br>

## Partition? <a name="partition"></a>

Topics are partitioned, meaning a topic is spread over a number of "buckets" located on different Kafka brokers. This
distributed placement of your data is very important for scalability because it allows client applications to both read
and write the data from/to many brokers at the same time. When a new event is published to a topic, it is actually
appended to one of the topic's partitions. Events with the same event key (e.g., a customer or vehicle ID) are written
to the same partition, and Kafka guarantees that any consumer of a given topic-partition will always read that
partition's events in exactly the same order as they were written. </br>
![partition.png](img%2Fpartition.png)  </br>

## I have an event and want all service listen event to receive it, what should I do? <a name="i_have_an_event_and_want_all_consumers_of_that_event_to_receive_it_what_should_i_sdo"></a>

Each service can have multiple cosumers and topics can have multiple partitions, I don't care about this. But in order
for all services to receive an event from the producer, no event is missed, each service needs to listen for a unique
GroupId. </br>

Link: https://stackoverflow.com/questions/35561110/can-multiple-kafka-consumers-read-same-message-from-the-partition </br>



