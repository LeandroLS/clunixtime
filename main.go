package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Primeiro commit. Testando :D")
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())
}
