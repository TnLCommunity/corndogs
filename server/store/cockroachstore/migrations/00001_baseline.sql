-- +goose Up
create table tasks (
   uuid UUID primary key,
   queue text not null,
   current_state text not null,
   auto_target_state text not null,
   submit_time int64 not null,
   update_time int64 not null,
   timeout int64 not null,
   payload bytes not null,

   INDEX index_queue (queue),
   INDEX index_current_state (current_state),
   INDEX index_submit_time (submit_time),
   INDEX index_update_time (update_time)
);

-- +goose Down
drop table tasks;