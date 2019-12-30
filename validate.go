// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// validate is a helper function to validate the CLI configuration.
func validate(c *cli.Context) error {
	logrus.Debug("validating CLI configuration")

	// validate default configuration
	err := validateDefault(c)
	if err != nil {
		return err
	}

	// validate netrc configuration
	err = validateNetrc(c)
	if err != nil {
		return err
	}

	return nil
}

// helper function to validate the default CLI configuration.
func validateDefault(c *cli.Context) error {
	logrus.Trace("validating default CLI configuration")

	if len(c.String("commit")) == 0 {
		return fmt.Errorf("commit (PARAMETER_COMMIT or BUILD_COMMIT) flag is not set")
	}

	if len(c.String("path")) == 0 {
		return fmt.Errorf("path (PARAMETER_PATH or BUILD_WORKSPACE) flag is not set")
	}

	if len(c.String("ref")) == 0 {
		return fmt.Errorf("ref (PARAMETER_REF or BUILD_REF) flag is not set")
	}

	if len(c.String("remote")) == 0 {
		return fmt.Errorf("remote (PARAMETER_REMOTE or REPOSITORY_CLONE) flag is not set")
	}

	return nil
}

// helper function to validate the netrc CLI configuration.
func validateNetrc(c *cli.Context) error {
	logrus.Trace("validating netrc CLI configuration")

	if len(c.String("netrc.machine")) == 0 {
		return fmt.Errorf("netrc.machine (PARAMETER_NETRC_MACHINE or VELA_NETRC_MACHINE) flag is not set")
	}

	if len(c.String("netrc.username")) == 0 {
		return fmt.Errorf("netrc.username (PARAMETER_NETRC_USERNAME or VELA_NETRC_USERNAME) flag is not set")
	}

	if len(c.String("netrc.password")) == 0 {
		return fmt.Errorf("netrc.password (PARAMETER_NETRC_PASSWORD or VELA_NETRC_PASSWORD) flag is not set")
	}

	return nil
}
