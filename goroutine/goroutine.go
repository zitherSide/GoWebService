package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		//fmt.Printf("%d", i)
	}
}

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	wg.Done()
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		//fmt.Printf("%c", i)
	}
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	wg.Done()
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func print2() {
	var wg sync.WaitGroup
	printNumbers2(&wg)
	printLetters2(&wg)
}

func goPrint2() {
	var wg sync.WaitGroup
	go printNumbers2(&wg)
	go printLetters2(&wg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()
}
