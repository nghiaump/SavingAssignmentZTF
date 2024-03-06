package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
	"time"
)

const KafkaTopicSavingAccount = "NewSavingAccountCreated"
const KafkaAddress = "kafka:9092"

func main() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{KafkaAddress}, config) // them http:// la bug luon !!! too many colon???
	if err != nil {
		log.Fatalf("Failed to start Sarama consumer: %s", err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Failed to close Sarama consumer: %s", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(KafkaTopicSavingAccount, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start consumer for partition: %s", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	log.Println("Listening messages:")
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			{
				// log.Printf("Received message: %s\n", string(msg.Value))
				log.Println("Received new message and Unmarshalled to struct:")
				var account SavingAccount
				if err := json.Unmarshal(msg.Value, &account); err != nil {
					log.Printf("Failed to unmarshal message: %s", err)
					continue
				}
				log.Printf("Message details: %+v\n", account)
			}

		case <-signals:
			break ConsumerLoop
		}
	}

	log.Println("Shutting down")
	// Đảm bảo rằng consumer đã consume tất cả các tin nhắn còn lại trước khi thoát
	time.Sleep(5 * time.Second)
}
