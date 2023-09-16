package worker

import (
	"context"
	"time"

	"github.com/snakesneaks/snakesneaks-cmd/internal/worker/drawer"
)

const (
	sneaking state = iota
	backhome
	end
)

type state int

type worker struct {
	DPF   time.Duration //duration per frame
	State state
}

func NewWorker(SPF time.Duration) *worker {
	return &worker{
		SPF,
		sneaking,
	}
}

func (w *worker) Run(ctx context.Context, stop context.CancelFunc) {
	drawer := drawer.NewDrawer(ctx)
	defer stop()
loop:
	for w.State == sneaking {
		select {
		case <-ctx.Done():
			break loop
		default:
			if done := drawer.DrawNext(); done {
				break loop
			}
			time.Sleep(w.DPF)
			//log.Println("snake is sneaking")
		}
	}
	w.State = end
	//log.Println("worker ended!")

	drawer.Clear()
}

func (w *worker) Shutdown(ctx context.Context) {
	//log.Println("shutting down...")
	if w.State != end {
		w.State = backhome
	loop:
		for w.State != end {
			select {
			case <-ctx.Done():
				break loop
			default:
				continue loop
			}
		}
	}

}
