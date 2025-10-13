BEGIN;
DROP INDEX IF EXISTS idx_coordinates_current;
DROP INDEX IF EXISTS idx_coordinates_entity;
DROP TABLE IF EXISTS coordinates;
COMMIT;
