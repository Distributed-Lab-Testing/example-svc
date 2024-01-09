-- +migrate Up

create table if not exists notes(
    id bigserial primary key,
    content text not null,
    created_at timestamp without time zone not null
);

-- +migrate Down
drop table if exists notes;
