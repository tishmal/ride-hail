package mq

import (
	"fmt"

	"ride-hail/internal/shared/config"

	"github.com/rabbitmq/amqp091-go"
)

func ConnectRabbit(cfg *config.Config) (*amqp091.Connection, *amqp091.Channel, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User, cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host, cfg.RabbitMQ.Port)

	fmt.Println("RabbitMQ URL:", url)

	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	// Создаём exchanges
	_ = ch.ExchangeDeclare("ride_topic", "topic", true, false, false, false, nil)
	_ = ch.ExchangeDeclare("driver_topic", "topic", true, false, false, false, nil)
	_ = ch.ExchangeDeclare("location_fanout", "fanout", true, false, false, false, nil)

	return conn, ch, nil
}
