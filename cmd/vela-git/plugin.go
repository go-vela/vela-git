// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/spf13/afero"
)

var appFS = afero.NewOsFs()

// Plugin represents the configuration loaded for the plugin.
type Plugin struct {
	// build arguments loaded for the plugin
	Build *Build
	// netrc arguments loaded for the plugin
	Netrc *Netrc
	// repo arguments loaded for the plugin
	Repo *Repo
}

// Exec formats and runs the commands for cloning a git repository.
func (p *Plugin) Exec() error {
	logrus.Debug("running plugin with provided configuration")

	// check if a build path is provided
	if len(p.Build.Path) > 0 {
		// send OS call to create path to build directory
		err := os.MkdirAll(p.Build.Path, 0777)
		if err != nil {
			return err
		}
	}

	// create .netrc file for authentication
	err := p.Netrc.Write()
	if err != nil {
		return err
	}

	// send OS call to change working directory to build directory
	err = os.Chdir(p.Build.Path)
	if err != nil {
		return err
	}

	// output the git version for troubleshooting
	err = execCmd(versionCmd())
	if err != nil {
		return err
	}

	// initialize git repo
	err = execCmd(initCmd())
	if err != nil {
		return err
	}

	// add remote to git repo
	err = execCmd(remoteAddCmd(p.Repo.Remote))
	if err != nil {
		return err
	}

	// output remotes for git repo
	err = execCmd(remoteVerboseCmd())
	if err != nil {
		return err
	}

	// check if repo tags are enabled
	if p.Repo.Tags {
		// fetch repo state with tags
		err = execCmd(fetchTagsCmd(p.Build.Ref, p.Build.Depth))
		if err != nil {
			return err
		}
	} else {
		// fetch repo state without tags
		err = execCmd(fetchNoTagsCmd(p.Build.Ref, p.Build.Depth))
		if err != nil {
			return err
		}
	}

	// check if it's a pull request
	if p.Repo.PrTargetBranch != "" {
		// fetch target branch state
		err = execCmd(fetchNoTagsCmd(p.Repo.PrTargetBranch, p.Build.Depth))
		if err != nil {
			return err
		}
		// create a reference branch to the target branch
		err = execCmd(createTargetBranchCmd(p.Repo.PrTargetBranch))
		if err != nil {
			return err
		}
	}

	// hard reset current state to build commit
	err = execCmd(resetCmd(p.Build.Sha))
	if err != nil {
		return err
	}

	// check if repo submodules are enabled
	if p.Repo.Submodules {
		// update submodules to expected state
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
