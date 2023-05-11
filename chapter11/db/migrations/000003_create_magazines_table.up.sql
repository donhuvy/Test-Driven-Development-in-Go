BEGIN;
CREATE TABLE IF NOT EXISTS magazines
(
    id           VARCHAR(50) PRIMARY KEY,
    name         VARCHAR(50) NOT NULL,
    issue_number INTEGER     NOT NULL,
    owner_id     VARCHAR(50) NOT NULL,
    status       VARCHAR(50) NOT NULL
);
COMMIT;