create table user_settings
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
);

comment on table user_settings is 'Настройки пользователей';

create unique index user_settings_id_uindex
    on user_settings (id);

create unique index user_settings_user_id_uindex
    on user_settings (user_id);

