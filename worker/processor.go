package worker

import (
	"context"

	db "github.com/WillChen936/simple-bank/db/sqlc"
	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProccessor interface {
	Start() error
	ProccessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProccessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedisTaskProccessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProccessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
		},
	)

	return &RedisTaskProccessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProccessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, processor.ProccessTaskSendVerifyEmail)

	return processor.server.Start(mux)
}
