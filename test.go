package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now().Format("20060102")
	fmt.Println(a)
}
