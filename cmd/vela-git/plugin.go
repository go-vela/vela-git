// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/spf13/afero"
)

var appFS = afero.NewOsFs()

// Plugin represents the CLI configuration loaded for the plugin.
type Plugin struct {
	// build arguments loaded for the plugin
	Build *Build
	// netrc arguments loaded for the plugin
	Netrc *Netrc
	// repo arguments loaded for the plugin
	Repo *Repo
}

// Exec formats the commands for cloning a git repository
func (p *Plugin) Exec() error {
	if len(p.Build.Path) == 0 {
		err := os.MkdirAll(p.Build.Path, 0777)
		if err != nil {
			return err
		}
	}

	err := p.Netrc.Write()
	if err != nil {
		return err
	}

	err = os.Chdir(p.Build.Path)
	if err != nil {
		return err
	}

	err = execCmd(initCmd())
	if err != nil {
		return err
	}

	err = execCmd(remoteAddCmd(p.Repo.Remote))
	if err != nil {
		return err
	}

	err = execCmd(remoteVerboseCmd())
	if err != nil {
		return err
	}

	if p.Repo.Tags {
		err = execCmd(fetchTagsCmd(p.Build.Ref))
		if err != nil {
			return err
		}
	} else {
		err = execCmd(fetchNoTagsCmd(p.Build.Ref))
		if err != nil {
			return err
		}
	}

	err = execCmd(resetCmd(p.Build.Sha))
	if err != nil {
		return err
	}

	if p.Repo.Submodules {
		err = execCmd(submoduleCmd())
		if err != nil {
			return err
		}
	}

	return nil
}

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
