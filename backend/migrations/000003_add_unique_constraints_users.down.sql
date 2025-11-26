-- Remove os índices
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_username;

-- Remove as constraints de unicidade
ALTER TABLE users 
DROP CONSTRAINT IF EXISTS users_email_unique;

ALTER TABLE users 
DROP CONSTRAINT IF EXISTS users_username_unique;