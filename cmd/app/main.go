package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Dreker052/email-delivery-service/internal/config"
	"github.com/Dreker052/email-delivery-service/internal/service"
	"github.com/hibiken/asynq"
)

type EmailTaskPayload struct {
	ToEmail string `json:"to_email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

const TaskTypeEmail = "email:delivery"

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load env file: %s", err.Error())
	}

	sender := service.NewSender(cfg)

	redisOpt := asynq.RedisClientOpt{Addr: cfg.RedisAddr}

	srv := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"emails": 10,
			},
		},
	)

	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskTypeEmail, func(ctx context.Context, t *asynq.Task) error {
		var p EmailTaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}

		log.Printf("Sending email to %s...", p.ToEmail)

		return sender.Send(p.ToEmail, p.Subject, p.Body)
	})

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
