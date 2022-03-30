-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_code (
    uc_code_id UUID PRIMARY KEY, 
    uc_user_id UUID, 
    uc_path TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_code;
-- +goose StatementEnd
