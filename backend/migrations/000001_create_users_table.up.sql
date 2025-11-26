CREATE TABLE users (

  id integer unique PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
  username    varchar(40),
  email   varchar(40),
  password varchar(900)
  
);


CREATE INDEX idx_users_username ON users(username);