/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockchainhandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	cb "github.com/hyperledger/fabric-protos-go/common"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/trustbloc/fabric-peer-ext/pkg/mocks"

	"github.com/trustbloc/sidetree-fabric/pkg/httpserver"
	"github.com/trustbloc/sidetree-fabric/pkg/observer/common"
	obmocks "github.com/trustbloc/sidetree-fabric/pkg/observer/mocks"
)

func TestNewTransactionsHandler(t *testing.T) {
	bcProvider := &obmocks.BlockchainClientProvider{}

	h := NewTransactionsHandler(channel1, handlerCfg, bcProvider)
	require.NotNil(t, h)

	require.Equal(t, "/blockchain/transactions", h.Path())
	require.Equal(t, http.MethodGet, h.Method())
}

func TestNewTransactionsSinceHandler(t *testing.T) {
	bcProvider := &obmocks.BlockchainClientProvider{}

	h := NewTransactionsSinceHandler(channel1, handlerCfg, bcProvider)
	require.NotNil(t, h)

	params := map[string]string{"since": "{since}", "transaction-time-hash": "{transaction-time-hash}"}

	require.Equal(t, "/blockchain/transactions", h.Path())
	require.Equal(t, http.MethodGet, h.Method())
	require.Equal(t, params, h.Params())
}

func TestTransactions_All(t *testing.T) {
	const blockNum = 1000
	const txn1 = "tx1"
	const anchor = "xxx"

	bcInfo := &cb.BlockchainInfo{
		Height:           1000,
		CurrentBlockHash: []byte{1, 2, 3, 4},
	}

	bcProvider := &obmocks.BlockchainClientProvider{}

	t.Run("Success", func(t *testing.T) {
		bcClient := &obmocks.BlockchainClient{}
		bcClient.GetBlockchainInfoReturns(bcInfo, nil)

		bb := mocks.NewBlockBuilder(channel1, blockNum)
		bb.Transaction(txn1, pb.TxValidationCode_VALID).ChaincodeAction("sidetree").Write(common.AnchorPrefix, getTxnInfo(anchor))
		bcClient.GetBlockByNumberReturns(bb.Build(), nil)

		bcProvider.ForChannelReturns(bcClient, nil)

		h := NewTransactionsHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusOK, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeJSON, rw.Header().Get(httpserver.ContentTypeHeader))

		resp := &TransactionsResponse{}
		require.NoError(t, json.Unmarshal(rw.Body.Bytes(), resp))
	})

	t.Run("Marshal error -> error", func(t *testing.T) {
		bcClient := &obmocks.BlockchainClient{}
		bcClient.GetBlockchainInfoReturns(bcInfo, nil)

		bb := mocks.NewBlockBuilder(channel1, blockNum)
		bb.Transaction(txn1, pb.TxValidationCode_VALID).ChaincodeAction("sidetree").Write(common.AnchorPrefix, getTxnInfo(anchor))
		bcClient.GetBlockByNumberReturns(bb.Build(), nil)

		bcProvider.ForChannelReturns(bcClient, nil)

		h := NewTransactionsHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		errExpected := errors.New("injected marshal error")
		h.jsonMarshal = func(v interface{}) ([]byte, error) { return nil, errExpected }

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusInternalServerError, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeText, rw.Header().Get(httpserver.ContentTypeHeader))
		require.Equal(t, httpserver.StatusServerError, rw.Body.String())
	})

	t.Run("Maximum reached -> Success", func(t *testing.T) {
		bcClient := &obmocks.BlockchainClient{}
		bcClient.GetBlockchainInfoReturns(bcInfo, nil)

		bb := mocks.NewBlockBuilder(channel1, blockNum)
		bb.Transaction(txn1, pb.TxValidationCode_VALID).ChaincodeAction("sidetree").Write(common.AnchorPrefix, getTxnInfo(anchor))
		bcClient.GetBlockByNumberReturns(bb.Build(), nil)

		bcProvider.ForChannelReturns(bcClient, nil)

		cfg := Config{
			BasePath:                  "/blockchain",
			MaxTransactionsInResponse: 0,
		}

		h := NewTransactionsHandler(channel1, cfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusOK, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeJSON, rw.Header().Get(httpserver.ContentTypeHeader))

		resp := &TransactionsResponse{}
		require.NoError(t, json.Unmarshal(rw.Body.Bytes(), resp))
	})

	t.Run("Blockchain client provider error -> error", func(t *testing.T) {
		errExpected := errors.New("injected blockchain client provider error")
		bcProvider.ForChannelReturns(nil, errExpected)

		h := NewTransactionsHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusInternalServerError, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeText, rw.Header().Get(httpserver.ContentTypeHeader))
	})

	t.Run("Blockchain client error -> error", func(t *testing.T) {
		errExpected := errors.New("injected blockchain client error")
		bcClient := &obmocks.BlockchainClient{}
		bcClient.GetBlockchainInfoReturns(nil, errExpected)
		bcProvider.ForChannelReturns(bcClient, nil)

		h := NewTransactionsHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusInternalServerError, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeText, rw.Header().Get(httpserver.ContentTypeHeader))
	})
}

func TestTransactions_Since(t *testing.T) {
	const blockNum = 1000
	const txn1 = "tx1"
	const anchor = "xxx"

	bcInfo := &cb.BlockchainInfo{
		Height:           1002,
		CurrentBlockHash: []byte{1, 2, 3, 4},
	}

	bcProvider := &obmocks.BlockchainClientProvider{}

	t.Run("Success", func(t *testing.T) {
		restoreParams := setBlockchainParams("Zi6XcyzoikY-OPOme_l2zQpexGdLov1-23ciPN66QQ8=", "0")
		defer restoreParams()

		bb := mocks.NewBlockBuilder(channel1, blockNum)
		bb.Transaction(txn1, pb.TxValidationCode_VALID).ChaincodeAction("sidetree").Write(common.AnchorPrefix, getTxnInfo(anchor))

		bcClient := &obmocks.BlockchainClient{}
		bcClient.GetBlockchainInfoReturns(bcInfo, nil)
		bcClient.GetBlockByHashReturns(bb.Build(), nil)
		bcClient.GetBlockByNumberReturns(bb.Build(), nil)

		bcProvider.ForChannelReturns(bcClient, nil)

		h := NewTransactionsSinceHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusOK, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeJSON, rw.Header().Get(httpserver.ContentTypeHeader))

		resp := &TransactionsResponse{}
		require.NoError(t, json.Unmarshal(rw.Body.Bytes(), resp))
	})

	t.Run("Block not found -> error", func(t *testing.T) {
		restoreParams := setBlockchainParams("Zi6XcyzoikY-OPOme_l2zQpexGdLov1-23ciPN66QQ8=", "0")
		defer restoreParams()

		bcClient := &obmocks.BlockchainClient{}
		bcClient.GetBlockchainInfoReturns(bcInfo, nil)
		bcClient.GetBlockByHashReturns(nil, errors.New("not found"))

		bcProvider.ForChannelReturns(bcClient, nil)

		h := NewTransactionsSinceHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusNotFound, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeText, rw.Header().Get(httpserver.ContentTypeHeader))
		require.Equal(t, httpserver.StatusNotFound, rw.Body.String())
	})

	t.Run("Invalid time hash -> error", func(t *testing.T) {
		restoreParams := setBlockchainParams("xxx_xxx", "0")
		defer restoreParams()

		h := NewTransactionsSinceHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusBadRequest, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeJSON, rw.Header().Get(httpserver.ContentTypeHeader))

		errResp := &ErrorResponse{}
		require.NoError(t, json.Unmarshal(rw.Body.Bytes(), errResp))
		require.Equal(t, InvalidTxNumOrTimeHash, errResp.Code)
	})

	t.Run("Invalid since param -> error", func(t *testing.T) {
		restoreParams := setBlockchainParams("Zi6XcyzoikY-OPOme_l2zQpexGdLov1-23ciPN66QQ8=", "xxx")
		defer restoreParams()

		h := NewTransactionsSinceHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusBadRequest, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeJSON, rw.Header().Get(httpserver.ContentTypeHeader))

		errResp := &ErrorResponse{}
		require.NoError(t, json.Unmarshal(rw.Body.Bytes(), errResp))
		require.Equal(t, InvalidTxNumOrTimeHash, errResp.Code)
	})

	t.Run("Blockchain client provider error -> error", func(t *testing.T) {
		restoreParams := setBlockchainParams("Zi6XcyzoikY-OPOme_l2zQpexGdLov1-23ciPN66QQ8=", "0")
		defer restoreParams()

		errExpected := errors.New("injected blockchain client provider error")
		bcProvider.ForChannelReturns(nil, errExpected)

		h := NewTransactionsSinceHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusInternalServerError, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeText, rw.Header().Get(httpserver.ContentTypeHeader))
	})

	t.Run("Blockchain client error -> error", func(t *testing.T) {
		restoreParams := setBlockchainParams("Zi6XcyzoikY-OPOme_l2zQpexGdLov1-23ciPN66QQ8=", "0")
		defer restoreParams()

		errExpected := errors.New("injected blockchain client error")
		bcClient := &obmocks.BlockchainClient{}
		bcClient.GetBlockByHashReturns(nil, errExpected)
		bcProvider.ForChannelReturns(bcClient, nil)

		h := NewTransactionsSinceHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusInternalServerError, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeText, rw.Header().Get(httpserver.ContentTypeHeader))
	})

	t.Run("GetBlockByNumber error -> error", func(t *testing.T) {
		restoreParams := setBlockchainParams("Zi6XcyzoikY-OPOme_l2zQpexGdLov1-23ciPN66QQ8=", "1000")
		defer restoreParams()

		bb := mocks.NewBlockBuilder(channel1, blockNum)
		bb.Transaction(txn1, pb.TxValidationCode_VALID).ChaincodeAction("sidetree").Write(common.AnchorPrefix, getTxnInfo(anchor))

		errExpected := errors.New("injected blockchain client error")
		bcClient := &obmocks.BlockchainClient{}
		bcClient.GetBlockchainInfoReturns(bcInfo, nil)
		bcClient.GetBlockByHashReturns(bb.Build(), nil)
		bcClient.GetBlockByNumberReturns(nil, errExpected)
		bcProvider.ForChannelReturns(bcClient, nil)

		h := NewTransactionsSinceHandler(channel1, handlerCfg, bcProvider)
		require.NotNil(t, h)

		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/blockchain/transactions", nil)

		h.Handler()(rw, req)

		require.Equal(t, http.StatusInternalServerError, rw.Result().StatusCode)
		require.Equal(t, httpserver.ContentTypeText, rw.Header().Get(httpserver.ContentTypeHeader))
	})
}

func setBlockchainParams(timeHash, since string) func() {
	restoreGetTimeHash := getTimeHash
	restoreGetSince := getSince

	getTimeHash = func(req *http.Request) string {
		return timeHash
	}
	getSince = func(req *http.Request) string {
		return since
	}

	return func() {
		getTimeHash = restoreGetTimeHash
		getSince = restoreGetSince
	}
}
