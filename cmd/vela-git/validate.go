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
	err := validateBuild(p.Build)
	if err != nil {
		return err
	}

	// validate netrc configuration
	err = validateNetrc(p.Netrc)
	if err != nil {
		return err
	}

	// validate repo configuration
	err = validateRepo(p.Repo)
	if err != nil {
		return err
	}

	return nil
}

// validateBuild is a helper function to verify the build plugin configuration.
func validateBuild(b *Build) error {
	logrus.Trace("validating build plugin configuration")

	if len(b.Path) == 0 {
		return fmt.Errorf("no build path provided")
	}

	if len(b.Sha) == 0 {
		return fmt.Errorf("no build sha provided")
	}

	if len(b.Ref) == 0 {
		return fmt.Errorf("no build ref provided")
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

// validateRepo is a helper function to verify the repo CLI configuration.
func validateRepo(r *Repo) error {
	logrus.Trace("validating repo CLI configuration")

	if len(r.Remote) == 0 {
		return fmt.Errorf("no repo remote provided")
	}

	return nil
}
