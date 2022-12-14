package brokerService

import (
	"encoding/json"
	"log"
	"multiverse/notifier/models"
	"os"

	"github.com/streadway/amqp"
)

type BrokerService interface {
	StartConsuming() (chan models.Task, chan error, error)
	CloseConnection() error
	CloseChannel() error
}

func NewBrokerService() BrokerService {
	return &RabbitMQBrokerService{}
}

type RabbitMQBrokerService struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (r *RabbitMQBrokerService) StartConsuming() (taskChann chan models.Task, errChann chan error, err error) {
	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	amqpQueueName := os.Getenv("AMQP_QUEUE_NAME")
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		return
	}
	r.Connection = connectRabbitMQ
	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return
	}
	r.Channel = channelRabbitMQ
	queue, err := r.Channel.QueueDeclare(
		amqpQueueName, // name
		true,          // durable (will survive server restarts)
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return
	}
	err = r.Channel.Qos( // it is for fair dispatch and means if there is no free workers, the message will be put in the queue and will be delivered to the next worker.
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return
	}

	// Subscribing to QueueService1 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		queue.Name, // queue name
		"",         // consumer
		false,      // auto-ack == > message.Ack(false) ==> message will be removed
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // arguments
	)
	if err != nil {
		return
	}

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")
	taskChann = make(chan models.Task, 10)
	errChann = make(chan error, 2)
	go func() {
		for message := range messages {
			var task models.Task
			err := json.Unmarshal(message.Body, &task)
			if err != nil {
				log.Println("Error:", err)
				errChann <- err
				return
			}
			message.Ack(false)
			taskChann <- task
		}
	}()
	return
}

func (r *RabbitMQBrokerService) CloseConnection() error {
	return r.Connection.Close()
}

func (r *RabbitMQBrokerService) CloseChannel() error {
	return r.Channel.Close()
}
