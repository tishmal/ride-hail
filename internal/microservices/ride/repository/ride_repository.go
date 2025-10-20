package repository

// import (
// 	"context"
// 	"fmt"

// 	//"ride-hail/internal/pkg/models"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type RideRepository struct {
// 	pool *pgxpool.Pool
// }

// func NewRideRepository(pool *pgxpool.Pool) *RideRepository {
// 	return &RideRepository{pool: pool}
// }

// // CreateRide inserts ride and returns ride id and ride number.
// // It expects to be called inside a transaction or will start one if tx nil.
// func (r *RideRepository) CreateRide(ctx context.Context, tx pgx.Tx, req *models.CreateRideRequest) (rideID string, rideNumber string, err error) {
// 	// If no tx provided, start a short-lived tx
// 	ownsTx := false
// 	if tx == nil {
// 		ownsTx = true
// 		conn, err := r.pool.Acquire(ctx)
// 		if err != nil {
// 			return "", "", fmt.Errorf("acquire conn: %w", err)
// 		}
// 		defer conn.Release()
// 		tx, err = conn.Begin(ctx)
// 		if err != nil {
// 			return "", "", fmt.Errorf("begin tx: %w", err)
// 		}
// 	}

// 	// SQL: insert ride, use gen_random_uuid() and generate ride_number with sequence or pattern.
// 	// Here we use returning id and a generated ride_number using to_char(now(), 'YYYYMMDD') || '_' || nextval(...)
// 	// Assumes sequence ride_number_seq exists; alternatively generate in app.
// 	const insertSQL = `
// 		INSERT INTO rides (
// 			passenger_id,
// 			vehicle_type,
// 			status,
// 			requested_at,
// 			pickup_coordinate_id,
// 			destination_coordinate_id,
// 			estimated_fare,
// 			ride_number
// 		)
// 		VALUES (
// 			$1, $2, 'REQUESTED', now(), NULL, NULL, NULL,
// 			CONCAT('RIDE_', to_char(now(),'YYYYMMDD'), '_', nextval('ride_number_seq'))
// 		)
// 		RETURNING id::text, ride_number
// 	`

// 	row := tx.QueryRow(ctx, insertSQL,
// 		req.PassengerID,
// 		req.RideType,
// 	)

// 	if err := row.Scan(&rideID, &rideNumber); err != nil {
// 		if ownsTx {
// 			_ = tx.Rollback(ctx)
// 		}
// 		return "", "", fmt.Errorf("insert ride: %w", err)
// 	}

// 	if ownsTx {
// 		if err := tx.Commit(ctx); err != nil {
// 			return "", "", fmt.Errorf("commit tx: %w", err)
// 		}
// 	}

// 	return rideID, rideNumber, nil
// }

// // --- helper: make sure sequence exists (run in migrations) ---
// // CREATE SEQUENCE IF NOT EXISTS ride_number_seq START 1;
