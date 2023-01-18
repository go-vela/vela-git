// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Repo represents the plugin configuration for repo information.
type Repo struct {
	// default branch of the Repo
	DefaultBranch string
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

	// verify default branch is provided
	if len(r.DefaultBranch) == 0 {
		return fmt.Errorf("no repo default branch provided")
	}

	// verify remote is provided
	if len(r.Remote) == 0 {
		return fmt.Errorf("no repo remote provided")
	}

	return nil
}
