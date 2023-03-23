package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// data pertama dengan tipe interface{}
	data1 := []interface{}{"coba1", "coba2", "coba3"}

	// data kedua dengan tipe interface{}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

	// membuat goroutine untuk data pertama dan kedua secara bergantian
	for i := 1; i <= 8; i++ {
		wg.Add(2)

		// goroutine untuk data pertama
		go func(data []interface{}, id int) {
			defer wg.Done()

			mutex.Lock()
			if id%2 == 0 {
				fmt.Println(data, id+1)
			}
			mutex.Unlock()
		}(data1, i%4)

		// goroutine untuk data kedua
		go func(data []interface{}, id int) {
			defer wg.Done()

			mutex.Lock()
			if id%2 == 1 {
				fmt.Println(data, id+1)
			}
			mutex.Unlock()
		}(data2, i%4)

		wg.Wait()
	}
}
