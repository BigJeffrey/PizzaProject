package rabbit

import (
	"os"

	"github.com/streadway/amqp"
)

func SendRabbitMessage(body string, name string) error {
	rabPass := os.Getenv("RABP")
	connRabbit, err := amqp.Dial("amqp://" + rabPass + "@130.61.54.93:5672/")
	if err != nil {
		return err
	}

	defer connRabbit.Close()

	chRabbit, err := connRabbit.Channel()
	if err != nil {
		return err
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
		return err
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
		return err
	}
	return nil
}
