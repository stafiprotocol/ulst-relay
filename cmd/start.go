package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"ulst-relay/pkg/config"
	"ulst-relay/pkg/crypto/secp256k1"
	"ulst-relay/pkg/keystore"
	"ulst-relay/pkg/log"
	"ulst-relay/pkg/utils"
	"ulst-relay/task"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	flagHome         = "home"
	flagEndpoint     = "endpoint"
	flagAccount      = "account"
	flagGasLimit     = "gas_limit"
	flagMaxGasPrice  = "max_gas_price"
	flagStakeManager = "stake_manager"
	flagLogLevel     = "log_level"

	defaultHomePath    = filepath.Join(os.Getenv("HOME"), ".stafi/ulst")
	defaultEndpoint    = "http://localhost:8545"
	defaultGasLimit    = "2000000"
	defaultMaxGasPrice = "600000000000"
	defaultLogLevel    = logrus.InfoLevel.String()
)

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Args:  cobra.ExactArgs(0),
		Short: "Start stablecoin lsd relay",
		RunE: func(cmd *cobra.Command, args []string) error {
			configHome, err := cmd.Flags().GetString(flagHome)
			if err != nil {
				return err
			}
			fmt.Printf("config home: %s\n", configHome)

			configEndpoint, err := cmd.Flags().GetString(flagEndpoint)
			if err != nil {
				return err
			}
			configAccount, err := cmd.Flags().GetString(flagAccount)
			if err != nil {
				return err
			}
			if !common.IsHexAddress(configAccount) {
				return fmt.Errorf("account not hex address: %s", configAccount)
			}

			configGasLimit, err := cmd.Flags().GetString(flagGasLimit)
			if err != nil {
				return err
			}

			configMaxGasPrice, err := cmd.Flags().GetString(flagMaxGasPrice)
			if err != nil {
				return err
			}
			configStakeManager, err := cmd.Flags().GetString(flagStakeManager)
			if err != nil {
				return err
			}
			if len(configStakeManager) > 0 && !common.IsHexAddress(configStakeManager) {
				return fmt.Errorf("stake manager not hex address: %s", configStakeManager)
			}
			if len(configStakeManager) == 0 {
				return errors.New("stake manager must be provided")
			}

			// check log level
			logLevelStr, err := cmd.Flags().GetString(flagLogLevel)
			if err != nil {
				return err
			}
			logLevel, err := logrus.ParseLevel(logLevelStr)
			if err != nil {
				return err
			}
			fmt.Printf("log level: %s\n", logLevelStr)
			logrus.SetLevel(logLevel)

			logFilePath := filepath.Join(configHome, "log_data")
			keystorePath := filepath.Join(configHome, "keystore")

			cfg := config.Config{}
			cfg.RpcEndpoint = configEndpoint
			cfg.Account = configAccount
			cfg.GasLimit = configGasLimit
			cfg.MaxGasPrice = configMaxGasPrice
			cfg.StakeMangerAddress = configStakeManager

			cfg.LogFilePath = logFilePath
			cfg.KeystorePath = keystorePath

			err = log.InitLogFile(cfg.LogFilePath)
			if err != nil {
				return err
			}
			logrus.Infof("cfg %+v", cfg)

			ctx := utils.ShutdownListener()
			kpI, err := keystore.KeypairFromAddress(cfg.Account, keystore.EthChain, cfg.KeystorePath, false)
			if err != nil {
				return err
			}
			kp, ok := kpI.(*secp256k1.Keypair)
			if !ok {
				return fmt.Errorf("keypair err")
			}

			logrus.Info("task starting...")
			t, err := task.NewTask(&cfg, kp)
			if err != nil {
				return err
			}
			err = t.Start()
			if err != nil {
				logrus.Errorf("task start err: %s", err)
				return err
			}
			defer func() {
				logrus.Infof("shutting down task ...")
				t.Stop()
			}()

			<-ctx.Done()
			return nil

		},
	}

	cmd.Flags().String(flagHome, defaultHomePath, "Home path")
	cmd.Flags().String(flagEndpoint, defaultEndpoint, "Rpc endpoint")
	cmd.Flags().String(flagGasLimit, defaultGasLimit, "Gas limit")
	cmd.Flags().String(flagMaxGasPrice, defaultMaxGasPrice, "Max gas price")
	cmd.Flags().String(flagLogLevel, defaultLogLevel, "The logging level (trace|debug|info|warn|error|fatal|panic)")

	// Required flags
	cmd.Flags().String(flagAccount, "", "Account hex string address")
	cmd.Flags().String(flagStakeManager, "", "Stake manager contract address")
	_ = cmd.MarkFlagRequired(flagAccount)
	_ = cmd.MarkFlagRequired(flagStakeManager)

	return cmd
}
