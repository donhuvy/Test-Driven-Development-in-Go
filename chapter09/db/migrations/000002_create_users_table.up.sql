BEGIN;
CREATE TABLE IF NOT EXISTS users
(
    id        VARCHAR(50) PRIMARY KEY,
    name      VARCHAR(50) NOT NULL,
    address   VARCHAR(50) NOT NULL,
    post_code VARCHAR(50) NOT NULL,
    country   VARCHAR(50) NOT NULL
);
COMMIT;