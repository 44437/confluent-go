package main

import (
	"confluent-go/model"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProvider(t *testing.T) {
	var producer *kafka.Producer
	var err error

	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": viper.Get("BOOTSTRAP_SERVERS_CONFLUENT"),
		"group.id":          viper.Get("GROUP_ID"),
		"security.protocol": viper.Get("SECURITY_PROTOCOL"),
		"sasl.mechanisms":   viper.Get("SASL_MECHANISM"),
		"sasl.username":     viper.Get("SASL_USERNAME"),
		"sasl.password":     viper.Get("SASL_PASSWORD"),
	})
	require.Nil(t, err)

	topicName := "topic_0"
	testOrder, err := json.Marshal(model.Order{
		Type: "order",
		Message: model.Message{
			ID:         "1234",
			Total:      123.4,
			Currency:   "TRY",
			CustomerID: "4321",
		},
	})

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName,
			Partition: kafka.PartitionAny},
		Value: testOrder}, nil)
	require.Nil(t, err)

	e := <-producer.Events()
	message := e.(*kafka.Message)
	require.Nil(t, message.TopicPartition.Error)

	producer.Close()
}
