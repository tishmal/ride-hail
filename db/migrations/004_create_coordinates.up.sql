BEGIN;

CREATE TABLE coordinates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    entity_id UUID NOT NULL,
    entity_type VARCHAR(20) NOT NULL CHECK (entity_type IN ('driver', 'passenger')),
    address TEXT NOT NULL,
    latitude DECIMAL(10,8) NOT NULL CHECK (latitude BETWEEN -90 AND 90),
    longitude DECIMAL(11,8) NOT NULL CHECK (longitude BETWEEN -180 AND 180),
    fare_amount DECIMAL(10,2) CHECK (fare_amount >= 0),
    distance_km DECIMAL(8,2) CHECK (distance_km >= 0),
    duration_minutes INTEGER CHECK (duration_minutes >= 0),
    is_current BOOLEAN DEFAULT true
);

CREATE INDEX idx_coordinates_entity ON coordinates(entity_id, entity_type);
CREATE INDEX idx_coordinates_current ON coordinates(entity_id, entity_type) WHERE is_current = true;

COMMIT;
