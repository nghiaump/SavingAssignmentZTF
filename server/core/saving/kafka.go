package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer() (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Thay thế địa chỉ broker của Kafka bằng địa chỉ thực tế của bạn
	brokerList := []string{"kafka:9092"}

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return &Producer{producer: producer}, nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}

func (p *Producer) Produce(topic string, value []byte) error {
	_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	})
	return err
}

func ProduceNewSavingAccountMessage(account *pb.SavingAccount) error {
	log.Println("Produce message to Kafka")
	producer, err := NewProducer()
	if err != nil {
		log.Printf("Failed to create Kafka producer: %s", err)
		return err
	}
	defer producer.Close()
	accJSON, _ := json.Marshal(account)
	if err := producer.Produce(KafkaTopicSavingAccount, accJSON); err != nil {
		log.Printf("Failed to produce message to Kafka: %s", err)
		return err
	}
	return nil
}
