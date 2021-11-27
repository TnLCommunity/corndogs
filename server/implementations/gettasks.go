package implementations

import (
	"context"
	corndogsv1alpha1 "github.com/TnLCommunity/corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/TnLCommunity/corndogs/server/store"
)

func (s *V1Alpha1Server) MustGetTaskStateByID(ctx context.Context, req *corndogsv1alpha1.GetTaskStateByIDRequest) *corndogsv1alpha1.GetTaskStateByIDResponse {
	response := store.AppStore.MustGetTaskStateByID(req)
	return response
}

func (s *V1Alpha1Server) GetNextTask(ctx context.Context, req *corndogsv1alpha1.GetNextTaskRequest) (*corndogsv1alpha1.GetNextTaskResponse, error) {
	response, nil := store.AppStore.GetNextTask(req)
	return response, nil
}


