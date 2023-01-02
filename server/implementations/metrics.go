package implementations

import (
	"context"

	"github.com/TnLCommunity/corndogs/server/store"
	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
)

func (s *V1Alpha1Server) GetQueues(ctx context.Context, req *corndogsv1alpha1.EmptyRequest) (*corndogsv1alpha1.GetQueuesResponse, error) {
	response, err := store.AppStore.GetQueues()
	return response, err
}

func (s *V1Alpha1Server) GetQueueTaskCounts(ctx context.Context, req *corndogsv1alpha1.EmptyRequest) (*corndogsv1alpha1.GetQueueTaskCountsResponse, error) {
	response, err := store.AppStore.GetQueueTaskCounts()
	return response, err
}

func (s *V1Alpha1Server) GetStateCounts(ctx context.Context, req *corndogsv1alpha1.GetStateCountsRequest) (*corndogsv1alpha1.GetStateCountsResponse, error) {
	response, err := store.AppStore.GetStateCounts(req)
	return response, err
}

func (s *V1Alpha1Server) GetQueueAndStateCounts(ctx context.Context, req *corndogsv1alpha1.EmptyRequest) (*corndogsv1alpha1.GetQueueAndStateCountsResponse, error) {
	response, err := store.AppStore.GetQueueAndStateCounts()
	return response, err
}
