package reclaim

import (
	"fmt"

	. "github.com/HBMY289/iotaZeroBalanceHelper/address"
)

func Search(seed string, endIndex int) ReclaimAddresses {
	var results ReclaimAddresses
	for i := 0; i <= endIndex; i++ {
		matchAddress(GetCurlPAddress(seed, i), &results)
		matchAddress(GetKerlAddress(seed, i), &results)
		fmt.Printf("\rchecking address #%d", i)
	}
	fmt.Println()
	return results
}

func matchAddress(addr string, matches *ReclaimAddresses) {

	for _, reclAddr := range Reclaims.Addresses {
		if addr == reclAddr.Address {
			matches.Add(reclAddr)
			fmt.Println ("\nMatched address:", addr)
			return
		}
	}

}



func (rs *ReclaimAddresses) Add(r ReclaimAddress) {
	rs.Addresses = append(rs.Addresses, r)
}

type ReclaimAddresses struct {
	Addresses []ReclaimAddress
}

type ReclaimAddress struct {
	Address, Reason string
	Balance         uint64
}
