/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package dcashandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/trustbloc/sidetree-fabric/pkg/httpserver"
	"github.com/trustbloc/sidetree-fabric/pkg/rest/versionhandler"
)

func TestNewVersionHandler(t *testing.T) {
	h := NewVersionHandler(channel1, handlerCfg)
	require.NotNil(t, h)

	require.Equal(t, "/cas/version", h.Path())
	require.Equal(t, http.MethodGet, h.Method())
}

func TestVersion_Handler(t *testing.T) {
	const v1 = "1.0.1"

	cfg := Config{
		BasePath:      "/cas",
		ChaincodeName: cc1,
		Collection:    coll1,
		Version:       v1,
	}

	h := NewVersionHandler(channel1, cfg)
	require.NotNil(t, h)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/cas/version", nil)

	h.Handler()(rw, req)

	require.Equal(t, http.StatusOK, rw.Result().StatusCode)
	require.Equal(t, httpserver.ContentTypeJSON, rw.Header().Get(httpserver.ContentTypeHeader))

	resp := &versionhandler.Response{}
	require.NoError(t, json.Unmarshal(rw.Body.Bytes(), resp))
	require.Equal(t, moduleName, resp.Name)
	require.Equal(t, v1, resp.Version)
}
