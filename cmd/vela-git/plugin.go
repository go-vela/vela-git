// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/afero"
)

var appFS = afero.NewOsFs()

// Netrc represents the CLI configuration for netrc information used for creating the .netrc file.
//
// https://www.gnu.org/software/inetutils/manual/html_node/The-_002enetrc-file.html
type Netrc struct {
	// remote machine name to communicate with
	Machine string
	// user name for communication with the remote machine
	Username string
	// password for communication with the remote machine
	Password string
}

// Plugin represents the CLI configuration loaded for the plugin.
type Plugin struct {
	// build arguments loaded for the plugin
	Build *Build
	// netrc arguments loaded for the plugin
	Netrc *Netrc
	// repo arguments loaded for the plugin
	Repo *Repo
}

const netrcFile = `
machine %s
login %s
password %s
`

// writeNetrc creates a netrc file and returns the file.
func writeNetrc(machine, login, password string) error {
	a := &afero.Afero{
		Fs: appFS,
	}

	if len(machine) == 0 || len(login) == 0 || len(password) == 0 {
		return nil
	}

	out := fmt.Sprintf(
		netrcFile,
		machine,
		login,
		password,
	)

	home := "/root"

	u, err := user.Current()
	if err == nil {
		home = u.HomeDir
	}

	path := filepath.Join(home, ".netrc")

	return a.WriteFile(path, []byte(out), 0600)
}

// Exec formats the commands for cloning a git repository
func (p *Plugin) Exec() error {
	if len(p.Build.Path) == 0 {
		err := os.MkdirAll(p.Build.Path, 0777)
		if err != nil {
			return err
		}
	}

	err := writeNetrc(p.Netrc.Machine, p.Netrc.Username, p.Netrc.Password)
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
