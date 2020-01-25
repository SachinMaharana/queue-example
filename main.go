package main

import (
	logger "SachinMaharana/twitbot/logger"
	"SachinMaharana/twitbot/rabbitmqstore"
	"SachinMaharana/twitbot/twitter"
	"os"
	"os/signal"

	"github.com/joho/godotenv"

	"syscall"
)

var log = logger.NewLogger()

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	twitterStream := make(chan string, 100)

	rabbitmqChannel, rabbitmqConn := rabbitmqstore.InitQueue()

	rabbitQueue := rabbitmqstore.DeclareQueue(rabbitmqChannel)

	go twitter.GetHashTagStream(twitterStream)
	go rabbitmqstore.Publish(rabbitmqChannel, rabbitQueue, twitterStream)

	defer rabbitmqConn.Close()
	defer rabbitmqChannel.Close()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	log.Notice("Shutting Down Gracefully", <-sigs)
}
