-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS collections
(
    id             CHAR(27)     NOT NULL,
    client_id      CHAR(27)     NOT NULL,
    title           VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (client_id)
        REFERENCES clients(id)
        ON DELETE CASCADE      
);

INSERT INTO collections (id,client_id,title,created_at) VALUES("2I9X3kNcNc4MDiFs5UVMyxSdNOH","2I9WbBoMEDnI7pxlmxfL1rMmGWK", "Sammlung" , NOW());
COMMIT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS collections;