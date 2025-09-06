/* 
############ CREATE TABLES ############
*/
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  email TEXT NOT NULL UNIQUE,
  username TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS events (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  location TEXT NOT NULL,
  dateTime TEXT NOT NULL,
  organizer INTEGER NOT NULL, -- refer to user id
  FOREIGN KEY (organizer) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS comments (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  event_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  parrent_id INTEGER,
  comment TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (event_id) REFERENCES events (id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (parrent_id) REFERENCES comments (id)
);