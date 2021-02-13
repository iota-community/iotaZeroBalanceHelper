package address

import (
	"github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/curl"
	"github.com/iotaledger/iota.go/kerl"
	"github.com/iotaledger/iota.go/signing"
	"github.com/iotaledger/iota.go/signing/key"
	"github.com/iotaledger/iota.go/signing/utils"
	"github.com/iotaledger/iota.go/trinary"
)


func GetKerlAddresses (seed string, addrCount int) []string{
	addrs := make([]string,addrCount,addrCount)
	for i:=0;i<addrCount;i++ {
		addrs[i] = GetKerlAddress(seed, i) 
	}
	return addrs
}

func GetCurlPAddress(seed string, index int) string {
	algo := curl.NewCurlP27()
	return getAddress(seed, uint64(index), algo)
}

func GetKerlAddress(seed string, index int) string {
	algo := kerl.NewKerl()
	return getAddress(seed, uint64(index), algo)
}

func getAddress(seed string, index uint64, algo sponge.SpongeFunction) string {
	subseed, err := signing.Subseed(seed, uint64(index), algo)
	must(err)

	prvKey, err := key.Sponge(subseed, consts.SecurityLevelMedium, algo)
	must(err)

	digests, err := signing.Digests(prvKey, algo)
	must(err)

	addressTrits, err := signing.Address(digests, algo)
	must(err)

	address := trinary.MustTritsToTrytes(addressTrits)
	return address
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}