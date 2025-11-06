package http

import (
	"encoding/json"
	"net/http"

	"ride-hail/internal/shared/logger"
	"ride-hail/internal/shared/utils"

	"github.com/rabbitmq/amqp091-go"
)

type RideRequest struct {
	PassengerID string  `json:"passenger_id"`
	PickupLat   float64 `json:"pickup_latitude"`
	PickupLng   float64 `json:"pickup_longitude"`
	RideType    string  `json:"ride_type"`
}

type RideResponse struct {
	RideID   string  `json:"ride_id"`
	Status   string  `json:"status"`
	Fare     float64 `json:"estimated_fare"`
	Message  string  `json:"message"`
	RideType string  `json:"ride_type"`
}

func HandleCreateRide(ch *amqp091.Channel) http.HandlerFunc {
	log := logger.New("RIDE")

	return func(w http.ResponseWriter, r *http.Request) {
		var req RideRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid request"})
			return
		}

		rideID, err := utils.GenerateUUID()
		if err != nil {
			return
		}
		resp := RideResponse{
			RideID:   rideID,
			Status:   "REQUESTED",
			Fare:     1000.0,
			Message:  "Ride created",
			RideType: req.RideType,
		}

		body, _ := json.Marshal(resp)
		err = ch.Publish(
			"ride_topic",
			"ride.request."+req.RideType,
			false, false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			// log.Log("ride-service", "ERROR", "publish_ride", err.Error())
		} else {
			log.Info("ride-service", "INFO", "ride_created", "Ride published to MQ")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(resp)
	}
}
