ALTER TABLE users
    DROP CONSTRAINT users_username_key,
    DROP CONSTRAINT users_email_key;

ALTER TABLE users add google_id varchar;
ALTER TABLE users ALTER COLUMN password DROP NOT NULL;

