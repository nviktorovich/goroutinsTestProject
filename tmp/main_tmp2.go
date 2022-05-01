package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

const Path = "RoutinesExercise/1/tmp/"

var wg sync.WaitGroup

type File struct {
	name    string
	filter  bool
	contain string
}

func CreateFile(idx int, in chan File) {
	defer wg.Done()
	filePath := Path + strconv.Itoa(idx) + ".txt"
	file := File{}
	file.name = filePath
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	file.contain = string(data)
	in <- file

}

func FilterFile(idx int, in chan File, out chan File) {
	//defer wg.Done()
	for {
		file, ok := <-in
		fmt.Println(idx, "sorter work with ", file)
		if !ok {
			close(out)
			return
		}
		file.filter = Filter([]byte(file.contain))
		if file.filter {
			out <- file
		}
	}
}

func ShowFiles(in <-chan File) {
	for file := range in {
		fmt.Println(file.name, file.contain)
	}
}

func main() {

	c1 := make(chan File, 10)
	c2 := make(chan File, 10)

	for w := 0; w < 10; w++ {
		wg.Add(1)
		go CreateFile(w, c1)
	}
	wg.Wait()
	close(c1)

	for w := 0; w < 2; w++ {
		go FilterFile(w, c1, c2)
	}
	ShowFiles(c2)

	wg.Wait()

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
