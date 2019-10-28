package sdk_types

import (
	"github.com/insight-chain/inb-go/common"
	"math/big"
)

type Configure struct {
	Url string `json:"url"`
}

type Account struct {
	Nonce     string    `json:"nonce"`
	Balance   string    `json:"balance"`
	Root      string    `json:"root"`
	CodeHash  string    `json:"codeHash"`
	Res       Resource  `json:"res"`
	Stakings  []Staking `json:"stakings"`
	UnStaking UnStaking `json:"unStaking"`
	//Regular
	Voted                        string `json:"voted"`
	LastReceivedVoteRewardHeight string `json:"lastReceivedVoteRewardHeight"`
}
type Resource struct {
	Used         *big.Int // used
	Usable       *big.Int // unuse
	StakingValue *big.Int // total number of mortgage
	Height       *big.Int
}
type Staking struct {
	Hash               common.Hash // transaction of regular mortgaging
	StartHeight        *big.Int    // start time
	LockHeights        *big.Int    // duration of mortgaging
	Value              *big.Int    // amount of mortgaging
	Received           *big.Int    // amount of already received value
	LastReceivedHeight *big.Int    // Last receive time
}

type UnStaking struct {
	StartHeight *big.Int // start time
	Value       *big.Int // amount of redeeming
}
