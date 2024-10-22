package main

import (
	"fmt"
	"github.com/kartaevana/BolshoiGolang/ht1/internal/pkg/storage"
	"log"
)

func main() {
	s, err := storage.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	s.Set("key1", "val1")
	s.Set("2", "val2")

	res1 := s.Get("key1")
	res2 := s.Get("2")

	fmt.Println(res1, res2)
	fmt.Println(s.GetKind("2"), s.GetKind("key1"))
}
