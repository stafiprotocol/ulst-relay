package stake_manager

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (_StakeManager *StakeManagerCaller) CurrentEraUint64(opts *bind.CallOpts) (uint64, error) {
	currentEra, err := _StakeManager.CurrentEra(opts)
	if err != nil {
		return 0, err
	}
	return currentEra.Uint64(), nil
}

func (_StakeManager *StakeManagerCaller) LatestEraUint64(opts *bind.CallOpts) (uint64, error) {
	latestEra, err := _StakeManager.LatestEra(opts)
	if err != nil {
		return 0, err
	}
	return latestEra.Uint64(), nil
}

type PoolInfo struct {
	Pool                         common.Address
	Era                          *big.Int
	Active                       *big.Int
	Bond                         map[common.Address]*big.Int
	Unbond                       map[common.Address]*big.Int
	FullfilledWithdrawalAmountOf map[common.Address]*big.Int
	Stablecoins                  []common.Address
}

func (p *PoolInfo) String() string {
	str := fmt.Sprintf("pool: %s\nera: %s\nactive: %s", p.Pool, p.Era, p.Active)
	for _, stablecoin := range p.Stablecoins {
		str += fmt.Sprintf("\n  stablecoin: %s", stablecoin)
		str += fmt.Sprintf("\n    bond: %s", p.Bond[stablecoin])
		str += fmt.Sprintf("\n    unbond: %s", p.Unbond[stablecoin])
		str += fmt.Sprintf("\n    fullfilledWithdrawalAmount: %s", p.FullfilledWithdrawalAmountOf[stablecoin])
	}
	return str
}

func (_StakeManager *StakeManagerCaller) GetPoolInfoOf(opts *bind.CallOpts, pool common.Address) (*PoolInfo, error) {
	_poolInfo, err := _StakeManager.PoolInfoOf(opts, pool)
	if err != nil {
		return nil, err
	}
	stablecoins, err := _StakeManager.GetStablecoins(opts)
	if err != nil {
		return nil, err
	}

	poolInfo := PoolInfo{
		Pool:                         pool,
		Era:                          _poolInfo.Era,
		Active:                       _poolInfo.Active,
		Bond:                         make(map[common.Address]*big.Int),
		Unbond:                       make(map[common.Address]*big.Int),
		FullfilledWithdrawalAmountOf: make(map[common.Address]*big.Int),
		Stablecoins:                  stablecoins,
	}

	for _, stablecoin := range stablecoins {
		stablecoinInfo, err := _StakeManager.GetPoolInfoByStablecoin(opts, pool, stablecoin)
		if err != nil {
			return nil, err
		}

		poolInfo.Bond[stablecoin] = stablecoinInfo.Bond
		poolInfo.Unbond[stablecoin] = stablecoinInfo.Unbond
		poolInfo.FullfilledWithdrawalAmountOf[stablecoin] = stablecoinInfo.FullfilledWithdrawalAmount
	}

	return &poolInfo, nil
}
