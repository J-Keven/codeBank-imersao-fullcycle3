package kafka

import (
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *ckafka.Producer
}

func NewProducerKafka() KafkaProducer {
	return KafkaProducer{}
}

// altera diretamente a classe
func (k *KafkaProducer) SetupProducer(bootstrapSerever string) {
	var err error
	k.Producer, err = ckafka.NewProducer(&ckafka.ConfigMap{"bootstrap.servers": bootstrapSerever})
	if err != nil {
		log.Fatalln("Error to connect in kafka", err)
	}
}

func (k *KafkaProducer) Publish(msg string, topic string) error {
	parseMessage := ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := k.Producer.Produce(&parseMessage, nil)
	if err != nil {
		return err
	}

	return nil
}
