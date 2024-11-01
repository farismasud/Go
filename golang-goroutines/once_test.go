package golang_goroutines

import (
	"fmt"
	"testing"
	"sync"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func ()  {
			group.Add(1)
			go once.Do(OnlyOnce)
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter", counter)
}