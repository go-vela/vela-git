package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

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
