CREATE TABLE users (
    id          BIGINT  PRIMARY KEY,
    username    TEXT    NOT NULL
);

COMMENT ON TABLE users IS 'Info about users';
COMMENT ON COLUMN users.id IS 'ID of user';
COMMENT ON COLUMN users.username IS 'Nickname of user';
