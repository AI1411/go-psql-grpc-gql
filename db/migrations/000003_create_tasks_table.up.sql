create table if not exists public.tasks
(
    id            serial PRIMARY KEY,
    title         varchar(64)             not null,
    "description" text                    null,
    due_date      timestamp default now() null,
    completed     bool      default false not null,
    user_id       integer                 null,
    "status"      varchar(64)             not null,
    created_at    timestamp default now() not null,
    updated_at    timestamp default now() not null
);

comment on table  public.tasks is 'ユーザ情報';
comment on column public.tasks.id            is 'ID';
comment on column public.tasks.title         is 'タイトル';
comment on column public.tasks."description" is '説明';
comment on column public.tasks.due_date      is '期限日時';
comment on column public.tasks.completed     is '完了フラグ';
comment on column public.tasks.user_id       is 'ユーザID';
comment on column public.tasks."status"      is 'ステータス';
comment on column public.tasks.created_at    is '作成日時';
comment on column public.tasks.updated_at    is '更新日時';