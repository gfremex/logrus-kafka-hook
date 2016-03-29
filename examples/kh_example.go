package main

import (
	lkh "github.com/gfremex/logrus-kafka-hook"
	"github.com/Sirupsen/logrus"
	"time"
)

func main() {
	// Create a new KafkaHook
	hook, err := lkh.NewKafkaHook(
		"kh",
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel},
		&logrus.JSONFormatter{},
		[]string{"192.168.60.5:9092", "192.168.60.6:9092", "192.168.60.7:9092"},
	)

	if err != nil {
		panic(err)
	}

	// Create a new logrus.Logger
	logger := logrus.New()

	// Add hook to logger
	logger.Hooks.Add(hook)

	// Add topics
	l := logger.WithField("topics", []string{"topic_1", "topic_2", "topic_3"})

	// Send message to logger
	l.Debug("This must not be logged")

	l.Info("This is an Info msg")

	l.Warn("This is a Warn msg")

	l.Error("This is an Error msg")

	// Ensure log messages were written to Kafka
	time.Sleep(time.Second)
}
