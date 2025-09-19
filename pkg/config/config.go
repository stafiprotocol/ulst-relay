// Copyright 2025 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package config

type Config struct {
	StakeMangerAddress string
	RpcEndpoint        string
	Account            string
	GasLimit           string
	MaxGasPrice        string

	LogFilePath  string
	KeystorePath string
}
