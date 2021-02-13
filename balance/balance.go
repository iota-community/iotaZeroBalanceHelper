package balance

import (
	"fmt"
	"io/ioutil"
)




func GetBalance (addresses []string) (uint64, map[string]uint64) {
	balances := make(map[string]uint64)
	var total uint64
	for _, addr := range(addresses){
		b := Snap.Balances[addr]
		if b != 0 {
			total += b
			balances[addr] = b
		}
	}
	return total, balances
}




func LoadSnapshot(){
	snapFile := "snapshot.txt"
	data, err := ioutil.ReadFile(snapFile)
	if err == nil {
		ParseJson([]byte(data))
		fmt.Printf("Loaded %d addresses from provided snapshot file.\n", len(Snap.Balances))
	}
}




type Snapshot struct {
	Balances       map[string]uint64 `json:"balances"`
	MilestoneIndex uint64
	Duration       int `json:"duration"`
}