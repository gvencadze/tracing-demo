-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(id INT PRIMARY KEY,first_name VARCHAR,last_name VARCHAR);

INSERT INTO users(id, first_name, last_name) VALUES (1, 'Alexey', 'Akulovich');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
