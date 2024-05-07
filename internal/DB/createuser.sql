create table if not exists users(
    id serial PRIMARY KEY , 
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);