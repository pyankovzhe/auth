-- +goose Up
-- +goose StatementBegin
CREATE table accounts (
    id uuid primary key default uuid_generate_v4(),
    login varchar UNIQUE,
    encrypted_password varchar not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table accounts;
-- +goose StatementEnd
