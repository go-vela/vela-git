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

// Plugin represents a one to one mapping of the CLI flags.
type Plugin struct {
	Remote        string // remote url for repository
	Path          string // local path to clone repository too
	CommitSha     string // specific sha to checkout to in repository
	CommitRef     string // specific ref generated for commit
	Tags          bool   // allow fetching of tags
	Submodules    bool   // allow fetching of first level submodules
	NetrcMachine  string // netrc machine used for authentication
	NetrcUsername string // netrc username used for authentication
	NetrcPassword string // netrc password used for authentication
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

	if p.Path != "" {
		err := os.MkdirAll(p.Path, 0777)
		if err != nil {
			return err
		}
	}

	err := writeNetrc(p.NetrcMachine, p.NetrcUsername, p.NetrcPassword)
	if err != nil {
		return err
	}

	os.Mkdir(p.Path, 0777)
	os.Chdir(p.Path)

	executeCommand(exec.Command("git", "init"))

	executeCommand(exec.Command("git", "remote", "add", "origin", p.Remote))

	executeCommand(exec.Command("git", "remote", "--verbose"))

	if p.Tags {
		executeCommand(exec.Command("git", "fetch", "--tags", "origin", p.CommitRef))
	} else {
		executeCommand(exec.Command("git", "fetch", "--no-tags", "origin", p.CommitRef))
	}

	executeCommand(exec.Command("git", "reset", "--hard", p.CommitSha))

	if p.Submodules {
		executeCommand(exec.Command("git", "submodule", "update", "--init"))
	}

	return nil
}
