// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Repo represents the plugin configuration for repo information.
type Repo struct {
	// full remote url for cloning
	Remote string
	// enable fetching of submodules
	Submodules bool
	// enable fetching of tags
	Tags bool
}

// Validate verifies the Repo is properly configured.
func (r *Repo) Validate() error {
	logrus.Trace("validating repo plugin configuration")

	// verify remote is provided
	if len(r.Remote) == 0 {
		return fmt.Errorf("no repo remote provided")
	}

	return nil
}
