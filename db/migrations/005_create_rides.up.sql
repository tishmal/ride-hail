BEGIN;

CREATE TABLE rides (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    ride_number VARCHAR(50) UNIQUE NOT NULL,
    passenger_id UUID NOT NULL REFERENCES users(id),
    driver_id UUID REFERENCES users(id),
    vehicle_type TEXT REFERENCES vehicle_type(value),
    status TEXT REFERENCES ride_status(value),
    priority INTEGER DEFAULT 1 CHECK (priority BETWEEN 1 AND 10),
    requested_at TIMESTAMPTZ DEFAULT now(),
    matched_at TIMESTAMPTZ,
    arrived_at TIMESTAMPTZ,
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    cancelled_at TIMESTAMPTZ,
    cancellation_reason TEXT,
    estimated_fare DECIMAL(10,2),
    final_fare DECIMAL(10,2),
    pickup_coordinate_id UUID REFERENCES coordinates(id),
    destination_coordinate_id UUID REFERENCES coordinates(id)
);

CREATE INDEX idx_rides_status ON rides(status);

COMMIT;
