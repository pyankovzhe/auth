-- +goose Up
-- +goose StatementBegin
CREATE table accounts (
    id serial primary key,
    login varchar,
    encrypted_password varchar not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table accounts;
-- +goose StatementEnd
