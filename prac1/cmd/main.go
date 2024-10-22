package main

import (
	"fmt"
	"github.com/kartaevana/BolshoiGolang/prac1/internal/pkg/storage"
	"log"
)

func main() {
	s, err := storage.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	s.Set("key1", "val1")

	res1 := s.Get("key1")
	res2 := s.Get("key2")
	fmt.Println(*res1, res2)
}
