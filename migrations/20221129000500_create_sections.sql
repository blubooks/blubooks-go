-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS sections
(
    id             CHAR(27)     NOT NULL,
    book_id        CHAR(27)     NOT NULL,
    section_id     CHAR(27)     null,
    sort           int          null,
    title          VARCHAR(255) NOT NULL,
    content        TEXT         NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (book_id)
        REFERENCES books(id)
        ON DELETE CASCADE,
    FOREIGN KEY (section_id)
        REFERENCES sections(id)
        ON DELETE SET null,
    INDEX sort_idx (sort),            
    FULLTEXT text_idx (title, content)  
);

INSERT INTO sections (id,sort,book_id,title,content,created_at) VALUES("2IoAJVJ7FNQaJHzX6pM2fWdzeLD", 1 ,"2I9XSLmWimWTepVU078Poe1Qx6w","Page1","# Text Inhalt\nText Test", NOW());
INSERT INTO sections (id,sort,book_id,title,content,created_at) VALUES("2IoDZHAfrLL7sfk1WR1pCXQNJX8", 2, "2I9XSLmWimWTepVU078Poe1Qx6w","Page 2","Text Inhalt 2", NOW());
COMMIT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS sections;