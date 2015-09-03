package kafka

import (
	"fmt"
	"github.com/Shopify/sarama" 
	"time"
    "os/exec"
    "encoding/json"
    "flag"
)

var groupName = "trash"
var topicName = "event"
var partition int32 = 0
var client *sarama.Client
var err error
var messages = flag.Int("messages", 10000, "Number of messages to send")

type Audit struct {
  AuditUUID string
  WhenAudited time.Time
  WhatURI string
  WhoURI string
  WhereURI string
  WhichChanged string
}


func producer() {
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
       audit_uuid, err := exec.Command("uuidgen").Output()
       if err != nil {
         panic(err)
       }

       what_uri, err := exec.Command("uuidgen").Output()
       if err != nil {
         panic(err)
       }

       who_uri, err := exec.Command("uuidgen").Output()
       if err != nil {
         panic(err)
       }

       where_uri, err := exec.Command("uuidgen").Output()
       if err != nil {
         panic(err)
       }

       which_uri, err := exec.Command("uuidgen").Output()
       if err != nil {
         panic(err)
       }

       when_audited := time.Now()

       m := Audit{AuditUUID: string(audit_uuid), WhenAudited: when_audited, WhatURI: string(what_uri), WhoURI: string(who_uri), WhereURI: string(where_uri), WhichChanged: string(which_uri)}

       b, err := json.Marshal(m)
  
       msg := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder(b)}
       producer.SendMessage(msg)
       fmt.Println("Producer send message to Kafka server")
       
       if err != nil {
         panic(err)
       }
   }
   
    
}

func main() {

	go producer()

	<-make(chan int)
}

// Reference: 
// [1].  https://gist.github.com/rayrod2030/8387924
// [2]. https://github.com/Shopify/sarama/blob/b86f86267368b80ae9aa3ae54306422c029e407d/functional_producer_test.go
// [3]. https://gist.github.com/JnBrymn/6fc38872b4d312886908