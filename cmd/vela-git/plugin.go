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

// Build represents the CLI configuration for build information.
type Build struct {
	// full path to workspace
	Path string
	// reference generated for commit
	Ref string
	// SHA-1 hash generated for commit
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

// Repo represents the CLI configuration for repo information.
type Repo struct {
	// full remote url for cloning
	Remote string
	// enable fetching of submodules
	Submodules bool
	// enable fetching of tags
	Tags bool
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

	executeCommand(exec.Command("git", "init"))

	executeCommand(exec.Command("git", "remote", "add", "origin", p.Repo.Remote))

	executeCommand(exec.Command("git", "remote", "--verbose"))

	if p.Repo.Tags {
		executeCommand(exec.Command("git", "fetch", "--tags", "origin", p.Build.Ref))
	} else {
		executeCommand(exec.Command("git", "fetch", "--no-tags", "origin", p.Build.Ref))
	}

	executeCommand(exec.Command("git", "reset", "--hard", p.Build.Sha))

	if p.Repo.Submodules {
		executeCommand(exec.Command("git", "submodule", "update", "--init"))
	}

	return nil
}
