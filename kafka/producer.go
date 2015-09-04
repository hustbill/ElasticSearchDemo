package kafka

import (
	"fmt"
	"github.com/Shopify/sarama" 
   
    "encoding/json"
    "flag"
    "../models"
)

var groupName = "trash"
var topicName = "event"
var partition int32 = 0
var client *sarama.Client
var err error
var messages = flag.Int("messages", 10, "Number of messages to send")


func Producer(product models.Product) {
	fmt.Println("creating producer")
     flag.Parse()
    
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
    
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

    // Send String
    /*
    for i := 0; i < 10; i++ {
            msg := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder(fmt.Sprintf("A%d", i))}
        producer.SendMessage(msg)
            fmt.Println("Producer send message to Kafka server")
    }*/
    
    // Send Json 
    for i := 1; i <= *messages; i++ {
     
       b, err := json.Marshal(product)
  
       msg := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder(b)}
       producer.SendMessage(msg)
       fmt.Println("Producer send message to Kafka server")
       
       if err != nil {
         panic(err)
       }
   }
   
    
}



// Reference: 
// [1].  https://gist.github.com/rayrod2030/8387924
// [2]. https://github.com/Shopify/sarama/blob/b86f86267368b80ae9aa3ae54306422c029e407d/functional_producer_test.go
// [3]. https://gist.github.com/JnBrymn/6fc38872b4d312886908