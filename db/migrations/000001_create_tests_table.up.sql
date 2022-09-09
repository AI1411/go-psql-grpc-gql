create table public.tests
(
    id   serial PRIMARY KEY,
    name varchar(64) not null
);

comment on table public.tests is 'テスト';
comment on column public.tests.id is 'ID';
comment on column public.tests.name is '名前';