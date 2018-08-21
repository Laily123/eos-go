package system

import (
	eos "github.com/eoscanada/eos-go"
)

// NewSetalimit sets the account limits. Requires signature from `eosio@active` account.
func NewSetalimit(account eos.AccountName, ramBytes, netWeight, cpuWeight int64) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setalimit"),
		Authorization: []eos.PermissionLevel{
			{Actor: eos.AccountName("eosio"), Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Setalimit{
			Account:   account,
			RAMBytes:  ramBytes,
			NetWeight: netWeight,
			CPUWeight: cpuWeight,
		}),
	}
	return a
}

// Setalimit represents the `eosio.system_contract::setalimit` action.
type Setalimit struct {
	Account   eos.AccountName `json:"account"`
	RAMBytes  int64           `json:"ram_bytes"`
	NetWeight int64           `json:"net_weight"`
	CPUWeight int64           `json:"cpu_weight"`
}
