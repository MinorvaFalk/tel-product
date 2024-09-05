-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id uuid DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    description varchar(255),
    price numeric(20,2) NOT NULL,
    variety varchar(255),
    rating numeric(2,1) DEFAULT 0,
    stock int NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz,
    deleted_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
