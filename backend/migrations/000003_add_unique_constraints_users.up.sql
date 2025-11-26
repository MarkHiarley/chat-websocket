-- Remove índices
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_username;


ALTER TABLE users 
DROP CONSTRAINT IF EXISTS users_email_unique;

ALTER TABLE users 
DROP CONSTRAINT IF EXISTS users_username_unique;


ALTER TABLE users 
ALTER COLUMN id DROP DEFAULT;

-- Remove a sequência
DROP SEQUENCE IF EXISTS users_id_seq CASCADE;