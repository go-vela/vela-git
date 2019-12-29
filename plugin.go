// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

// Default represents the CLI flags provided by default from Vela.
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

// Netrc represents the CLI flags used for creating the .netrc file.
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

// Optional represents the CLI flags to enable extra plugin functionality.
type Optional struct {
	// enable fetching of submodules
	Submodules bool
	// enable fetching of tags
	Tags bool
}

// Plugin represents the CLI flags loaded for the plugin.
type Plugin struct {
	// default arguments loaded for the plugin
	Default *Default
	// netrc arguments loaded for the plugin
	Netrc *Netrc
	// optional arguments loaded for the plugin
	Optional *Optional
}

// executeCommand runs the provided command and sanitizes the output.
func executeCommand(e *exec.Cmd) {
	e.Stdout = os.Stdout
	e.Stderr = os.Stderr
	fmt.Println("$", strings.Join(e.Args, " "))
	err := e.Run()
	if err != nil {
		log.Fatal(err)
	}
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

	executeCommand(exec.Command("git", "init"))

	executeCommand(exec.Command("git", "remote", "add", "origin", p.Default.Remote))

	executeCommand(exec.Command("git", "remote", "--verbose"))

	if p.Optional.Tags {
		executeCommand(exec.Command("git", "fetch", "--tags", "origin", p.Default.Ref))
	} else {
		executeCommand(exec.Command("git", "fetch", "--no-tags", "origin", p.Default.Ref))
	}

	executeCommand(exec.Command("git", "reset", "--hard", p.Default.Sha))

	if p.Optional.Submodules {
		executeCommand(exec.Command("git", "submodule", "update", "--init"))
	}

	return nil
}
