-- +goose Up
-- +goose StatementBegin
create table signals(
    ID bigserial primary key,
    address varchar(15) not null,
    port integer
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table signals;
-- +goose StatementEnd
