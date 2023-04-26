CREATE DATABASE booksdb;

CREATE TABLE IF NOT EXISTS book (
  id SERIAL NOT NULL,
  title varchar(255) NOT NULL,
  author varchar(255) NOT NULL,
  description text NOT NULL,
  price decimal(10,2) NOT NULL,
  created date NOT NULL,
  modified date NOT NULL,
  PRIMARY KEY (id)
);

-- Path: db\books.sql

INSERT INTO book (title, author, description, price, created, modified) VALUES ('JOKOWI', 'J.R.R. Tolkien', 'aaaaa', 12.99, '1945-01-01', '2017-01-01');
INSERT INTO book (title, author, description, price, created, modified) VALUES ('The adams', 'gtw', 'asep na kumaha', 19.99, '2017-01-01', '2017-01-01');

