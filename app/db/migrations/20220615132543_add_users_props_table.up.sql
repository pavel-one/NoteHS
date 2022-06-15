create table users_settings
(
    id        bigserial
        constraint users_settings_pk
            primary key,
    user_id   bigint                     not null
        constraint users_settings_users_id_fk
            references users
            on delete cascade,
    component varchar default 'NotePage' not null,
    post_id   varchar default '0'        not null
        constraint users_settings_posts_uuid_fk
            references posts
            on delete set default
);

comment on table users_settings is 'Настройки пользователей';

create unique index users_settings_id_uindex
    on users_settings (id);

create unique index users_settings_user_id_uindex
    on users_settings (user_id);

