package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	sendemail "github.com/Nishad4140/SkillSync_NotificationService/internal/helper/sendEmail"
)

type Acknowledgement struct {
	Email    string `json:"Email"`
	UserName string `json:"UserName"`
	Title    string `json:"Title"`
}

func StartConsumingAcknowledgements() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("Acknowledgement", 0, sarama.OffsetNewest)
	fmt.Println("offset ", sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var acknowledement Acknowledgement
			err := json.Unmarshal(msg.Value, &acknowledement)
			fmt.Println("message received")
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			go func(ack Acknowledgement) {
				message := fmt.Sprintf("Congratulations! %s user accepted your interest on the request they have posted titled as %s", ack.UserName, ack.Title)
				if err := sendemail.SendEmail(ack.Email, message); err != nil {
					log.Println(err)
				}
			}(acknowledement)

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}
