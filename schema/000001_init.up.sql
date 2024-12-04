CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE posts
(
    id          serial                                      not null unique,
    user_id     int references users (id) on delete cascade not null,
    title       varchar(255)                                not null,
    description varchar(255)
);

CREATE TABLE refresh_tokens
(
    id         serial                                      not null unique,
    user_id    int references users (id) on delete cascade not null,
    token      varchar(255)                                not null unique,
    expires_at timestamp                                   not null
);