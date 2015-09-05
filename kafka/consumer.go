package kafka

import (
	"flag"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

const (
	DefaultKafkaTopics   = "product"
	DefaultConsumerGroup = "consumer.go"
)

var (
	consumerGroup  = flag.String("group", DefaultConsumerGroup, "The name of the consumer group, used for coordination and load balancing")
	kafkaTopicsCSV = flag.String("topics", DefaultKafkaTopics, "The comma-separated list of topics to consume")
	zookeeper      = flag.String("zookeeper", "", "A comma-separated Zookeeper connection string (e.g. `zookeeper1.local:2181,zookeeper2.local:2181,zookeeper3.local:2181`)")

	zookeeperNodes []string
)

func init() {
	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
}

func Consumer() {
    log.Println("create consumer")
    flag.Parse()
    
    config := sarama.NewConfig()
    
    consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
    if err != nil {
        panic(err)
    }
    
    //eventCount := 0
	//offsets := make(map[string]map[int32]int64)
    
    topics, err := consumer.Topics()
    if err != nil {
        panic(err)
    }
    log.Println(topics[0])
    
    partitions, err := consumer.Partitions(topics[0])
    if err != nil {
        panic(err)
    }
    log.Println(partitions[0])
    
    pc, err := consumer.ConsumePartition(topics[0], partitions[0], 0)
    if err != nil {
        panic(err)
    }
    
    test0_msg := <-pc.Messages()
    log.Println(string(test0_msg.Value))
}


// Reference :
// [1].  https://github.com/wvanbergen/kafka/blob/master/examples/consumergroup/main.go
// [2]. https://github.com/Shopify/sarama/blob/master/consumer_test.go
