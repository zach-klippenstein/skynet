package main

import (
	"fmt"
	"sync"
	"time"
)

func skynet(result *int, wg *sync.WaitGroup, num, size, div int) {
	if size == 1 {
		*result = num
	} else {
		results := make([]int, div)
		sum := 0
		var sub_wg sync.WaitGroup
		sub_wg.Add(div)
		for i := 0; i < div; i++ {
			sub_num := num + i*(size/div)
			go skynet(&(results[i]), &sub_wg, sub_num, size/div, div)
		}
		sub_wg.Wait()
		for _, r := range results {
			sum += r
		}
		*result = sum
	}
	wg.Done()
}

func main() {
	var result int
	var wg sync.WaitGroup
	wg.Add(1)
	start := time.Now().UnixNano() / 1000000
	go skynet(&result, &wg, 0, 1000000, 10)
	wg.Wait()
	end := time.Now().UnixNano() / 1000000
	fmt.Printf("Result: %d in %d ms.\n", result, end-start)
}
