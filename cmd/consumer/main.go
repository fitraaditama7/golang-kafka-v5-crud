package main

import (
	"golang-kafka-v5-crud/consumer"
	"golang-kafka-v5-crud/helper"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func main() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := helper.GetKafkaConfig("", "")

	consumers, err := sarama.NewConsumer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Error creating kafka consumer got error: %v", err)
	}
	defer func() {
		if err := consumers.Close(); err != nil {
			logrus.Fatal(err)
			return
		}
	}()

	kafkaConsumer := consumer.KafkaConsumer{
		Consumer: consumers,
	}

	signals := make(chan os.Signal, 1)
	kafkaConsumer.Consume([]string{"test_topic"}, signals)
}
