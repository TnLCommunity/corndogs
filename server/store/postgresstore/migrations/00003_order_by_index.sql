-- +goose Up
drop index archived_task_idx_current_state;
CREATE INDEX IF NOT EXISTS task_idx_priority ON tasks (priority);
CREATE INDEX IF NOT EXISTS task_idx_priority_desc ON tasks (priority DESC);
CREATE INDEX IF NOT EXISTS task_idx_update_time_desc ON tasks (update_time DESC);

-- +goose Down
CREATE INDEX archived_task_idx_current_state ON archived_tasks (current_state);
drop index task_idx_priority;
drop index task_idx_priority_desc;
drop index task_idx_update_time_desc;
