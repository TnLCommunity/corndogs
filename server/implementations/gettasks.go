package implementations

import (
	"context"

	"github.com/TnLCommunity/corndogs/server/config"
	"github.com/TnLCommunity/corndogs/server/store"
	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
)

func (s *V1Alpha1Server) GetTaskStateByID(ctx context.Context, req *corndogsv1alpha1.GetTaskStateByIDRequest) (*corndogsv1alpha1.GetTaskStateByIDResponse, error) {
	response, err := store.AppStore.MustGetTaskStateByID(ctx, req)
	return response, err
}

func (s *V1Alpha1Server) GetNextTask(ctx context.Context, req *corndogsv1alpha1.GetNextTaskRequest) (*corndogsv1alpha1.GetNextTaskResponse, error) {
	if req.Queue == "" {
		req.Queue = config.DefaultQueue
	}
	if req.CurrentState == "" {
		req.CurrentState = config.DefaultStartingState
	}
	response, err := store.AppStore.GetNextTask(ctx, req)
	return response, err
}
