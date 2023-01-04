# Corndogs API Docs
This is where the API docs for corndogs lives for now. It contains the rpc endpoints available. It's mostly semantics and brief descriptions at the moment, since the [Flow](https://github.com/TnLCommunity/corndogs#flow) section in the readme covers a lot of how you would use these. See the [protos-corndogs](https://github.com/TnLCommunity/protos-corndogs) repo for specific fields and generated clients.

## General
The regular flow stuff.

### `SubmitTask`
Used for submitting a new task to a `queue`.\
Returns the created task.

### `GetTaskStateByID`
Gets a task by the `uuid`. Will return archived tasks.

### `GetNextTask`
Gets the next task from `queue` that has the same `current_state`.\
The `override_` fields override the task data *after* the states are switched.
See [State and Timeout Overrides](https://github.com/TnLCommunity/corndogs#state-and-timeout-overrides) in the readme for an example.\
Returns the next task.

### `UpdateTask`
Will update a task with the matching `uuid`, `queue`, and `current_state`. Use `new_state` to update the `current_state`.\
Returns the updated task.

### `CompleteTask`
Will complete a task with the matching `uuid`, `queue`, and `current_state`.\
Sets the `current_state` and `auto_target_state` to `completed` and archives the task.\
Returns the archived task.

### `CancelTask`
Will cancel a task with the matching `uuid`, `queue`, and `current_state`.\
Sets the `current_state` and `auto_target_state` to `canceled` and archives the task.\
Returns the archived task.

### `CleanUpTimedOut`
Will compare tasks to `at_time` to see if they're timed out. Optionally limited to a specific `queue`.\
See the [Timeouts](https://github.com/TnLCommunity/corndogs#timeouts) section in the readme for more info on how you might use timeouts.\
Returns the number of tasks `timed_out`.

---

## Metrics
For the proto based metrics stuff. 

### `GetQueues`
Returns `GetQueuesResponse` containing a list of `queues`, and the `total_task_count`.

### `GetQueueTaskCounts`
Returns `GetQueueTaskCountsResponse` containing:
- `queue_counts` map of queue name to number of tasks in that queue.
- Also returns `total_task_count`.

### `GetTaskStateCounts`
Accepts `GetTaskStateCountsRequest` with the `queue` you'd like to get the state task count for.\
Returns `GetTaskStateCountsResponse` containing:
- `queue` requested
- `count` of the total tasks in the queue
- `state_counts` map of state name to number of tasks in that state.

### `GetQueueAndStateCounts`
Returns `GetQueueAndStateCountsResponse` containing:
- `queue_and_state_counts` map of queue name to `QueueAndStateCounts` object.

`QueueAndStateCounts` contains:
- `queue` requested
- `count` of the total tasks in the queue
- `state_counts` map of state name to number of tasks in that state.