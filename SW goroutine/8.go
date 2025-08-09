package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 8. Graceful shutdown флаг
// Реализуйте механизм graceful shutdown, где:
// Одна горутина может установить флаг завершения
// Множество горутин могут проверять этот флаг
// Требуется максимальная производительность при проверках (atomic).
type ShutdownFlag struct {
	state atomic.Int32
}

func NewShotdownFlag() *ShutdownFlag {
	return &ShutdownFlag{}
}
func (s *ShutdownFlag) SignalShutdown() {
	s.state.Store(1)
}
func (s *ShutdownFlag) ShoulShutdown() bool {
	if s.state.Load() == 1 {
		return true
	}
	return false
}
func worker(id int, wg *sync.WaitGroup, flag *ShutdownFlag) {
	defer wg.Done()
	for {
		if flag.ShoulShutdown() {
			fmt.Printf("worker %d завершает работу\n", id)
			return
		}
		time.Sleep(3 * time.Second)
		fmt.Printf("worker %d work\n", id)
	}
}
func main() {
	flag := NewShotdownFlag()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg, flag)
	}
	time.Sleep(5 * time.Second)
	flag.SignalShutdown()
	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}
