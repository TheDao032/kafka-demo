package main

import (
	"encoding/json"
	"log"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"math/rand"

	"github.com/Shopify/sarama"
)

var (
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:29092").Strings()
	topic      = kingpin.Flag("topic", "Topic name").Default("demo_topic").String()
	maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

type MessageStruct struct {
	Message string
}

func main() {
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = *maxRetry
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(*brokerList, config)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	for {
		randString := RandStringRunes(10)
		mess := MessageStruct{
			Message: "Something Cool " + randString,
		}
		marshalMessage, _ := json.Marshal(mess)
		msg := &sarama.ProducerMessage{
			Topic: *topic,
			Value: sarama.StringEncoder(marshalMessage),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)/message(%v)\n", *topic, partition, offset, mess.Message)

		time.Sleep(2 * time.Second)
	}
}
