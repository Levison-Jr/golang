package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		A, B int
		cont []int = make([]int, 10)
	)
	for {
		fmt.Scanln(&A, &B)
		if A == 0 && B == 0 {
			break
		}
		for i := A; i <= B; i++ {
			s := strconv.FormatInt(int64(i), 10)
			cont[0] += strings.Count(s, "0")
			cont[1] += strings.Count(s, "1")
			cont[2] += strings.Count(s, "2")
			cont[3] += strings.Count(s, "3")
			cont[4] += strings.Count(s, "4")
			cont[5] += strings.Count(s, "5")
			cont[6] += strings.Count(s, "6")
			cont[7] += strings.Count(s, "7")
			cont[8] += strings.Count(s, "8")
			cont[9] += strings.Count(s, "9")
		}
		for _, v := range cont {
			fmt.Print(v, " ")
		}
		fmt.Println()
		cont = make([]int, 10)
	}
}
