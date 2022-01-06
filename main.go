package main

import (
	"fmt"
	"os"

	"github.com/iota-community/iotaZeroBalanceHelper/reclaim"
	"github.com/iota-community/iotaZeroBalanceHelper/userIO"
	"github.com/iota-community/iotaZeroBalanceHelper/balance"
	"github.com/iota-community/iotaZeroBalanceHelper/address"
)




func main() {
	welcome()	
	seed := userIO.GetSeed()
	mainMenu (seed)
	

}

func welcome (){
	fmt.Println ("This program is a collection of tools that can help you in case your Iota Trinity wallet unexpectedly shows zero balance.")
	fmt.Println ("It requires the input of your seed. For your own safety you should not run it on a computer that is connected to the internet.")
	fmt.Println ("Use at your own risk!")
	if !userIO.GetConfirmation ("Do you want to continue?") {
		os.Exit(0)
	}
}

func mainMenu(seed string) {
	for {
		opt := userIO.GetOption("Choose from available options:", []string{"check balance", "check for possible reclaim", "Exit"})
		switch opt {
			case 1:
				checkKerlBalance(seed)
			case 2:
				checkCurlPBalance(seed)
			case 3: 
				checkReclaim(seed)	
			case 4:
				os.Exit(0)
		}		
	}		
}

func checkKerlBalance(seed string) {
	balance.LoadSnapshot()
	addrCount := userIO.GetNumberOfAddresses ()
	fmt.Println("\nGenerating addresses")
	addrs := address.GetKerlAddresses(seed,addrCount)
	total, bals := balance.GetBalance(addrs)
	if total == 0 {
		fmt.Printf("\nNo funds were found on the first %d addresses. Either the funds have been moved or you entered the wrong seed.\n", addrCount)

		
	} else{
		fmt.Printf("\nA total balance of %d was found on the first %d addresses.\n", total, addrCount)
		for addr, bal := range bals {
			fmt.Printf("balance: %d on address: %s\n",bal,addr)
		}
	}
	if userIO.GetConfirmation ("Do you want to export the generated addresses for reference?") {
		userIO.ExportAddresses(addrs)
	}
	userIO.WaitforEnter()
}


func checkCurlPBalance(seed string) {
	balance.LoadSnapshot()
	addrCount := userIO.GetNumberOfAddresses ()
	fmt.Println("\nGenerating addresses")
	addrs := address.GetCurlPAddresses(seed,addrCount)
	total, bals := balance.GetBalance(addrs)
	if total == 0 {
		fmt.Printf("\nNo funds were found on the first %d addresses. Either the funds have been moved or you entered the wrong seed.\n", addrCount)

		
	} else{
		fmt.Printf("\nA total balance of %d was found on the first %d addresses.\n", total, addrCount)
		for addr, bal := range bals {
			fmt.Printf("balance: %d on address: %s\n",bal,addr)
		}
	}
	if userIO.GetConfirmation ("Do you want to export the generated addresses for reference?") {
		userIO.ExportAddresses(addrs)
	}
	userIO.WaitforEnter()
}


func checkReclaim(seed string) {

	addrCount := userIO.GetNumberOfAddresses ()

	fmt.Println("Starting search for matching reclaim addresses")
	results := reclaim.Search(seed, addrCount)
	fmt.Println("Finished search for matching reclaim addresses")
	fmt.Printf("\n%d addresses of the entered seed have been tested.\n\n", addrCount)
	if len(results.Addresses) == 0 {
		fmt.Println("No funds on the tested addresses have been taken into custody by the Iota Foundation.\nThis is not a reclaim case.")
	} else {
		var total uint64
		fmt.Println("This is a reclaim case. Funds on addresses of this seed have been taken into custody by the Iota Foundation.\n\nAffected addresses:\n")
		for _, res := range results.Addresses {
			fmt.Printf("balance: %d\t address: %s\t reason: %s\n", res.Balance, res.Address, res.Reason)
			total += res.Balance
		}
		fmt.Printf("\nTotal affected balance: %d\n\n", total)
		fmt.Println("It is not possible to determine if the funds have been reclaimed yet or not.")
	}	
	userIO.WaitforEnter()
}



