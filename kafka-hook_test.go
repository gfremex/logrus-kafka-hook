package logrus_kafka_hook

import (
	"testing"
	"github.com/Sirupsen/logrus"
	"time"
)

func TestKafkaHook(t *testing.T) {
	// Create a new KafkaHook
	hook, err := NewKafkaHook(
		"kh",
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel},
		&logrus.JSONFormatter{},
		[]string{"192.168.60.5:9092", "192.168.60.6:9092", "192.168.60.7:9092"},
	)

	if err != nil {
		t.Errorf("Can not create KafkaHook: %v\n", err)
	}

	// Create a new logrus.Logger
	logger := logrus.New()

	// Add hook to logger
	logger.Hooks.Add(hook)

	t.Logf("logger: %v", logger)
	t.Logf("logger.Out: %v", logger.Out)
	t.Logf("logger.Formatter: %v", logger.Formatter)
	t.Logf("logger.Hooks: %v", logger.Hooks)
	t.Logf("logger.Level: %v", logger.Level)

	// Add topics
	l := logger.WithField("topics", []string{"topic_1", "topic_2", "topic_3"})

	l.Debug("This must not be logged")

	l.Info("This is an Info msg")

	l.Warn("This is a Warn msg")

	l.Error("This is an Error msg")

	// Ensure log messages were written to Kafka
	time.Sleep(time.Second)
}
