// I give more than one iota for the Summer Olympics
// Here, I have given 11

package main

import "fmt"

func main() {

	const (
		_ = 1980 + (iota * 4)
		LosAngeles
		Seoul
		Barcelona
		Atlanta
		Sydney
		Athens
		Beijing
		London
		Rio
		Tokyo
	)

	fmt.Println("These cities hosted/[will host] the Summer Olympics...")
	fmt.Printf("%-18s %-18s \n", "City", "Year")
	fmt.Printf("%-18s %-18v \n", "Los Angeles", LosAngeles)
	fmt.Printf("%-18s %-18v \n", "Atlanta", Atlanta)
	fmt.Printf("%-18s %-18v \n", "Tokyo", Tokyo)

}
