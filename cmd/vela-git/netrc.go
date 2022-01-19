// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
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

	// verify machine is provided
	if len(n.Machine) == 0 {
		return fmt.Errorf("no netrc machine provided")
	}

	// verify username is provided
	if len(n.Username) == 0 {
		return fmt.Errorf("no netrc username provided")
	}

	// verify password is provided
	if len(n.Password) == 0 {
		return fmt.Errorf("no netrc password provided")
	}

	return nil
}

// Write creates a .netrc file in the home directory of the current user.
func (n *Netrc) Write() error {
	logrus.Trace("writing netrc configuration file")

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// check if machine, username and password are provided
	if len(n.Machine) == 0 || len(n.Username) == 0 || len(n.Password) == 0 {
		return nil
	}

	// create output string for .netrc file
	out := fmt.Sprintf(
		netrcFile,
		n.Machine,
		n.Username,
		n.Password,
	)

	// set default home directory for root user
	home := "/root"

	// capture current user running commands
	u, err := user.Current()
	if err == nil {
		// set home directory to current user
		home = u.HomeDir
	}

	// create full path for .netrc file
	path := filepath.Join(home, ".netrc")

	// send Filesystem call to create directory path for .netrc file
	err = a.Fs.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return err
	}

	return a.WriteFile(path, []byte(out), 0600)
}
