package cmd

import (
	"context"
	"fmt"
	"time"

	corndogsv1alpha1 "github.com/TnLCommunity/protos-corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var timeoutCommand = NewTimeoutCommand()

func NewTimeoutCommand() *cobra.Command {
	var address string
	var port string
	timeoutCommand := &cobra.Command{
		Use:   "timeout",
		Short: "Send a CleanUpTimedOut request at the current time to a corndogs service",
		Long:  "Send a CleanUpTimedOut request at the current time to a corndogs service",
		Run: func(cmd *cobra.Command, args []string) {
			SendCleanUpTimedOut(address, port)
		},
	}

	timeoutCommand.Flags().StringVarP(&address, "address", "a", "127.0.0.1", "The address to connect to the corndogs service")
	timeoutCommand.Flags().StringVarP(&port, "port", "p", "5080", "The port to connect to the corndogs service")
	rootCmd.AddCommand(timeoutCommand)
	return timeoutCommand
}

func SendCleanUpTimedOut(address, port string) {

	// connect
	connectTo := fmt.Sprintf("%s:%s", address, port)
	fmt.Println("Connecting to:", connectTo)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, connectTo, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cancel()
	corndogsClient := corndogsv1alpha1.NewCorndogsServiceClient(conn)
	fmt.Println("Connected")

	nowUTC := time.Now().Add(time.Duration(7) * time.Second).UTC()
	fmt.Println("Sending at time:", nowUTC)
	timeToTimeout := nowUTC.UnixNano()
	cleanUpTimedOutRequest := &corndogsv1alpha1.CleanUpTimedOutRequest{
		AtTime: timeToTimeout,
	}
	cleanUpTimedOutResponse, err := corndogsClient.CleanUpTimedOut(context.Background(), cleanUpTimedOutRequest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Timed out: %d\n", cleanUpTimedOutResponse.TimedOut)
}
