package service 

import (
	"log"
	"github.com/IBM/sarama"
)

type ServerImpl struct {} 
func NewServerImpl() ServerImpl {
	return ServerImpl{}  
}

const (
	Topic = "kafka_one"
	Group = "kakfa_group"
)
// GetConfig Create Kafka sarama config 
//	@return *sarama.Config 
func GetConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true 
	log.Println("Creating Kafka config successful")
	return config 
}

// NewClient Create Kafka sarama client 
//	@return sarama.Client 
func NewClient() sarama.Client {
	client, err := sarama.NewClient([]string{"localhost:9092"}, GetConfig())
	if err != nil {
		log.Fatal("Error creating client: ", err)
	}
	log.Printf("Creating client successful\n")
	return client
}

// GetProducer Create producer. REMEMBER TO CALL producer.Close()
//	@return sarama.SyncProducer 
func GetProducer() sarama.SyncProducer {
	client := NewClient() 
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatal("Error creating Kafka producer: ", err)
	}
	log.Println("Creating producer successful")
	return producer
}

// SendMessage send message 
//	@param line 
//	@param producer 
func SendMessage(line string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Value: sarama.StringEncoder(line),
		Topic: Topic,
	}
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal("Message send error: ", err)
	}
	// log.Printf("Message sent, partition{%v}, offset{%v}\n", partition, offset)
}
