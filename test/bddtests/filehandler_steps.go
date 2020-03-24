/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bddtests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cucumber/godog"
	jsonpatch "github.com/evanphx/json-patch"
	"github.com/pkg/errors"

	"github.com/trustbloc/fabric-peer-test-common/bddtests"

	"github.com/trustbloc/sidetree-core-go/pkg/document"
	"github.com/trustbloc/sidetree-core-go/pkg/docutil"
	"github.com/trustbloc/sidetree-core-go/pkg/restapi/helper"

	"github.com/trustbloc/sidetree-fabric/test/bddtests/restclient"
)

// FileHandlerSteps
type FileHandlerSteps struct {
	encodedCreatePayload string
	reqNamespace         string
	resp                 *restclient.HttpResponse
	bddContext           *bddtests.BDDContext
}

// NewFileHandlerSteps
func NewFileHandlerSteps(context *bddtests.BDDContext) *FileHandlerSteps {
	return &FileHandlerSteps{bddContext: context}
}

func (d *FileHandlerSteps) createDocument(url, content, namespace string) error {
	resolved, err := bddtests.ResolveAllVars(content)
	if err != nil {
		return err
	}

	if len(resolved) != 1 {
		return errors.Errorf("expecting 1 var but got %d", len(resolved))
	}

	content = resolved[0]

	logger.Infof("Creating document at [%s] in namespace [%s] with content %s", url, namespace, content)

	req, err := getCreateRequest(d.getOpaqueDocument(content))
	if err != nil {
		return err
	}

	d.encodedCreatePayload = docutil.EncodeToString(req)
	d.reqNamespace = namespace

	logger.Infof("Sending document to [%s]: %s", url, req)
	d.resp, err = restclient.SendRequest(url, req)

	logger.Infof("... got response from [%s]: %s", url, d.resp.Payload)

	return err
}

func (d *FileHandlerSteps) updateDocument(url, docID, jsonPatch string) error {
	logger.Infof("Updating document [%s] at [%s] with patch %s", docID, url, jsonPatch)

	resolvedPatch, err := bddtests.ResolveAllVars(jsonPatch)
	if err != nil {
		return err
	}

	if len(resolvedPatch) != 1 {
		return errors.Errorf("expecting 1 var but got %d", len(resolvedPatch))
	}

	jsonPatch = resolvedPatch[0]

	resolvedDocID, err := bddtests.ResolveAllVars(docID)
	if err != nil {
		return err
	}

	if len(resolvedDocID) != 1 {
		return errors.Errorf("expecting 1 var but got %d", len(resolvedDocID))
	}

	uniqueSuffix := getUniqueSuffix(resolvedDocID[0])

	patch, err := jsonpatch.DecodePatch([]byte(jsonPatch))
	if err != nil {
		return err
	}

	logger.Infof("Updating document [%s] at [%s] with patch %s - %+v", docID, url, jsonPatch, patch)

	req, err := d.getUpdateRequest(uniqueSuffix, patch)
	if err != nil {
		return err
	}

	logger.Infof("Sending update payload to [%s]: %s", url, req)

	d.resp, err = restclient.SendRequest(url, []byte(req))

	logger.Infof("... got response from [%s] - Status code: %d, Payload: %s", url, d.resp.StatusCode, d.resp.Payload)

	return err
}

func (d *FileHandlerSteps) uploadFile(url, path, contentType string) error {
	logger.Infof("Uploading file [%s] to [%s]", path, url)

	fileBytes := getFile(path)

	req := &UploadFile{
		ContentType: contentType,
		Content:     fileBytes,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	d.resp, err = restclient.SendRequest(url, reqBytes)
	return err
}

func (d *FileHandlerSteps) resolveFile(url string) error {
	logger.Infof("Resolving file: %s", url)

	remainingAttempts := 20
	for {
		var err error
		d.resp, err = restclient.SendResolveRequest(url)
		if err != nil {
			return err
		}

		bddtests.SetResponse(string(d.resp.Payload))

		if d.resp.StatusCode == http.StatusNotFound {
			logger.Infof("File not found: %s. Remaining attempts: %d", url, remainingAttempts)
			remainingAttempts--
			if remainingAttempts > 0 {
				time.Sleep(time.Second)
				continue
			}
		}

		return nil
	}
}

func (d *FileHandlerSteps) checkErrorResp(errorMsg string) error {
	if !strings.Contains(d.resp.ErrorMsg, errorMsg) {
		return errors.Errorf("error resp %s doesn't contain %s", d.resp.ErrorMsg, errorMsg)
	}
	return nil
}

func (d *FileHandlerSteps) retrievedFileContains(msg string) error {
	if d.resp.ErrorMsg != "" {
		return errors.Errorf("error resp: [%s]", d.resp.ErrorMsg)
	}

	logger.Infof("check success resp %s contain %s", string(d.resp.Payload), msg)
	if !strings.Contains(string(d.resp.Payload), msg) {
		return errors.Errorf("success resp %s doesn't contain %s", d.resp.Payload, msg)
	}
	return nil
}

func (d *FileHandlerSteps) checkErrorResponse(statusCode int, msg string) error {
	if d.resp.StatusCode != statusCode {
		return errors.Errorf("expecting status code %d but got %d", statusCode, d.resp.StatusCode)
	}

	if d.resp.ErrorMsg != msg {
		return errors.Errorf("expecting error message [%s] but got [%s]", msg, d.resp.ErrorMsg)
	}

	return nil
}

func (d *FileHandlerSteps) saveIDToVariable(varName string) error {
	if d.resp.ErrorMsg != "" {
		return errors.Errorf("error resp: [%s]", d.resp.ErrorMsg)
	}

	id := ""
	if err := json.Unmarshal(d.resp.Payload, &id); err != nil {
		return err
	}

	logger.Infof("Saving ID [%s] to variable [%s]", id, varName)

	bddtests.SetVar(varName, id)
	return nil
}

func (d *FileHandlerSteps) saveDocIDToVariable(varName string) error {
	if d.resp.ErrorMsg != "" {
		return errors.Errorf("error resp: [%s]", d.resp.ErrorMsg)
	}

	doc := document.Document{}
	if err := json.Unmarshal(d.resp.Payload, &doc); err != nil {
		return err
	}

	logger.Infof("Got doc %v", doc)
	logger.Infof("Saving ID [%s] to variable [%s]", doc["id"], varName)

	bddtests.SetVar(varName, doc["id"].(string))
	return nil
}

func (d *FileHandlerSteps) setJSONPatchVar(varName, patch string) error {
	var p []interface{}
	err := json.Unmarshal([]byte(patch), &p)
	if err != nil {
		panic(err)
	}

	obj, err := bddtests.ResolveVars(p)
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	logger.Infof("Setting variable [%s] to JSON patch %s", varName, bytes)

	bddtests.SetVar(varName, string(bytes))

	return nil
}

func (d *FileHandlerSteps) getOpaqueDocument(content string) string {
	doc, _ := document.FromBytes([]byte(content))
	bytes, _ := doc.Bytes()
	return string(bytes)
}

func (d *FileHandlerSteps) getUpdateRequest(uniqueSuffix string, patch jsonpatch.Patch) ([]byte, error) {
	return helper.NewUpdateRequest(&helper.UpdateRequestInfo{
		DidUniqueSuffix: uniqueSuffix,
		Patch:           patch,
		UpdateOTP:       docutil.EncodeToString([]byte(updateOTP)),
		MultihashCode:   sha2_256,
	})
}

func getFile(path string) []byte {
	r, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return data
}

func getUniqueSuffix(docID string) string {
	pos := strings.LastIndex(docID, docutil.NamespaceDelimiter)
	if pos == -1 {
		return docID
	}

	return docID[pos+1:]
}

// UploadFile contains the file upload request
type UploadFile struct {
	ContentType string `json:"contentType"`
	Content     []byte `json:"content"`
}

// RegisterSteps registers did sidetree steps
func (d *FileHandlerSteps) RegisterSteps(s *godog.Suite) {
	s.Step(`^client sends request to "([^"]*)" to create document with content "([^"]*)" in namespace "([^"]*)"$`, d.createDocument)
	s.Step(`^client sends request to "([^"]*)" to update document "([^"]*)" with patch "([^"]*)"$`, d.updateDocument)
	s.Step(`^client sends request to "([^"]*)" to retrieve file$`, d.resolveFile)
	s.Step(`^client sends request to "([^"]*)" to upload file "([^"]*)" with content type "([^"]*)"$`, d.uploadFile)
	s.Step(`^the ID of the file is saved to variable "([^"]*)"`, d.saveIDToVariable)
	s.Step(`^the ID of the returned document is saved to variable "([^"]*)"`, d.saveDocIDToVariable)
	s.Step(`^the retrieved file contains "([^"]*)"$`, d.retrievedFileContains)
	s.Step(`^the response has status code (\d+) and error message "([^"]*)"$`, d.checkErrorResponse)
	s.Step(`^variable "([^"]*)" is assigned the JSON patch '([^']*)'$`, d.setJSONPatchVar)
}
