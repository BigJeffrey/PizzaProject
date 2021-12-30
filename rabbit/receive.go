package rabbit

import (
	"log"

	"github.com/streadway/amqp"
)

func ReceiveRabbitMessage() {

	connRabbit, err := amqp.Dial("amqp://@130.61.54.93:5672/")
	if err != nil {
		log.Fatal(err)
	}

	defer connRabbit.Close()

	chRabbit, err := connRabbit.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer chRabbit.Close()

	q, err := chRabbit.QueueDeclare(
		"whatsup", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := chRabbit.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	<-forever
}
