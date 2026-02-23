package runner

import (
	"context"
	"fmt"
	"time"
)

func (r *Runner) Scheduler(ctx context.Context) error {
	return r.SchedulerWithCtx(ctx)
}

func (r *Runner) SchedulerWithCtx(ctx context.Context) error {
	for {
		loc, _ := time.LoadLocation("Europe/Paris")
		now := time.Now().In(loc)

		if now.Format("15:04:05") == "00:00:00" {
			fmt.Println("send")
		} else {
			fmt.Println("pas encore")
		}
	}
}
