package test

import (
	"context"
	"fmt"
	"testing"

	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/stretchr/testify/require"
)

func TestBasicPriority(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Payload:         testPayload,
	}
	// Priority 0
	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, int64(0), submitTaskResponse.Task.Priority)
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

	// Priority 1
	submitTaskRequest.Priority = 1
	submitTaskResponse, err = corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, int64(1), submitTaskResponse.Task.Priority)
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")
	onePriorityUuid := submitTaskResponse.Task.Uuid

	// Priority 0 (Sandwiched so no matter what "end" we take from we wont accidentally get the prioritized task)
	submitTaskRequest.Priority = 0
	submitTaskResponse, err = corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")

	getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:        testQueue,
		CurrentState: "testSubmitted",
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, onePriorityUuid, getNextTaskResponse.Task.Uuid, "didnt get top priority task")
	require.Equal(t, int64(1), getNextTaskResponse.Task.Priority, "expected priority to be one")

	getNextTaskResponse, err = corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, int64(0), getNextTaskResponse.Task.Priority, "expected priority to be zero")
}

func TestBasicDeprioritize(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Priority:        -1,
		Payload:         testPayload,
	}
	// Priority -1
	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, int64(0), submitTaskResponse.Task.Priority)
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

	// Priority 0
	submitTaskRequest.Priority = 0
	submitTaskResponse, err = corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, int64(-1), submitTaskResponse.Task.Priority)
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")
	defaultPriorityUuid := submitTaskResponse.Task.Uuid

	// Priority -1 (Sandwiched so no matter what "end" we take from we wont accidentally get the non-deprioritized task)
	submitTaskRequest.Priority = -1
	submitTaskResponse, err = corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")

	getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:        testQueue,
		CurrentState: "testSubmitted",
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, defaultPriorityUuid, getNextTaskResponse.Task.Uuid, "didnt get top priority task")
	require.Equal(t, int64(0), getNextTaskResponse.Task.Priority, "expected priority to be zero")

	getNextTaskResponse, err = corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, int64(-1), getNextTaskResponse.Task.Priority, "expected priority to be negative one")
}
