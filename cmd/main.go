package main

import (
	"github.com/dlc-01/BackendCrypto/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	// Запуск приложения
	app.Run("0.0.0.0:8080")

	// Ожидание сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Завершение приложения
	app.Shutdown()
}
