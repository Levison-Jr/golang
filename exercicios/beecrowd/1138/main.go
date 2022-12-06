package main

import (
	"fmt"
	"strconv"
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
			sb := []byte(s)

			for _, value := range sb {
				switch value {
				case 48:
					cont[0]++
				case 49:
					cont[1]++
				case 50:
					cont[2]++
				case 51:
					cont[3]++
				case 52:
					cont[4]++
				case 53:
					cont[5]++
				case 54:
					cont[6]++
				case 55:
					cont[7]++
				case 56:
					cont[8]++
				case 57:
					cont[9]++
				}
			}
		}

		fmt.Println(cont)
		cont = make([]int, 10)
	}
}
