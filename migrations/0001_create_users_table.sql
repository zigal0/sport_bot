CREATE TABLE users (
    id          BIGINT      PRIMARY KEY,
    username    TEXT        NOT NULL,
    created_at  TIMESTAMP   WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

COMMENT ON TABLE users IS 'Info about users';
COMMENT ON COLUMN users.id IS 'ID of user';
COMMENT ON COLUMN users.username IS 'Nickname of user';
