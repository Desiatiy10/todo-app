CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY, 
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_lists
(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text
);

CREATE TABLE IF NOT EXISTS todo_items
(
    id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text,
    done boolean NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS users_lists
(
    user_id int REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    list_id int REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL,
    created_at timestamptz DEFAULT now(),
    PRIMARY KEY (user_id, list_id)
);

CREATE TABLE IF NOT EXISTS lists_items
(
    item_id int REFERENCES todo_items(id) ON DELETE CASCADE NOT NULL,
    list_id int REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL,
    created_at timestamptz DEFAULT now(),
    PRIMARY KEY (item_id, list_id)
);
