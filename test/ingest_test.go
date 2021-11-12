package test

import (
	"context"
	corndogsv1alpha1 "github.com/TnLCommunity/corndogs/gen/proto/go/corndogs/v1alpha1"
	"testing"
	"time"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	"google.golang.org/grpc"
)

var client = GetIngestClient()

// assume an empty db, but we are using conventions to try to do this against a live server if needed
func init() {
	// Get the next task in the test_via_core_corndogs_repo queue, which should be none
	getNextRequest := &corndogsv1alpha1.GetNextTaskRequest{
		Queue: "test_via_core_corndogs_repo",
		CurrentState: "submitted",
	}
	nextTaskResponse, err := client.GetNextTask(context.Background(), getNextRequest)
	if err != nil {
		panic(err)
	}
	if nextTaskResponse.Task != nil {
		panic(nextTaskResponse.Task)
	}
}

func TestCrud(t *testing.T) {
	return
}

func GetIngestClient() corndogsv1alpha1.CorndogsServiceClient {
	// connect
	connectTo := "127.0.0.1:8080"
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, connectTo, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return corndogsv1alpha1.NewCorndogsServiceClient(conn)
}
