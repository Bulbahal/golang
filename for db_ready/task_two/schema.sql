CREATE TABLE IS NOT EXISTS students (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          "group" VARCHAR(20) NOT NULL,
                          birth_year INTEGER
);

CREATE TABLE IS NOT EXISTS courses (
                         id SERIAL PRIMARY KEY,
                         title VARCHAR(100) NOT NULL,
                         credits INTEGER
);

CREATE TABLE IS NOT EXISTS enrollments (
                             id SERIAL PRIMARY KEY,
                             student_id INTEGER REFERENCES students(id) ON DELETE CASCADE,
                             course_id INTEGER REFERENCES courses(id) ON DELETE CASCADE,
                             grade INTEGER CHECK (grade BETWEEN 1 AND 5),
                             paid BOOLEAN DEFAULT FALSE
);

CREATE TABLE IS NOT EXISTS  payments (
                          id SERIAL PRIMARY KEY,
                          enrollment_id INTEGER REFERENCES enrollments(id) ON DELETE CASCADE,
                          amount NUMERIC(10,2) NOT NULL,
                          paid_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_students_group ON students("group");
CREATE INDEX idx_enrollments_grade ON enrollments(grade);