BEGIN;
DROP INDEX IF EXISTS idx_ride_events_ride_id;
DROP TABLE IF EXISTS ride_events;
DROP TABLE IF EXISTS ride_event_type;
COMMIT;
