package main

import (
	"confluent-go/model"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.AllKeys())
}
func main() {
	var consumer *kafka.Consumer
	var err error

	consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": viper.Get("BOOTSTRAP_SERVERS_CONFLUENT"),
		"group.id":          viper.Get("GROUP_ID"),
		"security.protocol": viper.Get("SECURITY_PROTOCOL"),
		"sasl.mechanisms":   viper.Get("SASL_MECHANISM"),
		"sasl.username":     viper.Get("SASL_USERNAME"),
		"sasl.password":     viper.Get("SASL_PASSWORD"),
	})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{"topic_0"}, nil)

	if err != nil {
		panic(err)
	}

	for {
		var message *kafka.Message

		message, err = consumer.ReadMessage(-1)
		if err == nil {
			var order model.Order
			err := json.Unmarshal(message.Value, &order)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Message: %v\n", order)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, message)
		}
	}

	log.Fatal(consumer.Close())
}
