package test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	"google.golang.org/grpc"
)

var client = GetCorndogsClient()
var testID = gofakeit.Breakfast() + gofakeit.Dessert()

// assume an empty db, but we are using conventions to try to do this against a live server if needed
func init() {
	// Get the next task in the test_via_core_corndogs_repo queue, which should be none
	getNextRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue:        "testQueue" + testID,
		CurrentState: "submitted",
	}
	nextTaskResponse, err := client.GetNextTask(context.Background(), getNextRequest)
	if err != nil {
		panic(err)
	}
	if nextTaskResponse.Task != nil {
		panic(nextTaskResponse.Task.Uuid)
	}
}

func TestBasicFlow(t *testing.T) {
	corndogsClient := GetCorndogsClient()
	rand.Seed(time.Now().UnixNano())
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)

	// Differentiate values so they can be run multiple at a time on a live environment easily
	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           "testQueue" + testID,
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
		Queue:        "testQueue" + testID,
		CurrentState: "testSubmitted",
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, getNextTaskRequest.CurrentState+workingTaskSuffix, getNextTaskResponse.Task.CurrentState, "Task CurrentState is not the auto target state from before retrieval")
	require.Equal(t, getNextTaskRequest.CurrentState, getNextTaskResponse.Task.AutoTargetState, "Task AutoTargetState is not swapped with current state before retrieval")

	updateTaskRequest := &corndogsv1alpha1.UpdateTaskRequest{
		Uuid:            getNextTaskResponse.Task.Uuid,
		Queue:           "testQueue" + testID,
		CurrentState:    "testSubmitted" + workingTaskSuffix,
		AutoTargetState: "testSubmitted-completing",
		NewState:        "testSubmitted-updated", // This is not used in automated flow, we're adding a new flow
	}
	updateTaskResponse, err := corndogsClient.UpdateTask(context.Background(), updateTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, updateTaskResponse.Task, "Task in response was nil")
	require.Equal(t, updateTaskRequest.Queue, updateTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, updateTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, updateTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, updateTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, updateTaskRequest.NewState, updateTaskResponse.Task.CurrentState, "Task State was not updated")
	require.Equal(t, updateTaskRequest.AutoTargetState, updateTaskResponse.Task.AutoTargetState, "Task State was not updated")

	// Now get the updated task
	getNextTaskRequest = &corndogsv1alpha1.GetNextTaskRequest{
		Queue:        "testQueue" + testID,
		CurrentState: "testSubmitted-updated",
	}
	getNextTaskResponse, err = corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, "testSubmitted-completing", getNextTaskResponse.Task.CurrentState, "Task CurrentState is not the auto target state from before retrieval")
	require.Equal(t, getNextTaskRequest.CurrentState, getNextTaskResponse.Task.AutoTargetState, "Task CurrentState does not have correct suffix added")

	completeTaskRequest := &corndogsv1alpha1.CompleteTaskRequest{
		Uuid:         getNextTaskResponse.Task.Uuid,
		Queue:        "testQueue" + testID,
		CurrentState: "testSubmitted-updated" + workingTaskSuffix,
	}
	completeTaskResponse, err := corndogsClient.CompleteTask(context.Background(), completeTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, completeTaskResponse.Task, "Task in response was nil")
	require.Equal(t, completeTaskRequest.Queue, completeTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, completeTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, completeTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, completeTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, "completed", completeTaskResponse.Task.CurrentState, "Task CurrentState does not have correct suffix added")
	require.Equal(t, []byte(nil), completeTaskResponse.Task.Payload, "Task Payload is not nil after completion")
}

func TestGetNextTaskOverrideState(t *testing.T) {
	corndogsClient := GetCorndogsClient()
	rand.Seed(time.Now().UnixNano())
	workingTaskSuffix := "-working"
	testPayload := []byte("testPayload" + testID)

	submitTaskRequest := &corndogsv1alpha1.SubmitTaskRequest{
		Queue:           "testQueue" + testID,
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
		Queue:                   "testQueue" + testID,
		CurrentState:            "testSubmitted",
		OverrideCurrentState:    "testOverriddenCurrentState",
		OverrideAutoTargetState: "testOverriddenAutoTargetState",
	}
	getNextTaskResponse, err := corndogsClient.GetNextTask(context.Background(), getNextTaskRequest)
	require.Nil(t, err, fmt.Sprintf("error should be nil. error: \n%v", err))
	require.NotNil(t, getNextTaskResponse.Task, "Task in response was nil")
	require.Equal(t, getNextTaskRequest.Queue, getNextTaskResponse.Task.Queue, "Queue name is not equal")
	require.NotEmpty(t, getNextTaskResponse.Task.SubmitTime, "submit_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.UpdateTime, "update_time should not be empty")
	require.NotEmpty(t, getNextTaskResponse.Task.Uuid, "uuid should not be empty")
	require.Equal(t, getNextTaskRequest.OverrideCurrentState, getNextTaskResponse.Task.CurrentState, "Task CurrentState is not the overridden")
	require.Equal(t, getNextTaskRequest.OverrideAutoTargetState, getNextTaskResponse.Task.AutoTargetState, "Task AutoTargetState is not overridden")
}

func GetCorndogsClient() corndogsv1alpha1.CorndogsServiceClient {
	// connect
	connectTo := "127.0.0.1:5080"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, connectTo, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cancel()
	return corndogsv1alpha1.NewCorndogsServiceClient(conn)
}
