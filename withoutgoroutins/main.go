package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type File struct {
	name    string
	filter  bool
	contain string
}

const Path = "RoutinesExercise/1/tmp/"

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
func FileInit(id int) {
	filename := Path + strconv.Itoa(id) + ".txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}
	filterFlag := Filter(data)
	file := File{name: filename, filter: filterFlag, contain: string(data)}
	fmt.Println(file)
	fmt.Println(id)
}

func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		FileInit(i)
	}

	fmt.Println(time.Since(start))
}
