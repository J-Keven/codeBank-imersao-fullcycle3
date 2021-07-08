package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *ckafka.Producer
}

func NewProducerKafka() KafkaProducer {
	return KafkaProducer{}
}

func (k *KafkaProducer) SetupProducer(bootstrapSerever string) {
	k.Producer, _ = ckafka.NewProducer(&ckafka.ConfigMap{"bootstrap.servers": bootstrapSerever})
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
