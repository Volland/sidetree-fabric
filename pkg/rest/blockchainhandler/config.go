/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockchainhandler

import "github.com/trustbloc/sidetree-fabric/pkg/rest/authhandler"

// Config defines the configuration for a blockchain handler
type Config struct {
	Authorization authhandler.Config

	// BasePath is the base context path of the REST endpoint
	BasePath string
	// MaxTransactionsInResponse is the maximum number of transactions to return for the /blockchain/transactions request
	MaxTransactionsInResponse int
	// MaxBlocksInResponse is the maximum number of blocks to return for the /blockchain/blocks request
	MaxBlocksInResponse int
}
