package main

import (
	"fmt"
	"log"

	storage "github.com/mylatko/learning_1/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Uploaded !", file)
}
