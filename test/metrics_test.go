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

	getQueuesResponse, err := corndogsClient.GetQueues(context.Background(), &corndogsv1alpha1.EmptyRequest{})
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

	getQueueTaskCountsResponse, err := corndogsClient.GetQueueTaskCounts(context.Background(), &corndogsv1alpha1.EmptyRequest{})
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.GreaterOrEqual(t, len(getQueueTaskCountsResponse.QueueCounts), 1, "expected at least one queue in queue_counts")
	require.GreaterOrEqual(t, getQueueTaskCountsResponse.TotalTaskCount, int64(1), "expected a total_task_count value")
	require.Equal(t, getQueueTaskCountsResponse.QueueCounts[testQueue], int64(1), "test queue should have a value of one in queue_counts")
}
