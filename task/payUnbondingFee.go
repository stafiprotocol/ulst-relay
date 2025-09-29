package task

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"ulst-relay/bindings/stake_manager"
	"ulst-relay/bindings/stake_pool"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

func (t *Task) handlePayUnbondingFee(ctx context.Context, stakeManagerAddress common.Address) error {
	stakeManager, err := stake_manager.NewStakeManager(stakeManagerAddress, t.ethClient.Client())
	if err != nil {
		return err
	}
	bondedPools, err := stakeManager.GetBondedPools(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return err
	}
	if len(bondedPools) == 0 {
		return nil
	}

	stablecoins, err := stakeManager.GetStablecoins(nil)
	if err != nil {
		return err
	}
	for _, pool := range bondedPools {
		stakePool, err := stake_pool.NewCumstomStakePool(pool, t.ethClient.Client())
		if err != nil {
			return err
		}
		missingFees := make(map[common.Address]*big.Int)
		for _, stablecoin := range stablecoins {
			totalMissingUnbondingFee, err := stakePool.TotalMissingUnbondingFee(&bind.CallOpts{
				Context: ctx,
			}, stablecoin)
			if err != nil {
				return err
			}
			if totalMissingUnbondingFee.Cmp(big.NewInt(0)) > 0 {
				missingFees[stablecoin] = totalMissingUnbondingFee
			}
		}
		if len(missingFees) > 0 {
			err = t.invokePayUnbondingFee(ctx, stakePool, missingFees)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Task) invokePayUnbondingFee(ctx context.Context, stakePool *stake_pool.CustomStakePool, missingFees map[common.Address]*big.Int) error {
	stakecoins := make([]common.Address, 0, len(missingFees))
	amounts := make([]*big.Int, 0, len(missingFees))
	for stablecoin := range missingFees {
		stakecoins = append(stakecoins, stablecoin)
		amounts = append(amounts, missingFees[stablecoin])
		logrus.Infof("payMissingUnbondingFee %s %s to %s", stablecoin.String(), missingFees[stablecoin].String(), stakePool.Address())
	}

	if err := t.ethClient.LockAndUpdateOpts(t.gasLimit, big.NewInt(0)); err != nil {
		return err
	}
	tx, err := stakePool.PayMissingUnbondingFee(t.ethClient.Opts(), stakecoins, amounts)
	if err != nil {
		return err
	}
	err = t.waitTxOnChain(tx.Hash(), t.ethClient)
	if err != nil {
		return err
	}

	//wait until payMissingUnbondingFee executed
	retry := 0
Outter:
	for {
		// wait 2h
		if retry > 2*60*5 {
			return fmt.Errorf("wait payMissingUnbondingFee executed failed")
		}
		for _, stablecoin := range stakecoins {
			totalMissingUnbondingFee, err := stakePool.TotalMissingUnbondingFee(nil, stablecoin)
			if err != nil {
				return err
			}
			if totalMissingUnbondingFee.Cmp(big.NewInt(0)) > 0 {
				logrus.Warnf("waiting payMissingUnbondingFee executed...")
				time.Sleep(12 * time.Second)
				retry++
				continue Outter
			}
		}

		break
	}

	return nil
}
