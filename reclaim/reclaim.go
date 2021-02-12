package reclaim

import (
	"fmt"

	"github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/curl"
	"github.com/iotaledger/iota.go/kerl"
	. "github.com/iotaledger/iota.go/signing"
	"github.com/iotaledger/iota.go/signing/key"
	"github.com/iotaledger/iota.go/signing/utils"
	"github.com/iotaledger/iota.go/trinary"
)

func Search(seed string, endIndex int) ReclaimAddresses {
	var results ReclaimAddresses
	for i := 0; i <= endIndex; i++ {
		matchAddress(getCurlPAddress(seed, i), &results)
		matchAddress(getKerlAddress(seed, i), &results)
		fmt.Printf("\rchecking address #%d", i)
	}
	fmt.Println()
	return results
}

func matchAddress(addr string, matches *ReclaimAddresses) {

	for _, reclAddr := range Reclaims.Addresses {
		if addr == reclAddr.Address {
			matches.Add(reclAddr)
		}
	}

}

func getCurlPAddress(seed string, index int) string {
	algo := curl.NewCurlP27()
	return getAddress(seed, uint64(index), algo)
}

func getKerlAddress(seed string, index int) string {
	algo := kerl.NewKerl()
	return getAddress(seed, uint64(index), algo)
}

func getAddress(seed string, index uint64, algo sponge.SpongeFunction) string {
	subseed, err := Subseed(seed, uint64(index), algo)
	must(err)

	prvKey, err := key.Sponge(subseed, consts.SecurityLevelMedium, algo)
	must(err)

	digests, err := Digests(prvKey, algo)
	must(err)

	addressTrits, err := Address(digests, algo)
	must(err)

	address := trinary.MustTritsToTrytes(addressTrits)
	return address
}

func must(err error) {
	if err != nil {
		panic(err)
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
