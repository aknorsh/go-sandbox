package main

import (
	"context"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: %v", err)
			}
			log.Printf("Readfile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}

// Per wraps rate.Every
func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

// Open opens api connection
func Open() *APIConnection {
	return &APIConnection{
		apiLimit: multiLimiter(
			rate.NewLimiter(Per(2, time.Second), 2),
			rate.NewLimiter(Per(10, time.Minute), 10),
		),
		diskLimit: multiLimiter(
			rate.NewLimiter(rate.Limit(1), 1),
		),
		networkLimit: multiLimiter(
			rate.NewLimiter(Per(3, time.Second), 3),
		),
	}
}

// APIConnection is fake api connection
type APIConnection struct {
	apiLimit,
	diskLimit,
	networkLimit RateLimiter
}

// ReadFile pretends read-file process
func (a *APIConnection) ReadFile(ctx context.Context) error {
	err := multiLimiter(a.apiLimit, a.diskLimit).Wait(ctx)
	if err != nil {
		return err
	}
	return nil
}

// ResolveAddress pretends resolve address process
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	err := multiLimiter(a.apiLimit, a.networkLimit).Wait(ctx)
	if err != nil {
		return err
	}
	return nil
}

// RateLimiter is rate limiter
type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

// MultiLimiter returns multi rate limiter.
func multiLimiter(limiters ...RateLimiter) *MultiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	return &MultiLimiter{limiters: limiters}
}

// MultiLimiter is multi rate limiter
type MultiLimiter struct {
	limiters []RateLimiter
}

// Wait waits all limiters in multiLimiter
func (l *MultiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Limit returns most urgent limiter
func (l *MultiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}
