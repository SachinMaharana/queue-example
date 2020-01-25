package main

import (
	"SachinMaharana/twitbot/rabbitmqstore"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var count int32

func main() {
	rabbitmqChannel, rabbitmqConn := rabbitmqstore.InitQueue()
	rabbitQueue := rabbitmqstore.DeclareQueue(rabbitmqChannel)

	defer rabbitmqConn.Close()
	defer rabbitmqChannel.Close()

	msgs, err := rabbitmqChannel.Consume(
		rabbitQueue.Name, // queue
		"",               // consumer
		false,            // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			count += 1
			log.Printf("Received a message: %s %d", d.Body, count)
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
