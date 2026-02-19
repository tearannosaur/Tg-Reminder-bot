CREATE TABLE reminders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    reminder VARCHAR(500)
);