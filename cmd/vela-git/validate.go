// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/sirupsen/logrus"
)

// Validate verifies the plugin is properly configured.
func (p *Plugin) Validate() error {
	logrus.Debug("validating plugin configuration")

	// validate build configuration
	err := p.Build.Validate()
	if err != nil {
		return err
	}

	// validate netrc configuration
	err = p.Netrc.Validate()
	if err != nil {
		return err
	}

	// validate repo configuration
	err = p.Repo.Validate()
	if err != nil {
		return err
	}

	return nil
}
