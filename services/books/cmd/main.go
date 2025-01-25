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

	"github.com/SergeyBogomolovv/read-advisor/services/books/internal/config"
	"github.com/SergeyBogomolovv/read-advisor/services/books/internal/handler"
	"github.com/SergeyBogomolovv/read-advisor/services/books/internal/service"
	"google.golang.org/grpc"
)

func main() {
	config := config.New()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	service := service.NewBooksService(logger, config.APIKey)

	srv := grpc.NewServer()

	handler.NewGRPCHandler(logger, service).Serve(srv)

	listener, err := net.Listen("tcp", config.Addr)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		srv.GracefulStop()
		logger.Info("server stopped")
	}()

	logger.Info("server started", "addr", config.Addr)
	srv.Serve(listener)
	wg.Wait()
}
