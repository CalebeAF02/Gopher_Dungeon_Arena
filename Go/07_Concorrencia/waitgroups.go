package goguide

import "sync"

// ProcessItems usa um WaitGroup para aguardar todas as goroutines terminarem.
func ProcessItems(items []int) int {
	var wg sync.WaitGroup
	total := 0
	var mu sync.Mutex

	wg.Add(len(items))
	for _, item := range items {
		go func(v int) {
			defer wg.Done()
			mu.Lock()
			total += v
			mu.Unlock()
		}(item)
	}
	wg.Wait()
	return total
}

// WaitGroupExample demonstra a estrutura básica de sincronização com WaitGroup.
func WaitGroupExample(count int) int {
	var wg sync.WaitGroup
	sum := 0
	wg.Add(count)

	for i := 1; i <= count; i++ {
		go func(v int) {
			defer wg.Done()
			sum += v
		}(i)
	}
	wg.Wait()
	return sum
}
