drop table posts cascade;
drop table comments;

create table posts (
    id serial primary key,
    content text,
    author varchar(256)
);

create table comments (
    id serial primary key,
    content text,
    author varchar(256),
    post_id integer references posts(id)
);
