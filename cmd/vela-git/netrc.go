// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/sirupsen/logrus"
)

const netrcFile = `
machine %s
login %s
password %s
`

// Netrc represents the netrc configuration used for creating the .netrc file.
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

// Validate verifies the Netrc is properly configured.
func (n *Netrc) Validate() error {
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

// Write creates a .netrc file in the home directory of the current user.
func (n *Netrc) Write() error {
	a := &afero.Afero{
		Fs: appFS,
	}

	if len(n.Machine) == 0 || len(n.Username) == 0 || len(n.Password) == 0 {
		return nil
	}

	out := fmt.Sprintf(
		netrcFile,
		n.Machine,
		n.Username,
		n.Password,
	)

	home := "/root"

	u, err := user.Current()
	if err == nil {
		home = u.HomeDir
	}

	path := filepath.Join(home, ".netrc")

	return a.WriteFile(path, []byte(out), 0600)
}
