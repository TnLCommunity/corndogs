package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/stretchr/testify/require"
)

func TestBasicTimeout(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID
	var timeout int64 = 5
	timeoutDuration := time.Duration(timeout) * time.Second

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Timeout:         timeout,
		Payload:         testPayload,
	}
	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, timeout, submitTaskResponse.Task.Timeout, "Timeout is not equal")
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

	getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:        testQueue,
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

	timeWhenTimedout := time.Now().UTC().Add(timeoutDuration).UnixNano()
	cleanUpTimedOutRequest := &corndogsv1alpha1.CleanUpTimedOutRequest{
		AtTime: timeWhenTimedout,
		Queue:  testQueue,
	}
	cleanUpTimedOutResponse, err := corndogsClient.CleanUpTimedOut(context.Background(), cleanUpTimedOutRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.Equal(t, int64(1), cleanUpTimedOutResponse.TimedOut)

	// If everything is working things like state should be the same except Timeout should now be 0
	getNextTaskResponse, err = corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, int64(0), getNextTaskResponse.Task.Timeout, "Timeout should now be 0")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, getNextTaskRequest.CurrentState+workingTaskSuffix, getNextTaskResponse.Task.CurrentState, "Task CurrentState is not the auto target state from before retrieval")
	require.Equal(t, getNextTaskRequest.CurrentState, getNextTaskResponse.Task.AutoTargetState, "Task AutoTargetState is not swapped with current state before retrieval")
}

func TestNoTimeout(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID
	var timeout int64 = -1

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Timeout:         timeout,
		Payload:         testPayload,
	}

	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, int64(0), submitTaskResponse.Task.Timeout, "Timeout should be 0")
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

	getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:        testQueue,
		CurrentState: "testSubmitted",
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.Equal(t, int64(0), getNextTaskResponse.Task.Timeout, "Timeout should be zero meaning no timeout")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, getNextTaskRequest.CurrentState+workingTaskSuffix, getNextTaskResponse.Task.CurrentState, "Task CurrentState is not the auto target state from before retrieval")
	require.Equal(t, getNextTaskRequest.CurrentState, getNextTaskResponse.Task.AutoTargetState, "Task AutoTargetState is not swapped with current state before retrieval")

	timeoutDuration := time.Duration(5) * time.Second
	timeWhenTimedout := time.Now().UTC().Add(timeoutDuration).UnixNano()
	cleanUpTimedOutRequest := &corndogsv1alpha1.CleanUpTimedOutRequest{
		AtTime: timeWhenTimedout,
		Queue:  testQueue,
	}
	cleanUpTimedOutResponse, err := corndogsClient.CleanUpTimedOut(context.Background(), cleanUpTimedOutRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.Equal(t, int64(0), cleanUpTimedOutResponse.TimedOut)

	getNextTaskResponse, err = corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.Nil(t, getNextTaskResponse.Task, "Task was not nil.")
}

func TestGetNextTaskOverrideTimeout(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID
	var timeout int64 = 5

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

	getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		OverrideTimeout: timeout,
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, getNextTaskRequest.OverrideTimeout, getNextTaskResponse.Task.Timeout, "Task Timeout is not the overridden")

	timeoutDuration := time.Duration(timeout) * time.Second
	timeWhenTimedout := time.Now().UTC().Add(timeoutDuration).UnixNano()
	cleanUpTimedOutRequest := &corndogsv1alpha1.CleanUpTimedOutRequest{
		AtTime: timeWhenTimedout,
		Queue:  testQueue,
	}
	cleanUpTimedOutResponse, err := corndogsClient.CleanUpTimedOut(context.Background(), cleanUpTimedOutRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.Equal(t, int64(1), cleanUpTimedOutResponse.TimedOut)
}

func TestGetNextTaskOverrideNoTimeout(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID
	var timeout int64 = 5

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Timeout:         timeout,
		Payload:         testPayload,
	}
	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

	getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		OverrideTimeout: -1, // No timeout
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, int64(0), getNextTaskResponse.Task.Timeout, "Task Timeout is not the overridden with 0")

	timeoutDuration := time.Duration(timeout) * time.Second
	timeWhenTimedout := time.Now().UTC().Add(timeoutDuration).UnixNano()
	cleanUpTimedOutRequest := &corndogsv1alpha1.CleanUpTimedOutRequest{
		AtTime: timeWhenTimedout,
		Queue:  testQueue,
	}
	cleanUpTimedOutResponse, err := corndogsClient.CleanUpTimedOut(context.Background(), cleanUpTimedOutRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.Equal(t, int64(0), cleanUpTimedOutResponse.TimedOut)
}

func TestGetNextTaskOverrideTimeoutNotSet(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID
	var timeout int64 = 5

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           testQueue,
		CurrentState:    "testSubmitted",
		AutoTargetState: "testSubmitted" + workingTaskSuffix,
		Timeout:         timeout,
		Payload:         testPayload,
	}
	submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
	require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

	getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:        testQueue,
		CurrentState: "testSubmitted",
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.NotEqual(t, int64(0), getNextTaskResponse.Task.Timeout, "Task Timeout was wrongly overriden")
	require.Equal(t, timeout, getNextTaskResponse.Task.Timeout, "Task Timeout is not equal")

	timeoutDuration := time.Duration(timeout) * time.Second
	timeWhenTimedout := time.Now().UTC().Add(timeoutDuration).UnixNano()
	cleanUpTimedOutRequest := &corndogsv1alpha1.CleanUpTimedOutRequest{
		AtTime: timeWhenTimedout,
		Queue:  testQueue,
	}
	cleanUpTimedOutResponse, err := corndogsClient.CleanUpTimedOut(context.Background(), cleanUpTimedOutRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.Equal(t, int64(1), cleanUpTimedOutResponse.TimedOut)
}

func TestTimeoutSpecificQueue(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID
	var timeout int64 = 5
	queueSuffix := []string{"first", "second"}

	// Create and get a task in each queue
	for _, suffix := range queueSuffix {
		submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
			Queue:           testQueue + suffix,
			CurrentState:    "testSubmitted",
			AutoTargetState: "testSubmitted" + workingTaskSuffix,
			Timeout:         timeout,
			Payload:         testPayload,
		}

		submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
		require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
		require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
		require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
		require.Equal(t, timeout, submitTaskResponse.Task.Timeout, "Timeout should be 0")
		require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
		require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
		require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

		getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
			Queue:        testQueue + suffix,
			CurrentState: "testSubmitted",
		}
		getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
		require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
		require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
		require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
		require.Equal(t, timeout, getNextTaskResponse.Task.Timeout, "Timeout should be zero meaning no timeout")
		require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
		require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
		require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
		require.Equal(t, getNextTaskRequest.CurrentState+workingTaskSuffix, getNextTaskResponse.Task.CurrentState, "Task CurrentState is not the auto target state from before retrieval")
		require.Equal(t, getNextTaskRequest.CurrentState, getNextTaskResponse.Task.AutoTargetState, "Task AutoTargetState is not swapped with current state before retrieval")
	}

	// Timeout each queue individually
	for _, suffix := range queueSuffix {
		timeoutDuration := time.Duration(timeout) * time.Second
		timeWhenTimedout := time.Now().UTC().Add(timeoutDuration).UnixNano()
		cleanUpTimedOutRequest := &corndogsv1alpha1.CleanUpTimedOutRequest{
			AtTime: timeWhenTimedout,
			Queue:  testQueue + suffix,
		}
		cleanUpTimedOutResponse, err := corndogsClient.CleanUpTimedOut(context.Background(), cleanUpTimedOutRequest)
		require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
		require.Equal(t, int64(1), cleanUpTimedOutResponse.TimedOut, "didnt time out specific queue")
	}
}

func TestTimeoutNoQueue(t *testing.T) {
	testID := GetTestID()
	corndogsClient := GetCorndogsClient()
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)
	testQueue := "testQueue" + testID
	var timeout int64 = 5
	queueSuffix := []string{"first", "second"}

	// Create and get a task in each queue
	for _, suffix := range queueSuffix {
		submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
			Queue:           testQueue + suffix,
			CurrentState:    "testSubmitted",
			AutoTargetState: "testSubmitted" + workingTaskSuffix,
			Timeout:         timeout,
			Payload:         testPayload,
		}

		submitTaskResponse, err := corndogsClient.SubmitTask(context.Background(), submitTaskRequest)
		require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
		require.NotNil(t, submitTaskResponse.Task, "Task in response was nil")
		require.Equal(t, submitTaskRequest.Queue, submitTaskResponse.Task.Queue, "Queue name is not equal")
		require.Equal(t, timeout, submitTaskResponse.Task.Timeout, "Timeout should be 0")
		require.NotEmpty(t, submitTaskResponse.Task.SubmitTime, "submit_time should not be empty")
		require.NotEmpty(t, submitTaskResponse.Task.UpdateTime, "update_time should not be empty")
		require.NotEmpty(t, submitTaskResponse.Task.Uuid, "uuid should not be empty")

		getNextTaskRequest := &corndogsv1alpha1.GetNextTaskRequest{
			Queue:        testQueue + suffix,
			CurrentState: "testSubmitted",
		}
		getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
		require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
		require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
		require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
		require.Equal(t, timeout, getNextTaskResponse.Task.Timeout, "Timeout should be zero meaning no timeout")
		require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
		require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
		require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
		require.Equal(t, getNextTaskRequest.CurrentState+workingTaskSuffix, getNextTaskResponse.Task.CurrentState, "Task CurrentState is not the auto target state from before retrieval")
		require.Equal(t, getNextTaskRequest.CurrentState, getNextTaskResponse.Task.AutoTargetState, "Task AutoTargetState is not swapped with current state before retrieval")
	}

	// Timeout both queues
	timeoutDuration := time.Duration(timeout) * time.Second
	timeWhenTimedout := time.Now().UTC().Add(timeoutDuration).UnixNano()
	cleanUpTimedOutRequest := &corndogsv1alpha1.CleanUpTimedOutRequest{
		AtTime: timeWhenTimedout,
	}
	cleanUpTimedOutResponse, err := corndogsClient.CleanUpTimedOut(context.Background(), cleanUpTimedOutRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.GreaterOrEqual(t, cleanUpTimedOutResponse.TimedOut, int64(2), "didnt time out multiple queues")
}
