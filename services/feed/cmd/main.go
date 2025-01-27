package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	pb "github.com/SergeyBogomolovv/read-advisor/lib/api/gen/books"
	"github.com/SergeyBogomolovv/read-advisor/services/feed/internal/config"
	"github.com/SergeyBogomolovv/read-advisor/services/feed/internal/handler"
	"github.com/SergeyBogomolovv/read-advisor/services/feed/internal/service"
	"github.com/SergeyBogomolovv/read-advisor/services/feed/internal/storage"
	"github.com/SergeyBogomolovv/read-advisor/services/feed/pkg/amqp"
	"github.com/SergeyBogomolovv/read-advisor/services/feed/pkg/books"
	"github.com/SergeyBogomolovv/read-advisor/services/feed/pkg/db"
	"google.golang.org/grpc"
)

func main() {
	conf := config.New()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	amqpConn := amqp.MustNew(conf.AmqpURL)
	defer amqpConn.Close()

	booksConn := books.MustNew(conf.BooksGrpcURL)
	defer booksConn.Close()

	db := db.MustNew(conf.PostgresURL)
	defer db.Close()

	listener, err := net.Listen("tcp", conf.Addr)
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewStorage(db)
	booksClient := pb.NewBooksClient(booksConn)

	prefSvc := service.NewPreferenceService(logger, booksClient, storage)
	feedSvc := service.NewFeedService(logger, booksClient, storage)

	rmq := handler.NewRabbitMQHandler(logger, amqpConn, prefSvc)
	feed := handler.NewFeedHandler(logger, feedSvc)

	srv := grpc.NewServer()
	feed.Register(srv)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		srv.GracefulStop()
		logger.Info("server stopped")
	}()

	go rmq.Consume(ctx)
	logger.Info("server started", "addr", conf.Addr)
	srv.Serve(listener)
	wg.Wait()
}
