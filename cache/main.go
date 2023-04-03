package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	cache := NewCache(GetFibonacci)
	fibValues := []int{1, 45, 13, 17, 1, 45, 1}
	var wg sync.WaitGroup
	for _, value := range fibValues {
		go func(v int) {
			defer wg.Done()
			wg.Add(1)
			result, err := cache.GetData(v)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Printf("The Fibonnaci value for %d is %d\n", v, result)
		}(value)
	}

	wg.Wait()
}
