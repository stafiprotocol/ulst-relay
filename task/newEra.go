package task

import (
	"context"
	"fmt"
	"math/big"
	"time"

	stake_manager "ulst-relay/bindings/StakeManager"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (t *Task) handleNewEra(stakeManagerAddress common.Address) error {
	stakeManger, err := stake_manager.NewStakeManager(stakeManagerAddress, t.ethClient.Client())
	if err != nil {
		return err
	}
	bondedPools, err := stakeManger.GetBondedPools(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		return err
	}
	if len(bondedPools) == 0 {
		return nil
	}

	latestCallOpts := bind.CallOpts{
		Pending: false,
		From:    [20]byte{},
		Context: context.Background(),
	}

	currentEra, err := stakeManger.CurrentEraUint64(&latestCallOpts)
	if err != nil {
		return err
	}
	latestEra, err := stakeManger.LatestEraUint64(&latestCallOpts)
	if err != nil {
		return err
	}

	if currentEra <= latestEra {
		logrus.Debugf("skipped: already updated to latest era, currentEra(%d) <= latestEra(%d)", currentEra, latestEra)
		return nil
	}

	err = t.invokeNewEra(stakeManger, currentEra, latestEra, &latestCallOpts, bondedPools)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) invokeNewEra(stakeManager *stake_manager.StakeManager, currentEra, latestEra uint64, latestCallOpts *bind.CallOpts, bondedPools []common.Address) error {
	oldRate, err := stakeManager.GetRate(nil)
	if err != nil {
		return err
	}

	oldPoolInfoStr := ""
	for _, pool := range bondedPools {
		poolInfo, err := stakeManager.GetPoolInfoOf(nil, pool)
		if err != nil {
			return err
		}
		oldPoolInfoStr += poolInfo.String() + "\n"
	}

	// send tx
	err = t.ethClient.LockAndUpdateOpts(t.gasLimit, big.NewInt(0))
	if err != nil {
		return err
	}
	tx, err := stakeManager.NewEra(t.ethClient.Opts())
	t.ethClient.UnlockOpts()
	if err != nil {
		return err
	}

	err = t.waitTxOnChain(tx.Hash(), t.ethClient)
	if err != nil {
		return errors.Wrap(err, "waitTxOnChain failed")
	}

	willUseEra := latestEra + 1
	//wait until newEra executed
	retry := 0
	for {
		// wait 2h
		if retry > 2*60*5 {
			return fmt.Errorf("wait newEra %d executed failed", willUseEra)
		}
		latestEra, err := stakeManager.LatestEraUint64(latestCallOpts)
		if err != nil {
			logrus.Warnf("get latestEra failed: %s", err.Error())
			time.Sleep(12 * time.Second)
			retry++
			continue
		}

		if latestEra < willUseEra {
			logrus.Warnf("waiting newEra %d executed...", willUseEra)
			time.Sleep(12 * time.Second)
			retry++
			continue
		}
		break
	}

	newPoolInfoStr := ""
	for _, pool := range bondedPools {
		poolInfo, err := stakeManager.GetPoolInfoOf(nil, pool)
		if err != nil {
			return err
		}
		newPoolInfoStr += poolInfo.String() + "\n"
	}

	newRate, err := stakeManager.GetRate(nil)
	if err != nil {
		return err
	}

	logrus.Infof(`newEra %d executed successfully, new rate: %s, old rate: %s
new pool info:
%s
old pool info:
%s`, willUseEra, newRate, oldRate, newPoolInfoStr, oldPoolInfoStr)

	return nil
}
