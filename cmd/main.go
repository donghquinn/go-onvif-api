package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"org.donghyuns.com/onvif/ptz/configs"
	"org.donghyuns.com/onvif/ptz/internal/network"
	"org.donghyuns.com/onvif/ptz/internal/utils"
)

func main() {
	// 종료 신호를 받을 채널 생성
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	godotenv.Load(".env")

	if err := setEnv(); err != nil {
		slog.Error(fmt.Sprintf("read config err: %v", err))
		return
	}

	// Setup logger with environment-based configuration
	if err := utils.SetupGlobalLogger("logs", 1000, 1000, configs.GlobalConfig.Env); err != nil {
		slog.Error(fmt.Sprintf("Failed to setup logger: %v", err))
		return
	}

	server := network.Network()

	go func() {
		slog.Debug("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		slog.Debug(fmt.Sprintf("[DEBUG] App Host %s", configs.GlobalConfig.AppPort))
		slog.Debug(fmt.Sprintf("[START] Server Listening On: %s", configs.GlobalConfig.AppHost))
		slog.Debug("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Start Listening Error: %v", err)
		}
	}()

	<-quit
	slog.Info("Received Shut Down Signal")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error(fmt.Sprintf("[START] Start Failed Graceful Shutdown: %v", err))
		return
	}

	slog.Info("Server Has been Shutdown Gracefully")
}

func setEnv() error {
	if err := configs.ReadGlobalCfg(); err != nil {
		return fmt.Errorf("read global config err: %v", err)
	}
	if err := configs.ReadDatabaseCfg(); err != nil {
		return fmt.Errorf("read database config err: %v", err)
	}

	return nil
}
