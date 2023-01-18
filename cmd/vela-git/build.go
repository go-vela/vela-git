// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Build represents the plugin configuration for build information.
type Build struct {
	// full path to workspace
	Path string
	// reference generated for commit
	Ref string
	// SHA-1 hash generated for commit
	Sha string
	// depth at which to fetch with
	Depth string
}

// Validate verifies the Build is properly configured.
func (b *Build) Validate() error {
	logrus.Trace("validating build plugin configuration")

	// verify path is provided
	if len(b.Path) == 0 {
		return fmt.Errorf("no build path provided")
	}

	// verify reference is provided
	if len(b.Ref) == 0 {
		return fmt.Errorf("no build ref provided")
	}

	// verify sha is provided
	if len(b.Sha) == 0 {
		return fmt.Errorf("no build sha provided")
	}

	return nil
}
