package helpers

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	"api-order/src/order/application/services"
	"api-order/src/order/domain/entities"
)

type RabbitMQProducer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	exchange   string
}

func NewRabbitMQProducer(exchange string) (services.IOrderProducer, error) {
	rabbitMQURL := "amqp://user:password@34.228.148.30:5672"

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchange, 
		"topic",  
		true,     
		false,    
		false,    
		false,    
		nil,     
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQProducer{
		connection: conn,
		channel:    ch,
		exchange:   exchange,
	}, nil
}

func (p *RabbitMQProducer) PublishOrderCreated(order entities.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		p.exchange,      
		"orden_topic",   
		false,           
		false,          
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Println("✅ Mensaje enviado a RabbitMQ:", string(body))
	return nil
}
