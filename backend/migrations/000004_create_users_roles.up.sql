ALTER TABLE users 
ADD COLUMN role VARCHAR(50) NOT NULL DEFAULT 'user';

ALTER TABLE users
ADD constraint chk_user_role CHECK (role IN ('user', 'admin'));

CREATE INDEX idx_users_role ON users(role);