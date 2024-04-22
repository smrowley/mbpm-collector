package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Work struct {
	//Id uuid.UUID
	Description string
	//Test func()
}

func test() {
	uuid, err := uuid.NewRandom()

	if err != nil {
		
	}

	fmt.Printf("uuid: %v\n", uuid)
}