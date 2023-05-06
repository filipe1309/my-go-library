package mygolibrary

import (
	"fmt"

	mygolib2 "github.com/filipe1309/my-go-library-2"
)

func HelloFrom110() {
	fmt.Println("Hello from My Go Library v1.1.0")
}

func HelloFromL1v120() {
	fmt.Println("Hello from My Go Library v1.2.0")
	mygolib2.HelloFromL2v100()
}
