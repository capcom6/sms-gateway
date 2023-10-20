package tasks

import (
	"context"
	"fmt"
	"time"

	microbase "bitbucket.org/soft-c/gomicrobase"
	"gorm.io/gorm"
)

func init() {
	microbase.RegisterOnStartedListener(task)
}

func task(ctx context.Context, c chan error, d *gorm.DB) error {
	ticker := time.NewTicker(60 * time.Second)
	go func() {
		defer func() {
			c <- nil
		}()

		for {
			select {
			case <-ticker.C:
				fmt.Println("tick")
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
	return nil
}
