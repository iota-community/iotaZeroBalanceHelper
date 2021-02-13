package userIO

import (
	"fmt"
	"strings"
	"strconv"	
	"encoding/json"
	"os"


	
)

const seedLen = 81

func GetNumberOfAddresses () int {
	var endIndex int 
	opt := GetOption("Choose number of addresses to check:", []string{"100 addresses (recommended)", "1000 addresses (only required if the seed was heavily used)", "custom number of addresses"})
	switch opt {
		case 1:
			endIndex = 100
		case 2: 
			endIndex = 1000	
		case 3:
			endIndex = GetInt ("Enter number of addresses")	
	} 
	
	return endIndex
}

func WaitforEnter () {
	fmt.Println ("\nPress Enter to continue")
	fmt.Scanln ()
}

func GetSeed () string {
	var seed string
	var answer string
	for {
		fmt.Print("\nEnter seed: ")
		fmt.Scanln(&seed) 
		if !hasInvalidChars(seed){
			if !hasInvalidChars(seed){
			if len (seed) == seedLen {
				return seed
			}
			if len (seed) < seedLen {
				fmt.Printf("The seed has less than %d characters. Do you want to continue with this seed? (y/n): ", seedLen	)	
				seed = seed + 	strings.Repeat("9",seedLen-len(seed)	)   
			}
			if len (seed) > seedLen {
				fmt.Printf("The seed has more than %d characters. Do you want to continue with this seed? (y/n): ", seedLen)
				seed = seed[0:81]
			}
			fmt.Scanln(&answer)	
			if answer == "y" {
				return seed
			}   
		}
						   } else{
							   fmt.Println("\nValid seeds only contain upper case letters A-Z and the number 9.")
						   }
	}

	
	return seed
}	
	

	
func hasInvalidChars (seed string) bool {
	for _,r := range(seed){
		if (r < 'A' || r > 'Z') && r != '9' {
			return true
		}
	}
	return false

}


func GetOption(text string, options []string) int {
	var sel string
	
	fmt.Println("\n" +text)
	for i, descr := range options {
		fmt.Printf("%d: %s\n", i+1, descr)

	}
	
	for {
		fmt.Printf("\nSelect option (1-%d): ", len (options))
		fmt.Scanln(&sel)
		index := getOptIndex(sel, len(options))
		
		if index != -1 {
			return index
		}
		fmt.Printf("Invalid input. ")
	}
}


func GetConfirmation (text string) bool {
	var answer string
	fmt.Printf ("\n%s (y/n): ",text)
	for {
		fmt.Scanln(&answer)
		if answer == "y"{
			return true
		}
		if answer == "n" {
			return false
		}
		fmt.Print("Invalid input. Try again (y/n): ")
	}
	
}


func GetInt (text string) int {
	var input string
	fmt.Printf ("%s: ", text)
	for {
		fmt.Scanln(&input)	
		index, err := strconv.Atoi(input)
		if err == nil {
			return index
		}
		fmt.Print("Invalid input. Please enter a number: ")
	}
	
}
func getOptIndex(sel string, count int) int {

	index, err := strconv.Atoi(sel)
	if err == nil && index > 0 && index <= count {
		return index 
	}
	return -1
}

func ExportAddresses (addrs []string) {
	j, _ := json.Marshal(addrs)
	f, err := os.Create("addressExport.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString(string(j))
    if err != nil {
        fmt.Println(err)
        fmt.Println(l)
        f.Close()
        return
    }
}