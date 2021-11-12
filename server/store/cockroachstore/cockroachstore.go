package cockroachstore

import (
	"context"
	"embed"
	"errors"
	"fmt"
	corndogsv1alpha1 "github.com/TnLCommunity/corndogs/gen/proto/go/corndogs/v1alpha1"
	"github.com/TnLCommunity/corndogs/server/config"
	"github.com/TnLCommunity/corndogs/server/store/cockroachstore/models"
	"github.com/TnLCommunity/corndogs/server/utils"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// global db
var DB *gorm.DB

var DatabaseName = config.GetEnvOrDefault("DATABASE_NAME", "defaultdb")
var DatabaseHost = config.GetEnvOrDefault("DATABASE_HOST", "localhost")
var DatabaseUser = config.GetEnvOrDefault("DATABASE_USER", "root")
var DatabasePassword = config.GetEnvOrDefault("DATABASE_PASSWORD", "root")
var DatabasePort = config.GetEnvOrDefault("DATABASE_PORT", "26257")
var DatabaseSSLMode = config.GetEnvOrDefault("DATABASE_SSL_MODE", "disable")
var MaxIdleConns = config.GetEnvAsIntOrDefault("DATABASE_MAX_IDLE_CONNS", "1")
var MaxOpenConns = config.GetEnvAsIntOrDefault("DATABASE_MAX_OPEN_CONNS", "10")
var ConnMaxLifetime = time.Duration(config.GetEnvAsIntOrDefault("DATABASE_CONN_MAX_LIFETIME_SECONDS", "3600")) * time.Second

// sql files embedded at compile time, used by goose
//go:embed migrations/*.sql
var embedMigrations embed.FS

type CockroachStore struct{}

func (s CockroachStore) Initialize() (func(), error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", DatabaseHost, DatabaseUser, DatabasePassword, DatabaseName, DatabasePort, DatabaseSSLMode)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDb, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(MaxIdleConns)
	sqlDb.SetMaxOpenConns(MaxOpenConns)
	sqlDb.SetConnMaxLifetime(ConnMaxLifetime)
	// configure database connection settings
	fmt.Printf("Connected to %q", DB.Name())
	goose.SetBaseFS(embedMigrations)

	err = goose.Up(sqlDb, "migrations")
	if err != nil {
		return nil, err
	}
	return func() { sqlDb.Close() }, nil
}

func (s CockroachStore) SubmitTask(req *corndogsv1alpha1.SubmitTaskRequest) (*corndogsv1alpha1.SubmitTaskResponse, error) {
	taskProto := &corndogsv1alpha1.Task{}
	err := crdbgorm.ExecuteTx(context.Background(), DB, nil,
		func(tx *gorm.DB) error {
			// create model, only use fields relevant to this request, let DB do the rest
			model := models.Task{
				Queue: req.Queue,
				CurrentState: req.CurrentState,
				AutoTargetState: req.AutoTargetState,
				Timeout: req.Timeout,
				Payload: req.Payload,
			}
			result := DB.Create(&model)
			if result.Error != nil {
				return result.Error
			}
			// marshall result to response
			return utils.StructToProto(model, taskProto)
		},
	)

	return &corndogsv1alpha1.SubmitTaskResponse{Task: taskProto}, err
}

func (s CockroachStore) MustGetTaskStateByID(req *corndogsv1alpha1.GetTaskStateByIDRequest) *corndogsv1alpha1.GetTaskStateResponse{
	taskProto := &corndogsv1alpha1.Task{}
	err := crdbgorm.ExecuteTx(context.Background(), DB, nil,
		func(tx *gorm.DB) error {
			model := models.Task{UUID: req.Uuid}
			result := DB.First(&model)
			if result.Error != nil {
				if errors.Is(result.Error, gorm.ErrRecordNotFound) {
					// not found return nil
					taskProto = nil
					return nil
				} else {
					return result.Error
				}
			}
			// marshall result to response
			return utils.StructToProto(model, taskProto)
		},
	)
	if err != nil {
		panic(err)
	}
	return &corndogsv1alpha1.GetTaskStateResponse{Task: taskProto}
}

func (s CockroachStore) GetNextTask(req *corndogsv1alpha1.GetNextTaskRequest) (*corndogsv1alpha1.GetNextTaskResponse, error){
	taskProto := &corndogsv1alpha1.Task{}
	var err error
	return &corndogsv1alpha1.GetNextTaskResponse{Task: taskProto}, err
}

func (s CockroachStore) UpdateTask(req *corndogsv1alpha1.UpdateTaskRequest) (*corndogsv1alpha1.UpdateTaskResponse, error){
	taskProto := &corndogsv1alpha1.Task{}
	var err error
	return &corndogsv1alpha1.UpdateTaskResponse{Task: taskProto}, err
}

func (s CockroachStore) CompleteTask(req *corndogsv1alpha1.CompleteTaskRequest) (*corndogsv1alpha1.CompleteTaskResponse, error){
	taskProto := &corndogsv1alpha1.Task{}
	var err error
	return &corndogsv1alpha1.CompleteTaskResponse{Task: taskProto}, err
}

func (s CockroachStore) CancelTask(req *corndogsv1alpha1.CancelTaskRequest) (*corndogsv1alpha1.CancelTaskResponse, error){
	taskProto := &corndogsv1alpha1.Task{}
	var err error
	return &corndogsv1alpha1.CancelTaskResponse{Task: taskProto}, err
}
