create table if not exists public.users
(
    id         serial PRIMARY KEY,
    name       varchar(64) not null,
    email      varchar(64) not null,
    password   varchar(64) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

comment on table public.users is 'ユーザ情報';
comment on column public.users.id is 'ID';
comment on column public.users.name is '名前';
comment on column public.users.email is 'メールアドレス';
comment on column public.users.password is 'パスワード';
comment on column public.users.created_at is '作成日時';
comment on column public.users.updated_at is '更新日時';