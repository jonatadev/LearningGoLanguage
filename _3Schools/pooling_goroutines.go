package main

import (
	"runtime"
	"sync"
	"github.com/montanaflynn/stats"
)

func multiMedian(vectors [][]float64) {
	var wg sync.WaitGroup
	wg.Add(len(vectors))
	ch := make(chan []float64)
	for i := 0; i < runtime.NumCPU(); i++ {
		go poolWorker(ch, &wg)
	}
	for _, vec := range vectors {
		ch <- vec
	}
	wg.Wait()
	close(ch)
}

func poolWorker(ch <-chan []float64, wg *sync.WaitGroup) {
	for values := range ch {
		m, _ := stats.Median(values)
		log.Printf("Median %v -> %f", values, m)
		wg.Done()
	}
	log.Printf("Shutting Down")
}

func main() {
	vectors := [][]float64{
		{1.1, 2.2, 3.3},
		{2.2, 3.3, 4.4},
		{3.3, 4.4, 5.5},
		{4.4, 5.5, 6.6},
		{5.5, 6.6, 7.7},
	}
	multiMedian(vectors)
	time.Sleep(10 * time.Millisecond) //Let workers terminate
	fmt.Println("DONE")
}
