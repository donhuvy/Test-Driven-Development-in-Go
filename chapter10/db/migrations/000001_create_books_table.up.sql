BEGIN;
CREATE TABLE IF NOT EXISTS books
(
   id VARCHAR (50) PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   author VARCHAR (50) NOT NULL,
   owner_id VARCHAR (50) NOT NULL,
   status VARCHAR (50) NOT NULL
);
COMMIT;