create table if not exists like(
    id serial primary key,
    user_id int REFERENCES users(id) NOT NULL,
    post_id int REFERENCES post(id) NOT NULL
)