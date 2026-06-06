package main

import (
	"fmt"
	"log"

	"github.com/niviten/wfe/internal/config"
	"github.com/niviten/wfe/internal/db"
	"github.com/niviten/wfe/internal/wfe"
)

func main() {
	config.Init()
	db.Init()

	wf := wfe.New()

	output, err := wf.Insert(wfe.NewRunnableAction(func(inputs any) (any, error) {
		// log.Println("action 1")
		value, ok := inputs.(int64)
		if !ok {
			return nil, fmt.Errorf("Expected int64, got %v", inputs)
		}
		return value + 2, nil
	})).Insert(wfe.NewRunnableAction(func(inputs any) (any, error) {
		// log.Println("action 2")
		value, ok := inputs.(int64)
		if !ok {
			return nil, fmt.Errorf("Expected int64, got %v", inputs)
		}
		return value * 5, nil
	})).Execute(int64(11))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("output: %v\n", output)

	log.Println("___done")
}
