-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS clients
(
    id             CHAR(27)     NOT NULL,
    title           VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (ID)
);

INSERT INTO clients (id,title,created_at) VALUES("2I9WbBoMEDnI7pxlmxfL1rMmGWK", "Demo Mandant", NOW());
COMMIT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS clients;