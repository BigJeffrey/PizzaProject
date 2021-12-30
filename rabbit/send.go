package rabbit

import (
	"log"

	"github.com/streadway/amqp"
)

func SendRabbitMessage(body string, name string) {
	connRabbit, err := amqp.Dial("amqp://130.61.54.93:5672/")
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
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	err = chRabbit.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatal(err)
	}
}
