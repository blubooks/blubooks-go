-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS books
(
    id             CHAR(27)     NOT NULL,
    collection_id  CHAR(27)     NOT NULL,
    title          VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (collection_id)
        REFERENCES collections(id)
        ON DELETE CASCADE    
);

INSERT INTO books (id, collection_id,title,created_at) VALUES("2I9XSLmWimWTepVU078Poe1Qx6w", "2I9X3kNcNc4MDiFs5UVMyxSdNOH", "Buch1", NOW());
COMMIT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS books;