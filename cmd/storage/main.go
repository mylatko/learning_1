package main

import (
	"fmt"

	storage "github.com/mylatko/learning_1/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("Hello!", st)
}
