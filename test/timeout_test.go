package test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/stretchr/testify/require"
)

func TestBasicTimeout(t *testing.T) {
	corndogsClient := GetCorndogsClient()
	rand.Seed(time.Now().UnixNano())
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	var timeout int64 = 5
	timeoutDuration := time.Duration(timeout) * time.Second

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           "testQueue" + testID,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Timeout:         timeout,
		Payload:         testPayload,
	}
	/*
		TEST 1
		send with timeout.
		cleanup timed out
		check timed out.
		check CurrentState returns back to previous
	*/

	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, timeout, submitTaskResponse.Task.Timeout, "Timeout is not equal")
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

	getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:        "testQueue" + testID,
		CurrentState: "testSubmitted",
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, timeout, getNextTaskResponse.Task.Timeout, "Timeout is not equal")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, getNextTaskRequest.CurrentState+workingTaskSuffix, getNextTaskResponse.Task.CurrentState, "Task CurrentState is not the auto target state from before retrieval")
	require.Equal(t, getNextTaskRequest.CurrentState, getNextTaskResponse.Task.AutoTargetState, "Task AutoTargetState is not swapped with current state before retrieval")

	timeWhenTimedout := time.Now().Add(timeoutDuration)
	cleanUpTimedOutRequest := &corndogsv1alpha1.CleanUpTimedOutRequest{
		AtTime: timeWhenTimedout.UnixNano(),
	}
	cleanUpTimedOutResponse, err := corndogsClient.CleanUpTimedOut(context.Background(), cleanUpTimedOutRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.Equal(t, cleanUpTimedOutResponse.TimedOut, 1)

	// If everything is working this should be the same, meaning things like the state are returned to their previous values.
	getNextTaskResponse, err = corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, timeout, getNextTaskResponse.Task.Timeout, "Timeout is not equal")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, getNextTaskRequest.CurrentState+workingTaskSuffix, getNextTaskResponse.Task.CurrentState, "Task CurrentState is not the auto target state from before retrieval")
	require.Equal(t, getNextTaskRequest.CurrentState, getNextTaskResponse.Task.AutoTargetState, "Task AutoTargetState is not swapped with current state before retrieval")
}

func TestTimeoutDefault(t *testing.T) {
	// TEST 2
	// send with timeout 0
	// check timeout is default
}

func TestNoTimeout(t *testing.T) {
	// TEST 3
	// send with timeout -1
	// check timeout is 0 in DB?
	// cleanup timed out
	// Check it did not timeout
	// Check CurrentState is the same as the working state
}
