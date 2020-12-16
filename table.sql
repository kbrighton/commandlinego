create table todos
(
    id           serial not null
        constraint todos_pk
            primary key,
    taskname     varchar,
    description  text,
    is_complete  boolean default false,
    date_created date,
    date_updated date
);

alter table todos
    owner to postgres;

