create table if not exists public.products
(
    id             UUID                       NOT NULL DEFAULT gen_random_uuid(),
    name           varchar(64)                NOT NULL,
    description    text                       null,
    price          integer     default 100    NOT NULL,
    discount_price integer     default 100    NOT NULL,
    status         varchar(64) default 'sale' NOT NULL,
    user_id        integer                    NOT NULL,
    created_at     timestamp   default now()  NOT NULL,
    updated_at     timestamp   default now()  NOT NULL
);

comment on table public.products is 'ユーザ情報';
comment on column public.products.id is 'ID';
comment on column public.products.name is '商品名';
comment on column public.products.description is '商品説明';
comment on column public.products.price is '価格';
comment on column public.products.discount_price is '値引後価格';
comment on column public.products.status is '商品ステータス
    SALE = 1; // 販売中
    SOLD = 2; // 販売済
    EXHIBIT = 3 // 出品中止;';
comment on column public.products.user_id is 'ユーザID';
comment on column public.products.created_at is '作成日時';
comment on column public.products.updated_at is '更新日時';