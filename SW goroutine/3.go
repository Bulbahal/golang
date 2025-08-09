package main

import (
	"fmt"
	"sync"
	"time"
)

// 3. Конфигурация с горячей перезагрузкой
// Разработайте систему управления конфигурацией, которая позволяет:
// Часто читать параметры конфигурации (много горутин)
// Редко полностью перезагружать конфигурацию (одна горутина)
// Используйте sync.RWMutex.
type Config struct {
	mu   sync.RWMutex
	Name string
	Age  int
}

func (c *Config) GetConfig() (name string, age int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Name, c.Age
}
func (c *Config) Realoand(name string, age int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Name = name
	c.Age = age
}
func main() {
	config := &Config{Name: "Витя", Age: 28}
	for i := 0; i < 10; i++ {
		go func(id int) {
			for {
				confName, confAge := config.GetConfig()
				fmt.Println(id, confName, confAge)
				time.Sleep(5 * time.Second)
			}
		}(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Конфиг обнавлен")
	config.Realoand("Тимур", 10)
	time.Sleep(3 * time.Second)
}
