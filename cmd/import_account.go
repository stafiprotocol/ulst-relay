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

func importAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import-account",
		Args:  cobra.ExactArgs(0),
		Short: "Import account",
		RunE: func(cmd *cobra.Command, args []string) error {
			configHome, err := cmd.Flags().GetString(flagHome)
			if err != nil {
				return err
			}
			fmt.Printf("config home: %s\n", configHome)
			keystorePath := filepath.Join(configHome, "keystore")
			fmt.Printf("keystore path: %s\n", keystorePath)

			logrus.SetLevel(logrus.InfoLevel)

			return generateKeyFileByPrivateKey(keystorePath)
		},
	}
	cmd.Flags().String(flagHome, defaultHomePath, "Home path")
	return cmd
}

func generateKeyFileByPrivateKey(keypath string) error {
	var kp crypto.Keypair
	var err error

	key := keystore.GetPassword("Enter private key:")
	skey := string(key)

	if skey[0:2] == "0x" {
		kp, err = secp256k1.NewKeypairFromString(skey[2:])
	} else {
		kp, err = secp256k1.NewKeypairFromString(skey)
	}
	if err != nil {
		return fmt.Errorf("could not import secp256k1 keypair from given string: %s", err)
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
			logrus.Error("import keypair: could not close keystore file")
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
	}).Info("key imported")

	return nil
}
