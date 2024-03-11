package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 5
const MaxEatingCount = 2

type State int

const (
	THINKING State = iota
	HUNGRY
	EATING
)

var (
	state               [N]State
	criticalRegionMutex sync.Mutex
	outputMutex         sync.Mutex
	bothForksAvailable  = make([]chan struct{}, N)
	forksAvailable      = make([]bool, N)
	eatingCount         = make([]int, N)
	wg                  sync.WaitGroup
)

func left(i int) int {
	return (i - 1 + N) % N
}

func right(i int) int {
	return (i + 1) % N
}

func test(i int) {
	if state[i] == HUNGRY && state[left(i)] != EATING && state[right(i)] != EATING {
		state[i] = EATING
		if forksAvailable[i] {
			close(bothForksAvailable[i])
			forksAvailable[i] = false
		}
	}
}

func philosopher(i int) {
	defer wg.Done()
	for eatingCount[i] < MaxEatingCount {
		fmt.Printf("Philosopher %d is thinking..\n", i)
		time.Sleep(time.Millisecond * time.Duration(random(100, 500)))

		fmt.Printf("Philosopher %d is hungry..\n", i)
		state[i] = HUNGRY
		test(i)

		<-bothForksAvailable[i]
		fmt.Printf("Philosopher %d acquired forks and is eating\n", i)

		state[i] = THINKING
		fmt.Printf("Philosopher %d finished eating and released forks\n", i)
		test(left(i))
		test(right(i))

		eatingCount[i]++

		time.Sleep(time.Millisecond * 10)
	}
}

func random(min, max int) int {
	rnd := time.Now().UnixNano()
	return int(rnd%(int64(max)-int64(min)+1)) + min
}

func main() {
	for i := 0; i < N; i++ {
		bothForksAvailable[i] = make(chan struct{})
		forksAvailable[i] = true
		go philosopher(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("All philosophers have finished eating. Exiting")
}
