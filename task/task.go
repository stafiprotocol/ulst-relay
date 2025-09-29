package task

import (
	"context"
	"fmt"
	"math/big"
	"time"
	"ulst-relay/pkg/config"
	"ulst-relay/pkg/crypto/secp256k1"
	"ulst-relay/pkg/eth1"
	"ulst-relay/pkg/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Task struct {
	taskTicker  int64
	stop        chan struct{}
	rpcEndpoint string
	keyPair     *secp256k1.Keypair
	gasLimit    *big.Int
	maxGasPrice *big.Int

	stakeMangerAddress common.Address

	ethClient *eth1.Client
}

func NewTask(cfg *config.Config, keyPair *secp256k1.Keypair) (*Task, error) {
	gasLimitDeci, err := decimal.NewFromString(cfg.GasLimit)
	if err != nil {
		return nil, err
	}

	if gasLimitDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("gas limit is zero")
	}
	maxGasPriceDeci, err := decimal.NewFromString(cfg.MaxGasPrice)
	if err != nil {
		return nil, err
	}
	if maxGasPriceDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("max gas price is zero")
	}

	stakeManagerAddress := common.HexToAddress(cfg.StakeMangerAddress)

	s := &Task{
		taskTicker:         15,
		stop:               make(chan struct{}),
		rpcEndpoint:        cfg.RpcEndpoint,
		keyPair:            keyPair,
		gasLimit:           gasLimitDeci.BigInt(),
		maxGasPrice:        maxGasPriceDeci.BigInt(),
		stakeMangerAddress: stakeManagerAddress,
	}

	return s, nil
}

func (task *Task) Start() error {
	ethClient, err := eth1.NewClient(task.rpcEndpoint, task.keyPair, task.gasLimit, task.maxGasPrice)
	if err != nil {
		return err
	}
	task.ethClient = ethClient

	utils.SafeGoWithRestart(task.newEraHandler)
	utils.SafeGoWithRestart(task.newPayUnbondingFeeHandler)

	return nil
}

func (task *Task) Stop() {
	close(task.stop)
}

func (task *Task) newEraHandler() {
	logrus.Info("start new era Handler")
	ticker := time.NewTicker(time.Duration(task.taskTicker) * time.Second)
	defer ticker.Stop()

	for {

		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		case <-ticker.C:
			logrus.Debug("newEraHandler start -----------")

			err := task.handleNewEra(task.stakeMangerAddress)
			if err != nil {
				logrus.Warnf("newEraHandler failed, err: %s", err.Error())
				continue
			}

			logrus.Debug("newEraHandler end -----------")
		}
	}
}

func (task *Task) newPayUnbondingFeeHandler() {
	logrus.Info("start new pay unbonding fee Handler")
	ticker := time.NewTicker(time.Duration(task.taskTicker) * time.Second)
	defer ticker.Stop()

	for {

		select {
		case <-task.stop:
			logrus.Info("task has stopped")
			return
		case <-ticker.C:
			logrus.Debug("newPayUnbondingFeeHandler start -----------")

			err := task.handlePayUnbondingFee(context.Background(), task.stakeMangerAddress)
			if err != nil {
				logrus.Warnf("newPayUnbondingFeeHandler failed, err: %s", err.Error())
				continue
			}

			logrus.Debug("newPayUnbondingFeeHandler end -----------")
		}
	}
}

func (task *Task) waitTxOnChain(txHash common.Hash, client *eth1.Client) (err error) {
	retry := 0
	txSuccess := false
	for {
		if retry > utils.RetryLimit {
			return fmt.Errorf("waitTxOnChain %s reach retry limit", txHash.String())
		}
		_, pending, err := client.TransactionByHash(txHash)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"hash": txHash.String(),
				"err":  err.Error(),
			}).Warn("TransactionByHash")

			time.Sleep(utils.RetryInterval)
			retry++
			continue
		} else {
			if pending {
				logrus.WithFields(logrus.Fields{
					"hash":    txHash.String(),
					"pending": pending,
				}).Warn("TransactionByHash")

				time.Sleep(utils.RetryInterval)
				retry++
				continue
			} else {
				// check status
				var receipt *types.Receipt
				subRetry := 0
				for {
					if subRetry > utils.RetryLimit {
						return fmt.Errorf("TransactionReceipt %s reach retry limit", txHash.String())
					}

					receipt, err = task.ethClient.TransactionReceipt(txHash)
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"hash": txHash.String(),
							"err":  err.Error(),
						}).Warn("tx TransactionReceipt")

						time.Sleep(utils.RetryInterval)
						subRetry++
						continue
					}
					break
				}

				if receipt.Status == 1 { //success
					txSuccess = true
				}
				break
			}
		}
	}

	logrus.WithFields(logrus.Fields{
		"tx":         txHash.String(),
		"tx success": txSuccess,
	}).Info("tx already on chain")

	return nil
}
