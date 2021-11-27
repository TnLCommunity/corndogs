package implementations

import (
	"context"
	corndogsv1alpha1 "github.com/TnLCommunity/corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/TnLCommunity/corndogs/server/store"
)

func (s *V1Alpha1Server) UpdateTask(ctx context.Context, req *corndogsv1alpha1.UpdateTaskRequest) (*corndogsv1alpha1.UpdateTaskResponse, error){
	response, nil := store.AppStore.UpdateTask(req)
	return response, nil
}


func (s *V1Alpha1Server) CompleteTask(ctx context.Context, req *corndogsv1alpha1.CompleteTaskRequest) (*corndogsv1alpha1.CompleteTaskResponse, error){
	response, nil := store.AppStore.CompleteTask(req)
	return response, nil
}


func (s *V1Alpha1Server) CancelTask(ctx context.Context, req *corndogsv1alpha1.CancelTaskRequest) (*corndogsv1alpha1.CancelTaskResponse, error){
	response, nil := store.AppStore.CancelTask(req)
	return response, nil
}

