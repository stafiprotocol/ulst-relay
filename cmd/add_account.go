package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"ulst-relay/pkg/crypto"
	"ulst-relay/pkg/crypto/secp256k1"
	"ulst-relay/pkg/keystore"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func addAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-account",
		Args:  cobra.ExactArgs(0),
		Short: "Add new account",
		RunE: func(cmd *cobra.Command, args []string) error {
			configHome, err := cmd.Flags().GetString(flagHome)
			if err != nil {
				return err
			}
			fmt.Printf("config home: %s\n", configHome)
			keystorePath := filepath.Join(configHome, "keystore")
			fmt.Printf("keystore path: %s\n", keystorePath)

			logrus.SetLevel(logrus.InfoLevel)

			return generateKeyFileByRandom(keystorePath)
		},
	}
	cmd.Flags().String(flagHome, defaultHomePath, "Home path")
	return cmd
}

func generateKeyFileByRandom(keypath string) error {
	var kp crypto.Keypair
	var err error

	kp, err = secp256k1.GenerateKeypair()
	if err != nil {
		return fmt.Errorf("could not generate secp256k1 keypair, %s", err)
	}

	fp, err := filepath.Abs(keypath + "/" + kp.Address() + ".key")
	if err != nil {
		return fmt.Errorf("invalid filepath: %s", err)
	}

	if _, err := os.Stat(fp); err != nil {
		err := os.MkdirAll(filepath.Dir(fp), 0700)
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(filepath.Clean(fp), os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			logrus.Error("add keypair: could not close keystore file")
		}
	}()

	password := keystore.GetPassword("password for key:")
	err = keystore.EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		return fmt.Errorf("could not write key to file: %s", err)
	}
	logrus.WithFields(logrus.Fields{
		"address": kp.Address(),
		"file":    fp,
	}).Info("key added")

	return nil
}
