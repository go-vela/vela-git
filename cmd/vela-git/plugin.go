// SPDX-License-Identifier: Apache-2.0

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

	// configure default branch for init
	err = execCmd(defaultBranchCmd(p.Build.Branch))
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

	// fetch the repo
	err = execCmd(fetchCmd(p.Build.Ref, p.Repo.Tags, p.Build.Depth))
	if err != nil {
		return err
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

	// if LFS is enabled, get/resolve the LFS objects
	if p.Repo.LFS {
		err = execCmd(getLFSCmd())
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
