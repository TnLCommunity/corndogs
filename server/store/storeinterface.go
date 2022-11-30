package store

import (
	"time"

	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
)

var AppStore Store

type Store interface {
	Initialize() (deferredFunc func(), err error)
	CleanUpTimedOut(atTime time.Time) error

	SubmitTask(req *corndogsv1alpha1.SubmitTaskRequest) (*corndogsv1alpha1.SubmitTaskResponse, error)
	MustGetTaskStateByID(req *corndogsv1alpha1.GetTaskStateByIDRequest) *corndogsv1alpha1.GetTaskStateByIDResponse
	GetNextTask(req *corndogsv1alpha1.GetNextTaskRequest) (*corndogsv1alpha1.GetNextTaskResponse, error)
	UpdateTask(req *corndogsv1alpha1.UpdateTaskRequest) (*corndogsv1alpha1.UpdateTaskResponse, error)
	CompleteTask(req *corndogsv1alpha1.CompleteTaskRequest) (*corndogsv1alpha1.CompleteTaskResponse, error)
	CancelTask(req *corndogsv1alpha1.CancelTaskRequest) (*corndogsv1alpha1.CancelTaskResponse, error)
}

func SetStore(store Store) {
	AppStore = store
}
