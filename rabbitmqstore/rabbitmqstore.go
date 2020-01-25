package rabbitmqstore

import (
	logger "SachinMaharana/twitbot/logger"

	"github.com/streadway/amqp"
)

var log = logger.NewLogger()

// InitQueue ...
func InitQueue() (*amqp.Channel, *amqp.Connection) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	log.Notice("Connected to Queue")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch, conn
}

// DeclareQueue ..
func DeclareQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"twitter", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return q
}

// Publish ..
func Publish(ch *amqp.Channel, q amqp.Queue, b <-chan string) {
	for v := range b {
		err := ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(v),
			})
		failOnError(err, "Failed to publish a message")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Criticalf("%s: %s", msg, err)
	}
}
