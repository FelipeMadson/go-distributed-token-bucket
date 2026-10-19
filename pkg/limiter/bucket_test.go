package limiter

import (
	"sync"
	"testing"
	"time"
)

func TestTokenBucketRefill(t *testing.T) {
	bucket, err := NewTokenBucket(2.0, 10.0)
	if err != nil {
		t.Fatalf("unexpected error creating bucket: %v", err)
	}

	if !bucket.Allow(1.0) {
		t.Errorf("expected first token to be allowed")
	}
	if !bucket.Allow(1.0) {
		t.Errorf("expected second token to be allowed")
	}
	if bucket.Allow(1.0) {
		t.Errorf("expected third token to be rejected")
	}

	time.Sleep(120 * time.Millisecond)

	if !bucket.Allow(1.0) {
		t.Errorf("expected token to be allowed after refill")
	}
}

func TestDistributedLimiterConcurrency(t *testing.T) {
	limiter, err := NewDistributedLimiter(50.0, 100.0)
	if err != nil {
		t.Fatalf("failed to create limiter: %v", err)
	}

	var wg sync.WaitGroup
	var allowedCount int64
	var mu sync.Mutex

	workers := 20
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				allowed, err := limiter.Allow("client-alpha")
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if allowed {
					mu.Lock()
					allowedCount++
					mu.Unlock()
				}
			}
		}()
	}

	wg.Wait()

	if allowedCount == 0 {
		t.Errorf("expected at least some requests to be allowed")
	}
}
