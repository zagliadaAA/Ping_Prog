-- +goose Up
-- +goose StatementBegin
CREATE TABLE signals (
    id BIGSERIAL PRIMARY KEY,
    address VARCHAR(15) NOT NULL,
    port INTEGER,
    user_id INTEGER REFERENCES users(id),
    active BOOLEAN DEFAULT true
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table signals;
-- +goose StatementEnd