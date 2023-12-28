package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// submitTaskCmd represents the submitTask command
var submitTaskCmd = NewSubmitTaskCmd()

func NewSubmitTaskCmd() *cobra.Command {
	var address, port string
	var queue, currentState, autoTargetState, payload string
	var timeout, priority int64
	cmd := &cobra.Command{
		Use:   "submit-task",
		Short: "creates a corndogs task",
		Run: func(cmd *cobra.Command, args []string) {
			doSubmitTask(address, port, &corndogsv1alpha1.SubmitTaskRequest{
				Queue:           queue,
				CurrentState:    currentState,
				AutoTargetState: autoTargetState,
				Timeout:         timeout,
				Payload:         []byte(payload),
				Priority:        priority,
			})
		},
	}
	cmd.Flags().StringVarP(&address, "address", "a", "127.0.0.1", "The address to connect to the corndogs service")
	cmd.Flags().StringVarP(&port, "port", "p", "5080", "The port to connect to the corndogs service")
	cmd.Flags().StringVarP(&queue, "queue", "q", "", "The queue to submit the task to")
	cmd.Flags().StringVarP(&currentState, "current-state", "c", "", "The current state of the task")
	cmd.Flags().StringVarP(&autoTargetState, "auto-target-state", "t", "", "The target state of the task")
	cmd.Flags().Int64VarP(&timeout, "timeout", "o", 0, "The timeout of the task")
	cmd.Flags().StringVarP(&payload, "payload", "l", "", "The payload of the task")
	cmd.Flags().Int64VarP(&priority, "priority", "r", 0, "The priority of the task")
	rootCmd.AddCommand(cmd)
	return cmd
}

func doSubmitTask(address, port string, req *corndogsv1alpha1.SubmitTaskRequest) {
	conn, client, err := getClient(address, port, 5*time.Second)
	if err != nil {
		log.Err(err).Msg("failed to get client")
		os.Exit(1)
	}
	defer conn.Close()
	resp, err := client.SubmitTask(context.Background(), req)
	if err != nil {
		log.Err(err).Msg("failed to submit task")
	}
	log.Info().Msgf("response: %v", resp)
}

func getClient(address, port string, timeout time.Duration) (*grpc.ClientConn, corndogsv1alpha1.CorndogsServiceClient, error) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	target := fmt.Sprintf("%s:%s", address, port)
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	for err != nil {
		return nil, nil, err
	}
	corndogsClient := corndogsv1alpha1.NewCorndogsServiceClient(conn)
	return conn, corndogsClient, nil
}
