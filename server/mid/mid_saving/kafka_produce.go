package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/golang/glog"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"os"
	"strings"
)

func ReadKafkaConfig() kafka.ConfigMap {
	// reads the client configuration from client.properties
	// and returns it as a key-value map
	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open("client.properties")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return m
}

func NewKafkaProducer() *kafka.Producer {
	conf := ReadKafkaConfig()
	p, _ := kafka.NewProducer(&conf)
	return p
}

func (handler *MidServiceHandler) ProduceNewMessage(acc *pb.SavingAccount) {
	topic := "NewSavingAccountCreated"
	byteAcc, err := json.Marshal(acc)
	if err != nil {
		glog.Errorf("ProduceNewMessage: failed to marshal savingaccount: %v", acc)
	}

	handler.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte("acc"),
		Value:          byteAcc,
	}, nil)

	// send any outstanding or buffered messages to the Kafka broker and close the connection
	//p.Flush(15 * 1000)
	//p.Close()
}
