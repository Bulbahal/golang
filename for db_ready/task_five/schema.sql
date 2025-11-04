CREATE TABLE IF NOT EXISTS courses(
    id SERIAL PRIMARY KEY,
    title  TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS enrollments (
    id SERIAL PRIMARY KEY,
    student_id INTEGER NOT NULL,
    course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    grade INTEGER CHECK (grade BETWEEN 1 AND 5)
    );

CREATE TABLE IF NOT EXISTS course_avg_grade (
    course_id INTEGER PRIMARY KEY REFERENCES courses(id) ON DELETE CASCADE,
    avg_grade NUMERIC(4,2)
    );

CREATE OR REPLACE FUNCTION recalc_course_avg() RETURNS TRIGGER AS $$
DECLARE
cid  INTEGER;
    avgv NUMERIC(4,2);
BEGIN
    cid := COALESCE(NEW.course_id, OLD.course_id);

SELECT ROUND(AVG(grade)::NUMERIC, 2)
INTO avgv
FROM enrollments
WHERE course_id = cid;

IF avgv IS NULL THEN
DELETE FROM course_avg_grade WHERE course_id = cid;
ELSE
        INSERT INTO course_avg_grade (course_id, avg_grade)
        VALUES (cid, avgv)
        ON CONFLICT (course_id) DO UPDATE
                                       SET avg_grade = EXCLUDED.avg_grade;
END IF;

RETURN NULL;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS enrollments_after_change ON enrollments;
CREATE TRIGGER enrollments_after_change
    AFTER INSERT OR UPDATE OR DELETE ON enrollments
    FOR EACH ROW EXECUTE FUNCTION recalc_course_avg();
