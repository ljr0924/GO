package main

import (
	"fmt"
	"sync"
)

func main() {

	wg1 := new(sync.WaitGroup)
	fishCounter := 0
	catCounter := 0
	dogCounter := 0

	catCh := make(chan struct{}, 1)
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)

	wg1.Add(3)
	dogCh <- struct{}{}
	go printDog(wg1, dogCounter, dogCh, fishCh)
	go printFish(wg1, fishCounter, fishCh, catCh)
	go printCat(wg1, catCounter, catCh, dogCh)

	wg1.Wait()

	// *****************************************
	// 优化版本  通过配置 打印任意数量 任意顺序文本
	// *****************************************
	totalPrint := 4
	printOrder := []string{"one", "three", "two"}
	goroutineNum := len(printOrder)

	wg2 := new(sync.WaitGroup)
	chs := make([]chan struct{}, goroutineNum)
	for i := 0; i < goroutineNum; i++ {
		chs[i] = make(chan struct{}, 1)
	}
	wg2.Add(goroutineNum)
	for i := 0; i < goroutineNum; i++ {

		var curCh chan struct{}
		var nextCh chan struct{}
		if i == 0 {
			curCh = chs[0]
			nextCh = chs[1]
		} else if i == (goroutineNum - 1) {
			curCh = chs[goroutineNum - 1]
			nextCh = chs[0]
		} else {
			curCh = chs[i]
			nextCh = chs[i+1]
		}
		counter := 0
		go PrintAnimal(wg2, curCh, nextCh, printOrder[i], counter, totalPrint)
	}
	chs[0] <- struct{}{}

	wg2.Wait()

}

// 优化版本
func PrintAnimal(wg *sync.WaitGroup, curCh, nextCh chan struct{}, text string, counter, total int) {
	for {
		if counter == total {
			wg.Done()
			return
		}
		<- curCh
		fmt.Println(text)
		counter++
		nextCh <- struct{}{}
	}

}

// 顺序打印
func printDog(wg *sync.WaitGroup, counter int, dogCh, fishCh chan struct{}) {
	for {
		if counter == 10 {
			wg.Done()
			return
		}
		<- dogCh
		fmt.Println("dog")
		counter++
		fishCh <- struct{}{}
	}
}

func printFish(wg *sync.WaitGroup, counter int, fishCh, catCh chan struct{}) {
	for {
		if counter == 10 {
			wg.Done()
			return
		}
		<- fishCh
		fmt.Println("fish")
		counter++
		catCh <- struct{}{}
	}

}
func printCat(wg *sync.WaitGroup, counter int, catCh, dogCh chan struct{}) {
	for {
		if counter == 10 {
			wg.Done()
			return
		}
		<- catCh
		fmt.Println("cat")
		counter++
		dogCh <- struct{}{}
	}
}
