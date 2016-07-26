CREATE TABLE locations (
id INTEGER PRIMARY KEY IDENTITY(1,1),
name VARCHAR(64),
description TEXT NULL,
address TEXT,
regex TEXT
);

CREATE TABLE prices (
  id INTEGER PRIMARY KEY IDENTITY(1,1),
  locationId INTEGER,
  price INTEGER,
  accessDate DATETIME,
  FOREIGN KEY(locationId) REFERENCES locations(id)
);
