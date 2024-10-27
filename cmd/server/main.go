package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/jeffreywu1996/go-user/internal/config"
	"github.com/jeffreywu1996/go-user/internal/handler"
	"github.com/jeffreywu1996/go-user/internal/repository"
	"github.com/jeffreywu1996/go-user/internal/service"
	"github.com/jeffreywu1996/go-user/pkg/logger"
	pb "github.com/jeffreywu1996/go-user/proto/user"
)

func main() {
	// Initialize logger
	if err := logger.Init(); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	log := logger.Get()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Initialize dependencies
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Create gRPC server
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, userHandler)

	// Start listening
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatal("Failed to listen", zap.Error(err))
	}

	// Handle shutdown gracefully
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		log.Info("Received shutdown signal")
		server.GracefulStop()
	}()

	log.Info("Starting gRPC server", zap.String("port", cfg.Server.Port))
	if err := server.Serve(lis); err != nil {
		log.Fatal("Failed to serve", zap.Error(err))
	}
}
