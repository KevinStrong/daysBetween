package main

import (
	"fmt"

	"github.com/KevinStrong/daysbetween"
)

func main() {
	fmt.Println("Getting days between unix epoc of Jan 1st 1970 and today")
	daysFromEpoc := daysbetween.Get()
	fmt.Printf("Days since unix epoc: %d\n", daysFromEpoc)
}
