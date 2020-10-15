package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now().Hour()
	b := time.Now().Minute()
	c := time.Now().Second()
	fmt.Println(a, b, c)
}
