CREATE TABLE IS NOT EXISTS authors (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(100) NOT NULL
);

CREATE TABLE IS NOT EXISTS books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(200) NOT NULL,
                       author_id INTEGER REFERENCES authors(id),
                       publication_year INTEGER,
                       isbn VARCHAR(13),
                       price DECIMAL(10,2)
);