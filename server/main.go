package main

import (
	"fmt"
	corndogsv1alpha1 "github.com/TnLCommunity/corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/TnLCommunity/corndogs/server/implementations"
	"github.com/TnLCommunity/corndogs/server/logging"
	"github.com/TnLCommunity/corndogs/server/store"
	"github.com/TnLCommunity/corndogs/server/store/postgresstore"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	"google.golang.org/grpc"
)

func main() {
	// get logging setup
	logging.InitLogging()
	// use the postgres store
	store.SetStore(postgresstore.PostgresStore{})
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// initialize store
	deferredFunc, err := store.AppStore.Initialize()
	if err != nil {
		panic(err)
	}
	// if the store has a deferred call, defer it
	if deferredFunc != nil {
		defer deferredFunc()
	}
	// create listener
	listenOn := "0.0.0.0:5080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	// recovery handler - useless for now, fill in with sentry or something later
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			fmt.Printf("panicked: %v", p)
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}
	// create grpc server
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(opts...),
			),
		),
	)

	// register ingest service
	corndogsv1alpha1.RegisterCorndogsServiceServer(server, &implementations.V1Alpha1Server{})

	// register health service (used in k8s health checks)
	healthService := implementations.NewHealthChecker()
	grpc_health_v1.RegisterHealthServer(server, healthService)

	// serve
	log.Println("Listening on", listenOn)
	err = server.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
