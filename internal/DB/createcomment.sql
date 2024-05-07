CREATE TABLE if not exists comments (
    Id serial PRIMARY KEY,
    user_id UUID REFERENCES users(id) NOT NULL,
    post_id UUID REFERENCES posts(id) NOT NULL,
    content TEXT NOT NULL
);