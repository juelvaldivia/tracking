CREATE TABLE devices (
  id UUID PRIMARY KEY,
  type VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  mac VARCHAR(255) NOT NULL,
  alias VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  last_connection_date TIMESTAMP
);
