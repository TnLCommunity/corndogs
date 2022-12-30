-- +goose Up
alter table tasks add priority int;

-- +goose Down
alter table tasks drop column priority;
