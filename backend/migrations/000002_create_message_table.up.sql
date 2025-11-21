CREATE TABLE messages (
  id serial PRIMARY KEY,
  sender_id INT REFERENCES users(id) ON DELETE SET NULL,
  message varchar(900),
  created_at TIMESTAMP DEFAULT NOW()

);

CREATE INDEX idx_messages_sender ON messages(sender_id);
CREATE INDEX idx_messages_created_at ON messages(created_at DESC);