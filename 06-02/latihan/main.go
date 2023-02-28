package main

import (
	"fmt"
	"sync"
)

func incrementElement(m *sync.Mutex, wg *sync.WaitGroup, arr []int, counter *int) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}

	for i := 0; i < len(arr); i++ {
		arr[i] += i + 1
		wg2.Add(1)
		go func() {
			defer func() {
				m.Unlock()
				wg2.Done()
			}()
			m.Lock()
			*counter++
		}()
	}

	wg2.Wait()
}

func main() {
	arr := make([]int, 1000)
	wg := sync.WaitGroup{}
	m := &sync.Mutex{}
	counter := 0

	wg.Add(1)
	go incrementElement(m, &wg, arr, &counter)

	wg.Wait()

	fmt.Println("Finish")
	fmt.Println(arr)
	fmt.Println(counter)
}
