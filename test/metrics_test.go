package test

import (
	"context"
	"fmt"
	"testing"

	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/stretchr/testify/require"
)

func TestGetQueues(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Timeout:         -1, // No timeout
		Payload:         testPayload,
	}
	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

	getQueuesResponse, err := corndogsClient.GetQueues(context.Background(), &corndogsv1alpha1.GetQueuesRequest{})
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.GreaterOrEqual(t, len(getQueuesResponse.Queues), 1, "expected at least one queue in list of queues")
	require.GreaterOrEqual(t, getQueuesResponse.TotalTaskCount, int64(1), "expected a total_task_count value")
	require.Contains(t, getQueuesResponse.Queues, testQueue, "test queue should be in list of queues")
}

func TestGetQueueTaskCounts(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Timeout:         -1, // No timeout
		Payload:         testPayload,
	}
	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

	getQueueTaskCountsResponse, err := corndogsClient.GetQueueTaskCounts(context.Background(), &corndogsv1alpha1.GetQueueTaskCountsRequest{})
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.GreaterOrEqual(t, len(getQueueTaskCountsResponse.QueueCounts), 1, "expected at least one queue in queue_counts")
	require.GreaterOrEqual(t, getQueueTaskCountsResponse.TotalTaskCount, int64(1), "expected a total_task_count value")
	require.Equal(t, int64(1), getQueueTaskCountsResponse.QueueCounts[testQueue], "test queue should have a value of one in queue_counts")
}

func TestGetTaskStateCounts(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Timeout:         -1, // No timeout
		Payload:         testPayload,
	}
	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")
	// Another one
	submitTaskResponse, err = corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")

	// Move one to the AutoTargetState
	getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:        testQueue,
		CurrentState: "testSubmitted",
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")

	getTaskStateCountsRequest := &corndogsv1alpha1.GetTaskStateCountsRequest{
		Queue: testQueue,
	}

	getTaskStateCountsResponse, err := corndogsClient.GetTaskStateCounts(context.Background(), getTaskStateCountsRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.Equal(t, testQueue, getTaskStateCountsResponse.Queue, "Queue name is not equal")
	require.Equal(t, int64(2), getTaskStateCountsResponse.Count, "expected two tasks in queue")
	require.Equal(t, int64(1), getTaskStateCountsResponse.StateCounts[submitTaskRequest.CurrentState], "expected one task for initial state")
	require.Equal(t, int64(1), getTaskStateCountsResponse.StateCounts[submitTaskRequest.AutoTargetState], "expected one task for auto target state")
}

func TestGetQueueAndStateCounts(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID
	queueSuffix := []string{"first", "second"}
	currentState := "testSubmitted"
	autoTargetState := currentState + workingTaskSuffix

	// Create and get a task in each queue
	for _, suffix := range queueSuffix {
		submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
			Queue:           testQueue + suffix,
			CurrentState:    currentState,
			AutoTargetState: autoTargetState,
			Timeout:         -1, // No timeout
			Payload:         testPayload,
		}
		submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
		require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
		require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
		require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
		require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
		require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
		require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")
		// Another one
		submitTaskResponse, err = corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
		require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
		require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")

		// Move one to the AutoTargetState
		getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
			Queue:        testQueue + suffix,
			CurrentState: "testSubmitted",
		}
		getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
		require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
		require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
		require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	}

	getQueueAndStateCountsResponse, err := corndogsClient.GetQueueAndStateCounts(context.Background(), &corndogsv1alpha1.GetQueueAndStateCountsRequest{})
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	queueAndStateCounts := getQueueAndStateCountsResponse.QueueAndStateCounts

	for _, suffix := range queueSuffix {
		require.Contains(t, queueAndStateCounts, testQueue+suffix, "expected queue to exist in map")
		require.Equal(t, testQueue+suffix, queueAndStateCounts[testQueue+suffix].Queue, "Queue name is not equal")
		require.Equal(t, int64(2), queueAndStateCounts[testQueue+suffix].Count, "expected two tasks in queue")
		require.Equal(t, int64(1), queueAndStateCounts[testQueue+suffix].StateCounts[currentState], "expected one task for initial state")
		require.Equal(t, int64(1), queueAndStateCounts[testQueue+suffix].StateCounts[autoTargetState], "expected one task for auto target state")
	}
}
