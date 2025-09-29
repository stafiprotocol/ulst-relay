package stake_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type CustomStakePool struct {
	StakePool
	address common.Address
}

// NewStakePool creates a new instance of StakePool, bound to a specific deployed contract.
func NewCumstomStakePool(address common.Address, backend bind.ContractBackend) (*CustomStakePool, error) {
	contract, err := bindStakePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustomStakePool{
		address: address,
		StakePool: StakePool{
			StakePoolCaller:     StakePoolCaller{contract: contract},
			StakePoolTransactor: StakePoolTransactor{contract: contract},
			StakePoolFilterer:   StakePoolFilterer{contract: contract},
		},
	}, nil
}

func (s *CustomStakePool) Address() common.Address {
	return s.address
}
