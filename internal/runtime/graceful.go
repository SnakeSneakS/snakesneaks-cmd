package runtime

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var stopSignals = []os.Signal{syscall.SIGTERM, os.Interrupt, os.Kill}

type GracefulWorker interface {
	Run(context.Context, context.CancelFunc)
	Shutdown(context.Context)
}

func GracefulRun(worker GracefulWorker, shutdownTimeout time.Duration) {
	ctx, stop := signal.NotifyContext(context.Background(), stopSignals...)
	defer stop()

	go worker.Run(ctx, stop)

	<-ctx.Done()

	signal.Reset(stopSignals...) //this line allows "force quit"
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	worker.Shutdown(ctx)
}

func UnstoppableRun(worker GracefulWorker, shutdownTimeout time.Duration) {
	signal.Ignore(stopSignals...) // this makes it impossible to stop process

	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	go worker.Run(ctx, stop)
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	worker.Shutdown(ctx)
}
