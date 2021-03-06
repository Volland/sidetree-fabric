/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"strings"

	"github.com/hyperledger/fabric/peer/node"
	"github.com/spf13/cobra"
	viper "github.com/spf13/viper2015"

	sidetreepeer "github.com/trustbloc/sidetree-fabric/pkg/peer"
)

func main() {
	setup()

	sidetreepeer.Initialize()

	if err := startPeer(); err != nil {
		panic(err)
	}
}

func setup() {
	replacer := strings.NewReplacer(".", "_")

	viper.SetEnvPrefix(node.CmdRoot)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(replacer)

	node.InitCmd(&cobra.Command{}, nil)
}

func startPeer() error {
	return node.Start()
}
