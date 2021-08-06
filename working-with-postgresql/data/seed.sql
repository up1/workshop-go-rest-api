CREATE TABLE users(
   id serial PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL,
   age smallint,
   email VARCHAR (50) UNIQUE NOT NULL
);

INSERT INTO users VALUES (1, 'name 01', 10, 'user01@gmail.com');
INSERT INTO users VALUES (2, 'name 02', 20, 'user02@gmail.com');
INSERT INTO users VALUES (3, 'name 03', 30, 'user03@gmail.com');