-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_code (
    uc_user_id UUID PRIMARY KEY, 
    uc_code_id UUID, 
    uc_path TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS file_code;
-- +goose StatementEnd
