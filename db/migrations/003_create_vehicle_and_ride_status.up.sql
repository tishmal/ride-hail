BEGIN;

CREATE TABLE vehicle_type (
    value TEXT PRIMARY KEY
);
INSERT INTO vehicle_type (value) VALUES
('ECONOMY'),
('PREMIUM'),
('XL');

CREATE TABLE ride_status (
    value TEXT PRIMARY KEY
);
INSERT INTO ride_status (value) VALUES
('REQUESTED'),
('MATCHED'),
('EN_ROUTE'),
('ARRIVED'),
('IN_PROGRESS'),
('COMPLETED'),
('CANCELLED');

COMMIT;