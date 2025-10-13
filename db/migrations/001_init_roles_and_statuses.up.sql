BEGIN;

CREATE TABLE roles (
    value TEXT PRIMARY KEY
);
INSERT INTO roles (value) VALUES
('PASSENGER'),
('DRIVER'),
('ADMIN');

CREATE TABLE user_status (
    value TEXT PRIMARY KEY
);
INSERT INTO user_status (value) VALUES
('ACTIVE'),
('INACTIVE'),
('BANNED');

COMMIT;
