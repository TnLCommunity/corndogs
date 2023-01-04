package store

import (
	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
)

var AppStore Store

type Store interface {
	Initialize() (deferredFunc func(), err error)

	SubmitTask(req *corndogsv1alpha1.SubmitTaskRequest) (*corndogsv1alpha1.SubmitTaskResponse, error)
	MustGetTaskStateByID(req *corndogsv1alpha1.GetTaskStateByIDRequest) (*corndogsv1alpha1.GetTaskStateByIDResponse, error)
	GetNextTask(req *corndogsv1alpha1.GetNextTaskRequest) (*corndogsv1alpha1.GetNextTaskResponse, error)
	UpdateTask(req *corndogsv1alpha1.UpdateTaskRequest) (*corndogsv1alpha1.UpdateTaskResponse, error)
	CompleteTask(req *corndogsv1alpha1.CompleteTaskRequest) (*corndogsv1alpha1.CompleteTaskResponse, error)
	CancelTask(req *corndogsv1alpha1.CancelTaskRequest) (*corndogsv1alpha1.CancelTaskResponse, error)
	CleanUpTimedOut(req *corndogsv1alpha1.CleanUpTimedOutRequest) (*corndogsv1alpha1.CleanUpTimedOutResponse, error)
	// Metrics
	GetQueues(req *corndogsv1alpha1.GetQueuesRequest) (*corndogsv1alpha1.GetQueuesResponse, error)
	GetQueueTaskCounts(req *corndogsv1alpha1.GetQueueTaskCountsRequest) (*corndogsv1alpha1.GetQueueTaskCountsResponse, error)
	GetTaskStateCounts(req *corndogsv1alpha1.GetTaskStateCountsRequest) (*corndogsv1alpha1.GetTaskStateCountsResponse, error)
	GetQueueAndStateCounts(req *corndogsv1alpha1.GetQueueAndStateCountsRequest) (*corndogsv1alpha1.GetQueueAndStateCountsResponse, error)
}

func SetStore(store Store) {
	AppStore = store
}
