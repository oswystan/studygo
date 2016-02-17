---------------------------------------------------------------
-- 
-- initialize the database
--  1. create database;
--  2. create tables and indexes;
--
---------------------------------------------------------------
drop database if exists social;
create database social;
\c social;

create table users
(
    id bigserial not null,
    name varchar(128) not null,
    primary key(id)
);

create table relationships
(
    peer1 bigint not null,
    peer2 bigint not null,
    relation int not null,
    primary key(peer1, peer2)
);
create index rs_peer1 on relationships (peer1, relation);
create index rs_peer2 on relationships (peer2, relation);

-- insert test data into tables
insert into users (name) values ('bob');
insert into users (name) values ('mike');
insert into users (name) values ('tina');
select * from users;

insert into relationships values(1, 2, 1);
insert into relationships values(1, 3, 1);
select * from relationships;

