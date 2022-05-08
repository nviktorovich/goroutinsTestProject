package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const Path = "RoutinesExercise/1/tmp/"

type File struct {
	name    string
	filter  bool
	contain string
}

func SourceWorker(d chan File) {
	for i := 0; i < 1000; i++ {
		name := Path + strconv.Itoa(i) + ".txt"
		data, err := os.ReadFile(name)
		if err != nil {
			log.Println(err)
		}
		file := File{name: name, filter: Filter(data), contain: string(data)}
		d <- file
	}
	close(d)
}

func FilterWorker(in chan File, out chan File) {
	for {
		file, ok := <-in
		if !ok {
			close(out)
			return
		}
		if file.filter {
			out <- file
		}
	}
}

func RepresentWorker(out chan File) {
	for file := range out {
		fmt.Println(file.name)
		fmt.Println(file.contain)
		fmt.Println()
	}
}

func main() {
	start := time.Now()
	downChan := make(chan File)
	upChan := make(chan File)
	go SourceWorker(downChan)
	go FilterWorker(downChan, upChan)
	RepresentWorker(upChan)
	fmt.Println(time.Since(start))

}

func Filter(b []byte) bool {
	num, err := strconv.Atoi(string(b))
	if err != nil {
		log.Println(err)
		return false
	}
	if num > 6000000000000000000 {
		return true
	}
	return false
}
