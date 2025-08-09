package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//10. Rate limiter на atomic
//Реализуйте простой ограничитель запросов (rate limiter), который:
//Считает количество запросов в текущем интервале
//Сбрасывает счетчик по таймеру
//Должен работать максимально эффективно при проверках (atomic).

type RateLimiter struct {
	counter    atomic.Int32
	limit      int32
	resetEvery time.Duration
	stopflag   atomic.Bool
}

func NewRateLimiter(limit int32, every time.Duration) *RateLimiter {
	r := &RateLimiter{
		limit:      limit,
		resetEvery: every,
	}
	go r.resetLoop()
	return r
}
func (r *RateLimiter) resetLoop() {
	for !r.stopflag.Load() {
		time.Sleep(r.resetEvery)
		r.counter.Store(0)
	}
}
func (r *RateLimiter) Allow() bool {
	return r.counter.Add(1) <= r.limit
}
func (r *RateLimiter) Stop() {
	r.stopflag.Store(true)
}

func main() {
	limiter := NewRateLimiter(10, 1*time.Second)
	defer limiter.Stop()
	for i := 0; i < 5; i++ {
		go func(id int) {
			for !limiter.stopflag.Load() {
				if limiter.Allow() {
					fmt.Printf("%d: ok\n", id)
				} else {
					fmt.Printf("%d: fail\n", id)
				}
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}
	time.Sleep(5 * time.Second)
}
