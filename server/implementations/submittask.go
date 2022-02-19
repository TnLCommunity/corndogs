package implementations

import (
	"context"

	"github.com/TnLCommunity/corndogs/server/store"
	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
)

func (s *V1Alpha1Server) SubmitTask(ctx context.Context, req *corndogsv1alpha1.SubmitTaskRequest) (*corndogsv1alpha1.SubmitTaskResponse, error) {
	response, nil := store.AppStore.SubmitTask(req)
	return response, nil
}
