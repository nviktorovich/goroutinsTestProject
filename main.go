package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type File struct {
	name    string
	filter  bool
	contain string
}

const Path = "RoutinesExercise/1/tmp/"

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	c1 := make(chan int, 100000)
	c2 := make(chan File, 100000)

	for i := 0; i < 5; i++ {
		go FileInit(i, c1, c2)
		go FileReviewer(i, c2)
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		c1 <- i
	}
	wg.Wait()

	fmt.Println(time.Since(start))
}

func FileReviewer(idx int, in chan File) {
	defer close(in)
	for {
		fmt.Println(idx, "reviewer show file: ", <-in)
	}
}

func FileInit(worker int, input chan int, out chan File) {
	defer close(input)
	for {
		fileIdx := <-input
		fileName := Path + strconv.Itoa(fileIdx) + ".txt"
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			log.Println(err)
		}
		fileFlag := Filter(fileData)
		out <- File{name: fileName, filter: fileFlag, contain: string(fileData)}
		fmt.Println("worker ", worker, "work with", fileIdx)
		wg.Done()
	}
}

func Filter(b []byte) (ans bool) {
	num, err := strconv.Atoi(string(b))
	if err != nil {
		log.Println(err)
		return
	}
	if num > 6000000000000000000 {
		return true
	}
	return
}
