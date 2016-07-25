CREATE TABLE `locations` (
`id` INTEGER PRIMARY KEY AUTOINCREMENT,
`name` VARCHAR(64),
`description` TEXT NULL,
`address` TEXT
)

CREATE TABLE `prices` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `locationId` INTEGER,
  `price` INTEGER,
  `date` TIMESTAMP,
  FOREIGN KEY(locationId) REFERENCES locations(id)
)
