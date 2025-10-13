BEGIN;

CREATE TABLE ride_event_type (
    value TEXT PRIMARY KEY
);
INSERT INTO ride_event_type (value) VALUES
('RIDE_REQUESTED'),
('RIDE_MATCHED'),
('RIDE_EN_ROUTE'),
('RIDE_ARRIVED'),
('RIDE_STARTED'),
('RIDE_COMPLETED'),
('RIDE_CANCELLED');

CREATE TABLE ride_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    ride_id UUID NOT NULL REFERENCES rides(id),
    event_type TEXT NOT NULL REFERENCES ride_event_type(value),
    actor_id UUID REFERENCES users(id),
    metadata JSONB DEFAULT '{}'::JSONB
);

CREATE INDEX idx_ride_events_ride_id ON ride_events(ride_id);

COMMIT;
