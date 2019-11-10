package libs

import (
	"golang-kafka-v5-crud/helper"
	"golang-kafka-v5-crud/producer"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func kafkaProducer(msg string) {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := helper.GetKafkaConfig("", "")

	producers, err := sarama.NewSyncProducer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
		return
	}
	defer func() {
		if err := producers.Close(); err != nil {
			logrus.Errorf("Unable to stop kafka producer got error: %v", err)
			return
		}
	}()

	logrus.Infof("Success create kafka sync producer")
	kafka := producer.KafkaProducer{
		Producer: producers,
	}

	err = kafka.SendMessage("test_topic", msg)
	if err != nil {
		panic(err)
	}
}
