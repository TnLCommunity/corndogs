-- +goose Up
create table tasks (
                       uuid UUID primary key,
                       queue text not null,
                       current_state text not null,
                       auto_target_state text not null,
                       submit_time bigint not null,
                       update_time bigint not null,
                       timeout bigint not null,
                       payload bytea
);

CREATE INDEX task_idx_queue ON tasks (queue);
CREATE INDEX task_idx_current_state ON tasks (current_state);
CREATE INDEX task_idx_submit_time ON tasks (submit_time);
CREATE INDEX task_idx_update_time ON tasks (update_time);

create table archived_tasks (
                       uuid UUID primary key,
                       queue text not null,
                       current_state text not null,
                       auto_target_state text not null,
                       submit_time bigint not null,
                       update_time bigint not null
);

CREATE INDEX archived_task_idx_queue ON archived_tasks (queue);
CREATE INDEX archived_task_idx_current_state ON archived_tasks (current_state);
CREATE INDEX archived_task_idx_submit_time ON archived_tasks (submit_time);
CREATE INDEX archived_task_idx_update_time ON archived_tasks (update_time);

-- +goose Down
drop table tasks;
drop index task_idx_queue;
drop index task_idx_current_state;
drop index task_idx_submit_time;
drop index task_idx_update_time;
drop table archived_tasks;
drop index archived_task_idx_queue;
drop index archived_task_idx_current_state;
drop index archived_task_idx_submit_time;
drop index archived_task_idx_update_time;
