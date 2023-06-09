package implementations

import (
	"context"

	"github.com/TnLCommunity/corndogs/server/config"
	"github.com/TnLCommunity/corndogs/server/metrics"
	"github.com/TnLCommunity/corndogs/server/store"
	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
)

func (s *V1Alpha1Server) SubmitTask(ctx context.Context, req *corndogsv1alpha1.SubmitTaskRequest) (*corndogsv1alpha1.SubmitTaskResponse, error) {
	if req.Queue == "" {
		req.Queue = config.DefaultQueue
	}
	if req.CurrentState == "" {
		req.CurrentState = config.DefaultStartingState
	}
	if req.AutoTargetState == "" {
		req.AutoTargetState = req.CurrentState + config.DefaultWorkingSuffix
	}

	// Since protobuf default int values are 0, if they wanted 0, they have to send a negative value
	// which is otherwise invalid as a timeout
	if req.Timeout == 0 {
		req.Timeout = config.DefaultTimeout
	}
	if req.Timeout < 0 {
		req.Timeout = 0
	}
	response, err := store.AppStore.SubmitTask(ctx, req)
	if config.PrometheusEnabled && err == nil {
		metrics.TasksTotal.Inc()
	}
	return response, err
}
