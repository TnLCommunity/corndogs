-- +goose Up
create table tasks (
   uuid UUID primary key,
   queue text not null,
   current_state text not null,
   auto_target_state text not null,
   submit_time bigint not null,
   update_time bigint not null,
   timeout bigint not null,
   payload bytea not null
);

CREATE INDEX task_idx_queue ON tasks (queue);
CREATE INDEX task_idx_current_state ON tasks (current_state);
CREATE INDEX task_idx_submit_time ON tasks (submit_time);
CREATE INDEX task_idx_update_time ON tasks (update_time);

-- +goose Down
drop table tasks;
drop index task_idx_queue;
drop index task_idx_current_state;
drop index task_idx_submit_time;
drop index task_idx_update_time;