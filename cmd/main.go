package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/snakesneaks/snakesneaks-cmd/internal/worker"
)

func Run() {
	var w gracefulWorker = worker.NewWorker(100 * time.Millisecond)
	gracefulRun(w, time.Second*10)
}

type gracefulWorker interface {
	Run(context.Context, context.CancelFunc)
	Shutdown(context.Context)
}

func gracefulRun(worker gracefulWorker, shutdownDelay time.Duration) {
	ignoredSignals := []os.Signal{syscall.SIGTERM, os.Interrupt, os.Kill}
	ctx, stop := signal.NotifyContext(context.Background(), ignoredSignals...)
	defer stop()

	go worker.Run(ctx, stop)

	<-ctx.Done()

	//signal.Reset(ignoredSignals...) //if this line is comment out, it becomes unstoppable by Ctrl+C
	ctx, cancel := context.WithTimeout(context.Background(), shutdownDelay)
	defer cancel()

	worker.Shutdown(ctx)
}
