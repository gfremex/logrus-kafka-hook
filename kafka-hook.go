package logrus_kafka_hook

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/Sirupsen/logrus"
)

type KafkaHook struct {
	// Id of the hook
	id string

	// Log levels allowed
	levels []logrus.Level

	// Log entry formatter
	formatter logrus.Formatter

	// sarama.AsyncProducer
	producer sarama.AsyncProducer
}

// Create a new KafkaHook.
func NewKafkaHook(id string, levels []logrus.Level, formatter logrus.Formatter, brokers []string) (*KafkaHook, error) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	kafkaConfig.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	kafkaConfig.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokers, kafkaConfig)

	if err != nil {
		return nil, err
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			log.Printf("Failed to send log entry to kafka: %v\n", err)
		}
	}()

	hook := &KafkaHook{
		id,
		levels,
		formatter,
		producer,
	}

	return hook, nil
}

func (hook *KafkaHook) Id() string {
	return hook.id
}

func (hook *KafkaHook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *KafkaHook) Fire(entry *logrus.Entry) error {
	// Check time for partition key
	var partitionKey sarama.ByteEncoder

	// Get field time
	t, _ := entry.Data["time"].(time.Time)

	// Convert it to bytes
	b, err := t.MarshalBinary()

	if err != nil {
		return err
	}

	partitionKey = sarama.ByteEncoder(b)

	// Format before writing
	b, err = hook.formatter.Format(entry)

	if err != nil {
		return err
	}

	value := sarama.ByteEncoder(b)

	hook.producer.Input() <- &sarama.ProducerMessage{
		Key:   partitionKey,
		Topic: hook.Id(),
		Value: value,
	}

	return nil
}
