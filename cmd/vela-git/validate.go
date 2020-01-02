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

	// validate build configuration
	err := validateBuild(c)
	if err != nil {
		return err
	}

	// validate netrc configuration
	err = validateNetrc(c)
	if err != nil {
		return err
	}

	// validate repo configuration
	err = validateRepo(c)
	if err != nil {
		return err
	}

	return nil
}

// validateBuild is a helper function to validate the build CLI configuration.
func validateBuild(c *cli.Context) error {
	logrus.Trace("validating build CLI configuration")

	if len(c.String("build.sha")) == 0 {
		return fmt.Errorf("build.sha (PARAMETER_SHA or BUILD_COMMIT) flag is not set")
	}

	if len(c.String("build.path")) == 0 {
		return fmt.Errorf("build.path (PARAMETER_PATH or BUILD_WORKSPACE) flag is not set")
	}

	if len(c.String("build.ref")) == 0 {
		return fmt.Errorf("build.ref (PARAMETER_REF or BUILD_REF) flag is not set")
	}

	return nil
}

// validateNetrc is a helper function to validate the netrc CLI configuration.
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

// validateRepo is a helper function to validate the repo CLI configuration.
func validateRepo(c *cli.Context) error {
	logrus.Trace("validating repo CLI configuration")

	if len(c.String("repo.remote")) == 0 {
		return fmt.Errorf("repo.remote (PARAMETER_REMOTE or REPOSITORY_CLONE) flag is not set")
	}

	return nil
}
