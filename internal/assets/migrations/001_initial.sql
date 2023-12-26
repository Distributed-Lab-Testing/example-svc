-- +migrate Up

create table if not exists notes(
    id bigserial primary key,
    note text,
    created_at timestamp without time zone
);

-- +migrate Down
drop table if exists notes;
