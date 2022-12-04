\c books_psql
DO
$do$
    BEGIN
        IF EXISTS (
                SELECT FROM pg_catalog.pg_roles
                WHERE  rolname = 'books_psql') THEN

            RAISE NOTICE 'Role "books_psql" already exists. Skipping.';
        ELSE
            CREATE ROLE books_psql LOGIN PASSWORD 'books_psql_pass';
        END IF;
    END
$do$;

alter user books_psql with superuser;

grant all privileges on database books_psql to books_psql;

create table authors (
    id serial primary key,
    name varchar(255) not null,
    surname varchar(255) not null,
    birth_country varchar(255) not null
);

create table books (
    id serial primary key,
    title varchar(255) not null,
    genre varchar(255) not null,
    author_id int not null,
    year int not null,
    pages int not null,
    foreign key (author_id) references authors(id)
);
