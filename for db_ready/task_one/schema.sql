CREATE TABLE IF NOT EXISTS authors (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         birth_year INTEGER,
                         country VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(200) NOT NULL,
                       author_id INTEGER REFERENCES authors(id),
                       publication_year INTEGER,
                       isbn VARCHAR(13) UNIQUE,
                       price DECIMAL(10,2)
);

CREATE TABLE IF NOT EXISTS readers (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         email VARCHAR(100) UNIQUE NOT NULL,
                         registration_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS loans (
                       id SERIAL PRIMARY KEY,
                       book_id INTEGER REFERENCES books(id),
                       reader_id INTEGER REFERENCES readers(id),
                       loan_date DATE DEFAULT CURRENT_DATE,
                       return_date DATE,
                       UNIQUE(book_id, reader_id, loan_date)
);