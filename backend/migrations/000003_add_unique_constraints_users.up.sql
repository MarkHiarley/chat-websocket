DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_username;


ALTER TABLE users 
DROP CONSTRAINT IF EXISTS users_email_unique;

ALTER TABLE users 
DROP CONSTRAINT IF EXISTS users_username_unique;


ALTER TABLE users 
ADD CONSTRAINT users_email_unique UNIQUE (email);

ALTER TABLE users 
ADD CONSTRAINT users_username_unique UNIQUE (username);


CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);