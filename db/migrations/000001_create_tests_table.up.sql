CREATE TABLE IF NOT EXISTS tests
(
    id   serial PRIMARY KEY,
    name varchar(64) not null
);

comment on table tests is 'テスト';
comment on column tests.id is 'ID';
comment on column tests.name is '名前';