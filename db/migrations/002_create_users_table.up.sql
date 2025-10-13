BEGIN;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    email VARCHAR(100) UNIQUE NOT NULL,
    role TEXT NOT NULL REFERENCES roles(value),
    status TEXT NOT NULL REFERENCES user_status(value) DEFAULT 'ACTIVE',
    password_hash TEXT NOT NULL,
    attrs JSONB DEFAULT '{}'::JSONB
);

COMMIT;
