create table posts
(
    uuid        varchar
        constraint posts_pk
            primary key,
    post_data   json,
    user_id     bigserial not null
        constraint posts_users_id_fk
            references users
            on delete cascade,
    name        varchar   not null,
    description varchar,
    public      bool default false,
    updated_at  timestamp not null
);

comment on table posts is 'Заметки пользователей';

