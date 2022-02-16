create table dials
(
    id          bigserial primary key unique,
    user_id     bigserial not null,
    constraint fk_user_dial_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE,
    name        varchar,
    description varchar,
    url         varchar   not null,
    screen      varchar,
    final       boolean default false,
    created_at  timestamp not null,
    updated_at  timestamp not null
);


