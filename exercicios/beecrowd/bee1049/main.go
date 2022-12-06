package main

import "fmt"

func main() {
	var (
		tipo1, tipo2, tipo3 string
	)

	fmt.Scan(&tipo1, &tipo2, &tipo3)
	if tipo1 == "vertebrado" {
		if tipo2 == "ave" {
			switch tipo3 {
			case "carnivoro":
				fmt.Printf("aguia\n")
			case "onivoro":
				fmt.Printf("pomba\n")
			}
		} else if tipo2 == "mamifero" {
			switch tipo3 {
			case "onivoro":
				fmt.Printf("homem\n")
			case "herbivoro":
				fmt.Printf("vaca\n")
			}
		}
	} else if tipo1 == "invertebrado" {
		if tipo2 == "inseto" {
			switch tipo3 {
			case "hematofago":
				fmt.Printf("pulga\n")
			case "herbivoro":
				fmt.Printf("lagarta\n")
			}
		} else if tipo2 == "anelideo" {
			switch tipo3 {
			case "hematofago":
				fmt.Printf("sanguessuga\n")
			case "onivoro":
				fmt.Printf("minhoca\n")
			}
		}
	}
}
