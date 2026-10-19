package limiter

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrInvalidCapacity = errors.New("capacity and refill rate must be greater than zero")
	ErrEmptyKey        = errors.New("rate limit key cannot be empty")
)

type TokenBucket struct {
	mu         sync.Mutex
	capacity   float64
	tokens     float64
	refillRate float64 // tokens per second
	lastRefill time.Time
}

func NewTokenBucket(capacity float64, refillRate float64) (*TokenBucket, error) {
	if capacity <= 0 || refillRate <= 0 {
		return nil, ErrInvalidCapacity
	}
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}, nil
}

func (tb *TokenBucket) Allow(cost float64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.lastRefill = now

	tb.tokens = tb.tokens + (elapsed * tb.refillRate)
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	if tb.tokens >= cost {
		tb.tokens -= cost
		return true
	}
	return false
}

type DistributedLimiter struct {
	mu      sync.RWMutex
	buckets map[string]*TokenBucket
	rate    float64
	burst   float64
}

func NewDistributedLimiter(burst float64, rate float64) (*DistributedLimiter, error) {
	if burst <= 0 || rate <= 0 {
		return nil, ErrInvalidCapacity
	}
	return &DistributedLimiter{
		buckets: make(map[string]*TokenBucket),
		rate:    rate,
		burst:   burst,
	}, nil
}

func (dl *DistributedLimiter) Allow(key string) (bool, error) {
	if key == "" {
		return false, ErrEmptyKey
	}

	dl.mu.RLock()
	bucket, exists := dl.buckets[key]
	dl.mu.RUnlock()

	if !exists {
		dl.mu.Lock()
		bucket, exists = dl.buckets[key]
		if !exists {
			var err error
			bucket, err = NewTokenBucket(dl.burst, dl.rate)
			if err != nil {
				dl.mu.Unlock()
				return false, err
			}
			dl.buckets[key] = bucket
		}
		dl.mu.Unlock()
	}

	return bucket.Allow(1.0), nil
}
