CREATE TABLE IF NOT EXISTS users(
  id TEXT NOT NULL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);