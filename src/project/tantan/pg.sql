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
    name varchar(64) not null,
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

-- grant all PRIVILEGES on all tables in schema public to pgtest;
-- grant all PRIVILEGES on all sequences in schema public to pgtest;

alter table users owner to pgtest;
alter table relationships owner to pgtest;
alter sequence users_id_seq owner to pgtest;
select * from users;

create or replace function create_rs(id1 int, id2 int, state int) returns setof relationships as $$
    declare
        p1 bigint := 0;
        p2 bigint := 0;
        r1 int := 0;
        r2 int := 0;
        userid bigint := 0;
        rec record;

    begin
        select id from users where id=id1 into userid;
        if not found then
            raise exception 'can not find user %', id1;
        end if;
        select id from users where id=id2 into userid;
        if not found then
            raise exception 'can not find user %', id2;
        end if;

        if id1 > id2 then
            p1 = id2;
            p2 = id1;
            r2 = state;
        else
            p1 = id1;
            p2 = id2;
            r1 = state;
        end if;
        
        select * from relationships where peer1 = p1 and peer2 = p2 into rec;
        if not found then
            return query insert into relationships values (p1, p2, r1, r2) returning *;
        else
            if r1 > 0 then
                return query update relationships set relation1 = r1 where peer1=p1 and peer2=p2 returning *;
            else
                return query update relationships set relation2 = r2 where peer1=p1 and peer2=p2 returning *;
            end if;
        end if;
    end;

$$ language plpgsql;

