package rpc_server

import (
	"github.com/streadway/amqp"
	"fmt"
	"time"
	"log"
	"go/ast"
)

type RabbitConfig struct {
	User      string
	Password  string
	Host      string
	Port      string
	QueueName string
}

type RPCServer struct {
	config RabbitConfig
}

func (srv *RPCServer) Init (config RabbitConfig) {
	srv.config = config
}

func (srv *RPCServer) Start () {
	conn, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			srv.config.User,
			srv.config.Password,
			srv.config.Host,
			srv.config.Port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		srv.config.QueueName, // name
		false,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//err = ch.Qos(
	//	10,     // prefetch count
	//	0,     // prefetch size
	//	false, // global
	//)
	//failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			go func(d amqp.Delivery) {
				//log.Printf(" [.] %s", string(d.Body))
				response := "OK"
				time.Sleep(5 * time.Second)

				err = ch.Publish(
					"",        // exchange
					d.ReplyTo, // routing key
					false,     // mandatory
					false,     // immediate
					amqp.Publishing{
						ContentType:   "application/json",
						CorrelationId: d.CorrelationId,
						Body:          []byte(response),
					})
				failOnError(err, "Failed to publish a message")

				d.Ack(false)
			}(d)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
