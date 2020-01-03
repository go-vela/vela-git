// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"

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
	err = validateNetrc(p.Netrc)
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

// validateNetrc is a helper function to verify the netrc CLI configuration.
func validateNetrc(n *Netrc) error {
	logrus.Trace("validating netrc plugin configuration")

	if len(n.Machine) == 0 {
		return fmt.Errorf("no netrc machine provided")
	}

	if len(n.Username) == 0 {
		return fmt.Errorf("no netrc username provided")
	}

	if len(n.Password) == 0 {
		return fmt.Errorf("no netrc password provided")
	}

	return nil
}
