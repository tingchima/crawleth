// Package cronjob provides
package cronjob

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

// Job job function with context
type Job func(ctx context.Context)

func cmd(ctx context.Context, job Job) func() {
	return func() {
		job(ctx)
	}
}

// Cronjob .
type Cronjob struct {
	*cron.Cron
	sync.Mutex
	mapTable map[string]jobInfo
	ctx      context.Context
	cancel   context.CancelFunc
}

type jobInfo struct {
	EntryID cron.EntryID
	Spec    string
}

// NewCronjob .
func NewCronjob() (*Cronjob, error) {
	ctx, cancel := context.WithCancel(context.Background())
	return &Cronjob{
		Cron:     cron.New(cron.WithSeconds()),
		mapTable: make(map[string]jobInfo),
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// Shutdown graceful shutdown cron scheduler with 5 min timeout
func (c *Cronjob) Shutdown() error {
	c.cancel()
	ctx := c.Cron.Stop()
	timeout := time.After(5 * time.Minute)
	select {
	case <-ctx.Done():
		return nil
	case <-timeout:
		return errors.New("timeout")
	}
}

// AddJob .
func (c *Cronjob) AddJob(key string, spec string, job Job) (cron.EntryID, error) {
	entryID, err := c.Cron.AddFunc(spec, cmd(c.ctx, job))
	if err != nil {
		return 0, err
	}
	c.Lock()
	c.mapTable[key] = jobInfo{EntryID: entryID, Spec: spec}
	c.Unlock()

	fmt.Printf("scheduler add job %s with %s", key, spec)
	return entryID, nil
}
