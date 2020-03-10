/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package filehandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/trustbloc/sidetree-core-go/pkg/document"

	"github.com/trustbloc/sidetree-fabric/pkg/mocks"
)

//go:generate counterfeiter -o ../../mocks/documentresolver.gen.go --fake-name DocumentResolver . documentResolver

const (
	channelID = "channel1"
	schema1   = "schema1.json"
)

func TestFileRetrieveHandler(t *testing.T) {
	docResolver := &mocks.DocumentResolver{}
	dcasProvider := &mocks.DCASClientProvider{}

	dcasClient := &mocks.DCASClient{}
	dcasProvider.ForChannelReturns(dcasClient, nil)

	cfg := Config{
		BasePath:       "/schema",
		ChaincodeName:  "file_cc",
		Collection:     "schemas",
		IndexNamespace: "file:idx",
		IndexDocID:     "file:idx:1234",
	}

	h := NewRetrieveHandler(channelID, cfg, docResolver, dcasProvider)
	require.Equal(t, cfg.BasePath+"/{resourceName}", h.Path())
	require.Equal(t, http.MethodGet, h.Method())
	require.NotNil(t, h.Handler())

	t.Run("Bad request", func(t *testing.T) {
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/schema", nil)
		h.Handler()(rw, req)
		require.Equal(t, http.StatusBadRequest, rw.Code)
		require.Contains(t, rw.Body.String(), "resource name not provided")
	})

	t.Run("File index not found", func(t *testing.T) {
		docResolver.ResolveDocumentReturns(nil, errors.New("not found"))

		getResourceName = func(req *http.Request) string { return schema1 }
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
		h.Handler()(rw, req)
		require.Equal(t, http.StatusNotFound, rw.Code)
		require.Equal(t, "file index document not found", rw.Body.String())
	})

	t.Run("File index was deleted", func(t *testing.T) {
		docResolver.ResolveDocumentReturns(nil, errors.New("was deleted"))

		getResourceName = func(req *http.Request) string { return schema1 }
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
		h.Handler()(rw, req)
		require.Equal(t, http.StatusGone, rw.Code)
		require.Equal(t, "document is no longer available", rw.Body.String())
	})

	t.Run("File not found in index", func(t *testing.T) {
		docResolver.ResolveDocumentReturns(make(document.Document), nil)

		getResourceName = func(req *http.Request) string { return schema1 }
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
		h.Handler()(rw, req)
		require.Equal(t, http.StatusNotFound, rw.Code)
		require.Equal(t, fileNotFound, rw.Body.String())
	})

	t.Run("File not found in DCAS", func(t *testing.T) {
		doc := make(document.Document)
		doc[schema1] = "1234567890"
		docResolver.ResolveDocumentReturns(doc, nil)
		dcasClient.GetReturns(nil, nil)

		getResourceName = func(req *http.Request) string { return schema1 }
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
		h.Handler()(rw, req)
		require.Equal(t, http.StatusNotFound, rw.Code)
		require.Equal(t, fileNotFound, rw.Body.String())
	})

	t.Run("DCAS provider error", func(t *testing.T) {
		dcasProvider.ForChannelReturns(nil, errors.New("injected DCAS provider error"))
		defer func() { dcasProvider.ForChannelReturns(dcasClient, nil) }()

		doc := make(document.Document)
		doc[schema1] = "1234567890"
		docResolver.ResolveDocumentReturns(doc, nil)
		dcasClient.GetReturns(nil, errors.New("injected DCAS error"))

		getResourceName = func(req *http.Request) string { return schema1 }
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
		h.Handler()(rw, req)
		require.Equal(t, http.StatusInternalServerError, rw.Code)
		require.Equal(t, serverError, rw.Body.String())
	})

	t.Run("DCAS error", func(t *testing.T) {
		doc := make(document.Document)
		doc[schema1] = "1234567890"
		docResolver.ResolveDocumentReturns(doc, nil)
		dcasClient.GetReturns(nil, errors.New("injected DCAS error"))

		getResourceName = func(req *http.Request) string { return schema1 }
		rw := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
		h.Handler()(rw, req)
		require.Equal(t, http.StatusInternalServerError, rw.Code)
		require.Equal(t, serverError, rw.Body.String())
	})

	t.Run("File retrieved from DCAS", func(t *testing.T) {
		doc := make(document.Document)
		doc[schema1] = "1234567890"
		docResolver.ResolveDocumentReturns(doc, nil)
		getResourceName = func(req *http.Request) string { return schema1 }

		t.Run("Invalid file", func(t *testing.T) {
			dcasClient.GetReturns([]byte("{"), nil)
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
			h.Handler()(rw, req)
			require.Equal(t, http.StatusInternalServerError, rw.Code)
			require.Contains(t, rw.Body.String(), serverError)
		})

		t.Run("Missing content-type", func(t *testing.T) {
			dcasClient.GetReturns([]byte("{}"), nil)
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
			h.Handler()(rw, req)
			require.Equal(t, http.StatusInternalServerError, rw.Code)
			require.Contains(t, rw.Body.String(), serverError)
		})

		t.Run("Empty content", func(t *testing.T) {
			f := &File{
				ContentType: "application/json",
			}

			fileBytes, err := json.Marshal(f)
			require.NoError(t, err)

			dcasClient.GetReturns(fileBytes, nil)
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
			h.Handler()(rw, req)
			require.Equal(t, http.StatusNotFound, rw.Code)
		})

		t.Run("Success", func(t *testing.T) {
			fileContents := `{"field1":"value1"}`
			f := &File{
				ContentType: "application/json",
				Content:     []byte(fileContents),
			}

			fileBytes, err := json.Marshal(f)
			require.NoError(t, err)

			dcasClient.GetReturns(fileBytes, nil)
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/schema/schema1.json", nil)
			h.Handler()(rw, req)
			require.Equal(t, http.StatusOK, rw.Code)
			require.Equal(t, fileContents, rw.Body.String())
		})
	})
}
