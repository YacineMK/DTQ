package retry

import (
	"context"
	"log"
	"time"
)

type Config struct {
	MaxRetries int
	InitWait   time.Duration
	MaxWait    time.Duration
}

func NewRetryConfig(maxRetries int, initWait, maxWait time.Duration) *Config {
	return &Config{
		MaxRetries: maxRetries,
		InitWait:   initWait,
		MaxWait:    maxWait,
	}
}

func WithBackoff(ctx context.Context, cfg Config, operation func() error) error {
	var err error
	waitTime := cfg.InitWait

	for try := 0; try <= cfg.MaxRetries; try++ {
		if try > 0 {
			log.Printf("Retry attempt %d/%d. Waiting %v before retrying.", try, cfg.MaxRetries, waitTime)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(waitTime):
			}

			waitTime *= 2
			if waitTime > cfg.MaxWait {
				waitTime = cfg.MaxWait
			}
		}

		if err = operation(); err == nil {
			return nil
		}

		log.Printf("Operation failed on attempt %d/%d: %v", try+1, cfg.MaxRetries+1, err)
	}

	return err
}
