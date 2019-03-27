package middleware

import (
	"github.com/streadway/amqp"
	"log"
)

var AmqpConn  *amqp.Connection

func Conn() (*amqp.Connection){
	var  err error
	AmqpConn, err = amqp.Dial("amqp://rabbitmq:rabbitmq@58.87.86.221:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	return AmqpConn
}


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func Send(quene_name string,body string){
	ch, err := AmqpConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		quene_name, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//body := "Hello World!"

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

}

func SendTest(body string){
	Send("go_test",body)
}

//func main()  {
//	SendTest("test")
//}