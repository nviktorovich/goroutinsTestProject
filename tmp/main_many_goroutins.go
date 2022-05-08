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
	c1 := make(chan File, 100000)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go FileInit(i, c1)
	}
	wg.Wait()
	close(c1)

	for {
		ans, ok := <-c1
		if !ok {
			break
		}
		fmt.Println(ans)
	}
	fmt.Println(time.Since(start))
}

func FileInit(id int, c chan<- File) {
	wg.Add(1)
	defer wg.Done()
	filename := Path + strconv.Itoa(id) + ".txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}
	filterFlag := Filter(data)
	file := File{name: filename, filter: filterFlag, contain: string(data)}
	c <- file
	fmt.Println(id)
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
