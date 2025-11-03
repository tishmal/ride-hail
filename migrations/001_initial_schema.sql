BEGIN;

-- Активировать PostGIS (если нужно)
CREATE EXTENSION IF NOT EXISTS postgis;

-- Справочники / enum-типы
CREATE TABLE roles (value TEXT PRIMARY KEY);
INSERT INTO roles (value) VALUES ('PASSENGER'), ('DRIVER'), ('ADMIN');

CREATE TABLE user_status (value TEXT PRIMARY KEY);
INSERT INTO user_status (value) VALUES ('ACTIVE'), ('INACTIVE'), ('BANNED');

CREATE TABLE driver_status (value TEXT PRIMARY KEY);
INSERT INTO driver_status (value) VALUES ('OFFLINE'), ('AVAILABLE'), ('BUSY'), ('EN_ROUTE');

CREATE TABLE vehicle_type (value TEXT PRIMARY KEY);
INSERT INTO vehicle_type (value) VALUES ('ECONOMY'), ('PREMIUM'), ('XL');

CREATE TABLE ride_status (value TEXT PRIMARY KEY);
INSERT INTO ride_status (value) VALUES ('REQUESTED'), ('MATCHED'), ('EN_ROUTE'), ('ARRIVED'), ('IN_PROGRESS'), ('COMPLETED'), ('CANCELLED');

CREATE TABLE ride_event_type (value TEXT PRIMARY KEY);
INSERT INTO ride_event_type (value) VALUES ('RIDE_REQUESTED'), ('DRIVER_MATCHED'), ('DRIVER_ARRIVED'), ('RIDE_STARTED'), ('RIDE_COMPLETED'), ('RIDE_CANCELLED'), ('STATUS_CHANGED'), ('LOCATION_UPDATED'), ('FARE_ADJUSTED');

-- Функции (до триггеров)
CREATE OR REPLACE FUNCTION update_coordinates_geom()
RETURNS trigger AS $$
BEGIN
  NEW.geom := ST_SetSRID(ST_MakePoint(NEW.longitude, NEW.latitude), 4326);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Таблица users
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  email VARCHAR(100) UNIQUE NOT NULL,
  role TEXT REFERENCES roles(value) NOT NULL,
  status TEXT REFERENCES user_status(value) NOT NULL DEFAULT 'ACTIVE',
  password_hash TEXT NOT NULL,
  attrs JSONB DEFAULT '{}'::JSONB
);

-- Таблица coordinates
CREATE TABLE coordinates (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  entity_id UUID NOT NULL,
  entity_type VARCHAR(50) NOT NULL CHECK (entity_type IN ('driver','passenger')),
  address TEXT NOT NULL,
  latitude DECIMAL(10,8) NOT NULL CHECK (latitude BETWEEN -90 AND 90),
  longitude DECIMAL(11,8) NOT NULL CHECK (longitude BETWEEN -180 AND 180),
  fare_amount DECIMAL(10,2) CHECK (fare_amount >= 0),
  distance_km DECIMAL(8,2) CHECK (distance_km >= 0),
  duration_minutes INTEGER CHECK (duration_minutes >= 0),
  is_current BOOLEAN DEFAULT TRUE,
  geom GEOMETRY(Point, 4326)
);

CREATE TRIGGER trigger_update_coordinates_geom
BEFORE INSERT OR UPDATE OF latitude, longitude
ON coordinates
FOR EACH ROW
EXECUTE FUNCTION update_coordinates_geom();

-- Индексы координат
CREATE INDEX idx_coordinates_entity ON coordinates(entity_id, entity_type);
CREATE INDEX idx_coordinates_current ON coordinates(entity_id, entity_type) WHERE is_current = TRUE;
CREATE INDEX idx_coordinates_geom ON coordinates USING GIST (geom);

-- Таблица drivers
CREATE TABLE drivers (
  id UUID PRIMARY KEY REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  license_number VARCHAR(50) UNIQUE NOT NULL,
  vehicle_type TEXT REFERENCES vehicle_type(value),
  vehicle_attrs JSONB,
  rating DECIMAL(3,2) DEFAULT 5.0 CHECK (rating BETWEEN 1.0 AND 5.0),
  total_rides INTEGER DEFAULT 0 CHECK (total_rides >= 0),
  total_earnings DECIMAL(10,2) DEFAULT 0 CHECK (total_earnings >= 0),
  status TEXT REFERENCES driver_status(value),
  is_verified BOOLEAN DEFAULT FALSE
);
CREATE INDEX idx_drivers_status ON drivers(status);

-- Таблица driver_sessions
CREATE TABLE driver_sessions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  driver_id UUID REFERENCES drivers(id),
  started_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  ended_at TIMESTAMPTZ,
  total_rides INTEGER DEFAULT 0,
  total_earnings DECIMAL(10,2) DEFAULT 0
);

-- Таблица rides
CREATE TABLE rides (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  ride_number VARCHAR(50) UNIQUE NOT NULL,
  passenger_id UUID NOT NULL REFERENCES users(id),
  driver_id UUID REFERENCES drivers(id),
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

-- Таблица ride_events
CREATE TABLE ride_events (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  ride_id UUID REFERENCES rides(id),
  event_type TEXT REFERENCES ride_event_type(value),
  event_data JSONB NOT NULL
);

-- Таблица location_history
CREATE TABLE location_history (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  coordinate_id UUID REFERENCES coordinates(id),
  driver_id UUID REFERENCES drivers(id),
  latitude DECIMAL(10,8) NOT NULL CHECK (latitude BETWEEN -90 AND 90),
  longitude DECIMAL(11,8) NOT NULL CHECK (longitude BETWEEN -180 AND 180),
  accuracy_meters DECIMAL(6,2),
  speed_kmh DECIMAL(5,2),
  heading_degrees DECIMAL(5,2) CHECK (heading_degrees BETWEEN 0 AND 360),
  recorded_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  ride_id UUID REFERENCES rides(id)
);

COMMIT;
