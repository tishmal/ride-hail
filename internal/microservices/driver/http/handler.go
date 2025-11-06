package driver

import (
	"encoding/json"
	"fmt"
	"ride-hail/internal/shared/logger"

	"github.com/rabbitmq/amqp091-go"
)

type RideMsg struct {
	RideID   string `json:"ride_id"`
	RideType string `json:"ride_type"`
}

func ConsumeRides(ch *amqp091.Channel) {
	log := logger.New("Driver")
	q, _ := ch.QueueDeclare("driver_matching", true, false, false, false, nil)
	_ = ch.QueueBind(q.Name, "ride.request.*", "ride_topic", false, nil)

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	go func() {
		for msg := range msgs {
			var r RideMsg
			_ = json.Unmarshal(msg.Body, &r)
			log.Info("driver-service", "INFO", "ride_received", fmt.Sprintf("Ride %s received", r.RideID))

			// Отправляем ответ (мнимый "accepted")
			resp := map[string]any{
				"ride_id":   r.RideID,
				"driver_id": "driver-123",
				"accepted":  true,
			}
			b, _ := json.Marshal(resp)
			_ = ch.Publish("driver_topic", "driver.response."+r.RideID, false, false,
				amqp091.Publishing{ContentType: "application/json", Body: b})
			log.Info("driver-service", "INFO", "driver_response_sent", r.RideID)
		}
	}()
}
