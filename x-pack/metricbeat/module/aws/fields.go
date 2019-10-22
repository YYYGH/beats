// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "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"
}