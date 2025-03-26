-- +goose Up
-- +goose StatementBegin
CREATE TABLE ping_results (
    id BIGSERIAL PRIMARY KEY,
    result BOOLEAN NOT NULL,
    statistic TEXT,
    signal_id BIGINT REFERENCES signals(id),
    user_id BIGINT REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table ping_results;
-- +goose StatementEnd
