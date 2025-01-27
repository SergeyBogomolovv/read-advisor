package handler

import (
	"context"
	"log"
	"log/slog"

	d "github.com/SergeyBogomolovv/read-advisor/services/feed/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PrefService interface {
	AddPreference(ctx context.Context, userID int64, bookID string, prefType d.PreferenceType) error
}

type rmqHandler struct {
	logger *slog.Logger
	conn   *amqp.Connection
	svc    PrefService
}

func NewRabbitMQHandler(logger *slog.Logger, conn *amqp.Connection, svc PrefService) *rmqHandler {
	return &rmqHandler{
		logger: logger,
		conn:   conn,
		svc:    svc,
	}
}

func (r *rmqHandler) Consume(ctx context.Context) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-msgs:
			log.Printf("Received a message: %s", msg.Body)
		}
	}
}
