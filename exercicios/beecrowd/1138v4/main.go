package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var (
		A, B int64
		cont []int64 = make([]int64, 10)
	)
	for {
		fmt.Scanln(&A, &B)
		if A == 0 && B == 0 {
			break
		}
		for i := A; i <= B; i++ {
			s := strconv.FormatInt(i, 10)
			cont[0] += int64(bytes.Count([]byte(s), []byte("0")))
		}
		for _, v := range cont {
			fmt.Print(v, " ")
		}
		fmt.Println()
		cont = make([]int64, 10)
	}
}
