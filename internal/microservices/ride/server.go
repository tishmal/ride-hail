package ride

import (
	"database/sql"
	"net/http"
)

type Server struct {
	DB *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{DB: db}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/rides", s.handleGetRides)
	mux.HandleFunc("/rides/create", s.handleCreateRide)

	// Оборачиваем все маршруты цепочкой middleware
	handler := CORSMiddleware(
		LoggingMiddleware(
			AuthMiddleware(mux),
		),
	)

	return handler
}

// ---------------- Handlers ----------------

func (s *Server) handleGetRides(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of rides"))
}

func (s *Server) handleCreateRide(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ride created"))
}
