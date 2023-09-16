package cmd

import (
	"time"

	"github.com/snakesneaks/snakesneaks-cmd/internal/runtime"
	"github.com/snakesneaks/snakesneaks-cmd/internal/worker"
)

func Run() {
	dps := 100 * time.Millisecond

	var w runtime.GracefulWorker = worker.NewWorker(dps)
	runtime.UnstoppableRun(w, time.Second*10)
}
