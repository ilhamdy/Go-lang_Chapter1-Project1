package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup
	wg.Add(2)

	go func(data []interface{}) {
		defer wg.Done()
		for i := 1; i <= 4; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(data, i)
		}
	}(data1)

	go func(data []interface{}) {
		defer wg.Done()
		for i := 1; i <= 4; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(data, i)
		}
	}(data2)

	wg.Wait()
}
