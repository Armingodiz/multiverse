package brokerService

import (
	"encoding/json"
	"multiverse/core/models"
	"os"

	"github.com/streadway/amqp"
)

type BrokerService interface {
	Publish(tack models.Task) error
}

func NewBrokerService() BrokerService {
	return &RabbitMQBrokerService{}
}

type RabbitMQBrokerService struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (r *RabbitMQBrokerService) Publish(task models.Task) error {
	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		return err
	}
	defer connectRabbitMQ.Close()

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return err
	}
	defer channelRabbitMQ.Close()
	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		"QueueService1", // queue name
		false,           // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		return err
	}
	taskBytes, err := json.Marshal(task)
	if err != nil {
		return err
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        taskBytes,
	}

	// Attempt to publish a message to the queue.
	return channelRabbitMQ.Publish(
		"",              // exchange
		"QueueService1", // queue name
		false,           // mandatory
		false,           // immediate
		message,         // message to publish
	)
}
