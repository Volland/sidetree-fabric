/*
    Copyright SecureKey Technologies Inc. All Rights Reserved.

    SPDX-License-Identifier: Apache-2.0
*/

var {defineSupportCode} = require('cucumber');

defineSupportCode(function ({And, But, Given, Then, When}) {
    And(/^the hash of the base64-encoded value "([^"]*)" equals "([^"]*)"$/, function (arg1, arg2, callback) {
        callback.pending();
    });
    And(/^the hash of the base64URL-encoded value "([^"]*)" equals "([^"]*)"$/, function (arg1, arg2, callback) {
        callback.pending();
    });
    And(/^core index file URI is parsed from anchor string "([^"]*)" and saved to variable "([^"]*)"$/, function (arg1, arg2, callback) {
        callback.pending();
    });
    And(/^response is decompressed using "([^"]*)"$/, function (arg1, arg2, callback) {
        callback.pending();
    });
    And(/^core index file URI is parsed from transaction info "([^"]*)" and saved to variable "([^"]*)"$/, function (arg1, arg2, callback) {
        callback.pending();
    });
});
