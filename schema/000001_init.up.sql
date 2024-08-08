CREATE TABLE users
(
    id serial not null primary key,
    name varchar(150) not null,
    username varchar(150) not null unique,
    password_hash varchar(150) not null
);

CREATE TABLE todo_lists
(
    id serial not null unique,
    title varchar(200) not null,
    description varchar(200)
);

CREATE TABLE users_lists
(
    id serial not null primary key,
    user_id int references users (id) no delete cascade not null,
    list_id int references todo_lists (id) no delete cascade no null
);

CREATE TABLE todo_items
(
    id serial not null primary key,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false 
);

CREATE TABLE lists_items
(
    id serial not null primary key,
    item_id int references todo_items (id) no delete cascade not null, 
    list_id  int references todo_lists (id) no delete cascade not null
);


