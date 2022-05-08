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

func main() {
	begin := time.Now()
	for i := 0; i < 1000; i++ {
		name := Path + strconv.Itoa(i) + ".txt"
		data, err := os.ReadFile(name)
		if err != nil {
			log.Println(err)
		}
		file := File{name: name, filter: Filter(data), contain: string(data)}
		if file.filter {
			fmt.Println(file.name)
			fmt.Println(file.contain)
			fmt.Println()
		}

	}
	fmt.Println(time.Since(begin))
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
