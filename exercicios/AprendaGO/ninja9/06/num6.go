package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	time.Sleep(5*time.Second)
}
