-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id bigserial PRIMARY KEY,
    username varchar(120) NOT NULL,
    chat_id INTEGER UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd