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
INSERT INTO sections (id,section_id,sort,book_id,title,content,created_at) VALUES("2IoxbR1gZRLXo1NHxJYdQfIXpDm","2IoDZHAfrLL7sfk1WR1pCXQNJX8", 1, "2I9XSLmWimWTepVU078Poe1Qx6w","Page 2 1","Text Inhalt 2 1", NOW());
INSERT INTO sections (id,section_id,sort,book_id,title,content,created_at) VALUES("2IpD8yQKEX3QsZKvsX2MatPwJR7","2IoDZHAfrLL7sfk1WR1pCXQNJX8", 2, "2I9XSLmWimWTepVU078Poe1Qx6w","Page 2 2","Text Inhalt 2 2", NOW());
INSERT INTO sections (id,section_id,sort,book_id,title,content,created_at) VALUES("2IpDCV2Pk6O16DqGHOH5CzgHrqk","2IpD8yQKEX3QsZKvsX2MatPwJR7", 2, "2I9XSLmWimWTepVU078Poe1Qx6w","Page 2 3","Text Inhalt 2 3", NOW());
INSERT INTO sections (id,section_id,sort,book_id,title,content,created_at) VALUES("2IpDI93x4ZQLP1ebyVqDPNEDhHi","2IoDZHAfrLL7sfk1WR1pCXQNJX8", 1, "2I9XSLmWimWTepVU078Poe1Qx6w","Page 3 4","Text Inhalt 2 1", NOW());
INSERT INTO sections (id,sort,book_id,title,content,created_at) VALUES("2IpD5VcM1DUOve0yZ2kpcc6moat", 3, "2I9XSLmWimWTepVU078Poe1Qx6w","Page 2","Text Inhalt 3", NOW());
   
COMMIT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS sections;