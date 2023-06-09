package implementations

import (
	"context"

	"github.com/TnLCommunity/corndogs/server/config"
	"github.com/TnLCommunity/corndogs/server/metrics"
	"github.com/TnLCommunity/corndogs/server/store"
	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
)

func (s *V1Alpha1Server) UpdateTask(ctx context.Context, req *corndogsv1alpha1.UpdateTaskRequest) (*corndogsv1alpha1.UpdateTaskResponse, error) {
	if req.CurrentState == "" {
		req.CurrentState = config.DefaultStartingState
	}
	if req.NewState == "" {
		req.NewState = "updated"
	}
	if req.AutoTargetState == "" {
		req.AutoTargetState = req.NewState + config.DefaultWorkingSuffix
	}
	response, err := store.AppStore.UpdateTask(ctx, req)
	return response, err
}

func (s *V1Alpha1Server) CompleteTask(ctx context.Context, req *corndogsv1alpha1.CompleteTaskRequest) (*corndogsv1alpha1.CompleteTaskResponse, error) {
	response, err := store.AppStore.CompleteTask(ctx, req)
	if config.PrometheusEnabled && err == nil {
		metrics.CompletedTasksTotal.Inc()
	}
	return response, err
}

func (s *V1Alpha1Server) CancelTask(ctx context.Context, req *corndogsv1alpha1.CancelTaskRequest) (*corndogsv1alpha1.CancelTaskResponse, error) {
	response, err := store.AppStore.CancelTask(ctx, req)
	if config.PrometheusEnabled && err == nil {
		metrics.CanceledTasksTotal.Inc()
	}
	return response, err
}

func (s *V1Alpha1Server) CleanUpTimedOut(ctx context.Context, req *corndogsv1alpha1.CleanUpTimedOutRequest) (*corndogsv1alpha1.CleanUpTimedOutResponse, error) {
	response, err := store.AppStore.CleanUpTimedOut(ctx, req)
	if config.PrometheusEnabled && err == nil {
		metrics.TimedOutTasksTotal.Inc()
	}
	return response, err
}
