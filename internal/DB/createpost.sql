CREATE TABLE if not exists posts (
    id serial PRIMARY KEY,
    user_id int REFERENCES users(id) NOT NULL,
    content TEXT NOT NULL
);