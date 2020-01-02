// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

// Default represents the CLI configuration provided by default from Vela.
type Default struct {
	// path to clone repository to
	Path string
	// commit ref generated for commit
	Ref string
	// remote url for repository
	Remote string
	// commit sha to checkout to in repository
	Sha string
}

// Netrc represents the CLI configuration used for creating the .netrc file.
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

// Optional represents the CLI configuration to enable extra plugin functionality.
type Optional struct {
	// enable fetching of submodules
	Submodules bool
	// enable fetching of tags
	Tags bool
}

// Plugin represents the CLI configuration loaded for the plugin.
type Plugin struct {
	// default arguments loaded for the plugin
	Default *Default
	// netrc arguments loaded for the plugin
	Netrc *Netrc
	// optional arguments loaded for the plugin
	Optional *Optional
}

const netrcFile = `
machine %s
login %s
password %s
`

// writeNetrc creates a netrc file and returns the file.
func writeNetrc(machine, login, password string) error {
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

	return ioutil.WriteFile(path, []byte(out), 0600)
}

// Exec formats the commands for cloning a git repository
func (p Plugin) Exec() error {
	if len(p.Default.Path) == 0 {
		err := os.MkdirAll(p.Default.Path, 0777)
		if err != nil {
			return err
		}
	}

	err := writeNetrc(p.Netrc.Machine, p.Netrc.Username, p.Netrc.Password)
	if err != nil {
		return err
	}

	err = os.Chdir(p.Default.Path)
	if err != nil {
		return err
	}

	err = execCmd(initCmd())
	if err != nil {
		return err
	}

	err = execCmd(remoteAddCmd(p.Default.Remote))
	if err != nil {
		return err
	}

	err = execCmd(remoteVerboseCmd())
	if err != nil {
		return err
	}

	if p.Optional.Tags {
		err = execCmd(fetchTagsCmd(p.Default.Ref))
		if err != nil {
			return err
		}
	} else {
		err = execCmd(fetchNoTagsCmd(p.Default.Ref))
		if err != nil {
			return err
		}
	}

	err = execCmd(resetCmd(p.Default.Sha))
	if err != nil {
		return err
	}

	if p.Optional.Submodules {
		err = execCmd(submoduleCmd())
		if err != nil {
			return err
		}
	}

	return nil
}
