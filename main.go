package main

import (
	"fmt"

	"github.com/HBMY289/iotaZeroBalanceHelper/reclaim"
)

var nodeURL = "https://nodeHBMY289.goip.de:443"

//var testSeed = "JQFDINFVANETQ9GXQXNOZXARAGKTERA9AM9MHPRDPAOMBRWLMN9TVGXQHEJOWEGODTNLTUKQTYE9RJFWU"

var testSeed = "KDHCTWOXCMSHGXTQWPSMGUHOUOZUERWDASNMVUROKLJOUPOPRZETWIWPS9CJG9H9A9GBNHJTZFEDJGLR9"

func main() {

	checkReclaim(testSeed, 100)

}

func checkReclaim(seed string, endIndex int) {

	fmt.Println("Starting search for matching reclaim addresses")
	results := reclaim.Search(testSeed, endIndex)
	fmt.Println("Finished search for matching reclaim addresses")
	fmt.Printf("\n%d addresses of the entered seed have been tested.\n\n", endIndex)
	if len(results.Addresses) == 0 {
		fmt.Println("No funds for the tested addresses been taken into custody by the Iota Foundation. This is not a reclaim case.")
		return
	}
	var total uint64
	fmt.Println("This is a reclaim case. Funds on addresses of this seed have been taken into custody by the Iota Foundation.\n\nAffected addresses:\n")
	for _, res := range results.Addresses {
		fmt.Printf("balance: %d\t address: %s\t reason: %s\n", res.Balance, res.Address, res.Reason)
		total += res.Balance
	}
	fmt.Printf("\nTotal affected balance: %d\n\n", total)
	fmt.Println("It is not possible to determine if the funds have been reclaimed yet or not.")

}
