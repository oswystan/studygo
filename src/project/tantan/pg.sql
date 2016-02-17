---------------------------------------------------------------
-- 
-- initialize the database
--  1. create database;
--  2. create tables and indexes;
--
---------------------------------------------------------------

drop database if exists socialdb;
create database socialdb;
\c socialdb;

drop user if exists pgtest;
create user pgtest with CREATEDB LOGIN PASSWORD '123456';

create table users
(
    id bigserial not null,
    name varchar(128) not null,
    primary key(id)
);
create unique index users_name on users (name);

create table relationships
(
    peer1 bigint not null,
    peer2 bigint not null,
    relation1 int not null,
    relation2 int not null,
    primary key(peer1, peer2)
);
create index rs_peer1 on relationships (peer1, relation1);
create index rs_peer2 on relationships (peer2, relation2);

-- insert test data into tables
-- insert into users (name) values ('bob');
-- insert into users (name) values ('tina');
-- insert into users (name) values ('tata');
-- select * from users;

-- insert into relationships values(1, 2, 1, 1);
-- insert into relationships values(1, 3, 1, 2);
-- select * from relationships where (peer1 = 1 and peer2 = 2) or (peer1=2 and peer2=1);

grant all PRIVILEGES on all tables in schema public to pgtest;
grant all PRIVILEGES on all sequences in schema public to pgtest;
