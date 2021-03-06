/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sidetreesvc

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	extmocks "github.com/trustbloc/fabric-peer-ext/pkg/mocks"
	protocolApi "github.com/trustbloc/sidetree-core-go/pkg/api/protocol"

	"github.com/trustbloc/sidetree-fabric/pkg/config"
	cfgmocks "github.com/trustbloc/sidetree-fabric/pkg/config/mocks"
	sidetreectx "github.com/trustbloc/sidetree-fabric/pkg/context"
	ctxmocks "github.com/trustbloc/sidetree-fabric/pkg/context/mocks"
	"github.com/trustbloc/sidetree-fabric/pkg/mocks"
	peermocks "github.com/trustbloc/sidetree-fabric/pkg/peer/mocks"
	"github.com/trustbloc/sidetree-fabric/pkg/rest/sidetreehandler"
)

//go:generate counterfeiter -o ../mocks/protocolversionfactory.gen.go --fake-name ProtocolVersionFactory . protocolVersionFactory

func TestContext(t *testing.T) {
	nsCfg := sidetreehandler.Config{
		Namespace: didTrustblocNamespace,
		BasePath:  didTrustblocBasePath,
	}

	dcasCfg := config.DCAS{
		ChaincodeName: "cc1",
		Collection:    "dcas",
	}

	restCfg := &peermocks.RestConfig{}

	ctxProviders := &ContextProviders{
		Providers: &sidetreectx.Providers{
			TxnProvider:                &peermocks.TxnServiceProvider{},
			DCASProvider:               &mocks.DCASClientProvider{},
			OperationQueueProvider:     &mocks.OperationQueueProvider{},
			LedgerProvider:             &extmocks.LedgerProvider{},
			OperationProcessorProvider: &ctxmocks.CachingOpProcessorProvider{},
		},
		VersionFactory: &peermocks.ProtocolVersionFactory{},
	}

	cacheProvider := &ctxmocks.CachingOpProcessorProvider{}

	t.Run("Success", func(t *testing.T) {
		protocolVersions := map[string]protocolApi.Protocol{
			"0.5": {
				GenesisTime:         100,
				MultihashAlgorithms: []uint{18},
				MaxOperationCount:   100,
				MaxOperationSize:    1000,
			},
		}

		stConfigService := &cfgmocks.SidetreeConfigService{}
		stConfigService.LoadProtocolsReturns(protocolVersions, nil)

		ctx, err := newContext(channel1, nsCfg, dcasCfg, stConfigService, ctxProviders, &mocks.OperationStoreProvider{}, restCfg, cacheProvider)
		require.NoError(t, err)
		require.NotNil(t, ctx)

		require.NotNil(t, ctx.BatchWriter())

		require.NoError(t, ctx.Start())

		time.Sleep(20 * time.Millisecond)

		ctx.Stop()
	})

	t.Run("Operation store error", func(t *testing.T) {
		errExpected := errors.New("injected operation store error")

		protocolVersions := map[string]protocolApi.Protocol{
			"0.5": {
				GenesisTime:         100,
				MultihashAlgorithms: []uint{18},
				MaxOperationCount:   100,
				MaxOperationSize:    1000,
			},
		}

		stConfigService := &cfgmocks.SidetreeConfigService{}
		stConfigService.LoadProtocolsReturns(protocolVersions, nil)

		opStoreProvider := &mocks.OperationStoreProvider{}
		opStoreProvider.ForNamespaceReturns(nil, errExpected)

		ctx, err := newContext(channel1, nsCfg, dcasCfg, stConfigService, ctxProviders, opStoreProvider, restCfg, cacheProvider)
		require.EqualError(t, err, errExpected.Error())
		require.Nil(t, ctx)
	})

	t.Run("Sidetree config service error", func(t *testing.T) {
		errExpected := errors.New("injected sidetree config service error")

		protocolVersions := map[string]protocolApi.Protocol{
			"0.5": {
				GenesisTime:         100,
				MultihashAlgorithms: []uint{18},
				MaxOperationCount:   100,
				MaxOperationSize:    1000,
			},
		}

		stConfigService := &cfgmocks.SidetreeConfigService{}
		stConfigService.LoadProtocolsReturns(protocolVersions, nil)
		stConfigService.LoadSidetreeReturns(config.Sidetree{}, errExpected)

		ctx, err := newContext(channel1, nsCfg, dcasCfg, stConfigService, ctxProviders, &mocks.OperationStoreProvider{}, restCfg, cacheProvider)
		require.EqualError(t, err, errExpected.Error())
		require.Nil(t, ctx)
	})

	t.Run("No protocols -> error", func(t *testing.T) {
		stConfigService := &cfgmocks.SidetreeConfigService{}

		ctx, err := newContext(channel1, nsCfg, dcasCfg, stConfigService, ctxProviders, &mocks.OperationStoreProvider{}, restCfg, cacheProvider)
		require.Error(t, err)
		require.Contains(t, err.Error(), "no protocols defined")
		require.Nil(t, ctx)
	})

	t.Run("Initialize protocols -> error", func(t *testing.T) {
		errExpected := errors.New("injected sidetreeCfgService error")
		stConfigService := &cfgmocks.SidetreeConfigService{}
		stConfigService.LoadProtocolsReturns(nil, errExpected)

		ctx, err := newContext(channel1, nsCfg, dcasCfg, stConfigService, ctxProviders, &mocks.OperationStoreProvider{}, restCfg, cacheProvider)
		require.EqualError(t, err, errExpected.Error())
		require.Nil(t, ctx)
	})
}
