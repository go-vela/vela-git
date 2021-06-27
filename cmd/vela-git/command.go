// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// execCmd is a helper function to
// run the provided command.
func execCmd(e *exec.Cmd) error {
	logrus.Tracef("executing cmd %s", strings.Join(e.Args, " "))

	// set command stdout to OS stdout
	e.Stdout = os.Stdout
	// set command stderr to OS stderr
	e.Stderr = os.Stderr

	// output "trace" string for command
	fmt.Println("$", strings.Join(e.Args, " "))

	return e.Run()
}

// fetchTagsCmd is a helper function to
// download all objects, including tags,
// from the ref for a git repo.
func fetchTagsCmd(ref string, depth string) *exec.Cmd {
	logrus.Trace("returning fetchTagsCmd")

	args := []string{
		"fetch",
		"--tags",
		"origin",
		ref,
	}

	if depth != "" {
		args = append(args, []string{"--depth", depth}...)
	}

	return exec.Command(
		"git",
		args...,
	)
}

// fetchNoTagsCmd is a helper function to
// download all objects, excluding tags,
// from the ref for a git repo.
func fetchNoTagsCmd(ref string, depth string) *exec.Cmd {
	logrus.Trace("returning fetchNoTagsCmd")

	args := []string{
		"fetch",
		"--no-tags",
		"origin",
		ref,
	}

	if depth != "" {
		args = append(args, []string{"--depth", depth}...)
	}

	return exec.Command(
		"git",
		args...,
	)
}

// initCmd is a helper function to
// create an empty or reinitialize
// an existing git repo.
func initCmd() *exec.Cmd {
	logrus.Trace("returning initCmd")

	return exec.Command(
		"git",
		"init",
	)
}

// remoteAddCmd is a helper function to
// add a remote for a git repo.
func remoteAddCmd(remote string) *exec.Cmd {
	logrus.Trace("returning remoteAddCmd")

	return exec.Command(
		"git",
		"remote",
		"add",
		"origin",
		remote,
	)
}

// remoteVerboseCmd is a helper function to
// output al remotes for a git repo.
func remoteVerboseCmd() *exec.Cmd {
	logrus.Trace("returning remoteVerboseCmd")

	return exec.Command(
		"git",
		"remote",
		"--verbose",
	)
}

// resetCmd is a helper function to
// hard reset the current HEAD to
// the sha for a git repo.
func resetCmd(sha string) *exec.Cmd {
	logrus.Trace("returning resetCmd")

	return exec.Command(
		"git",
		"reset",
		"--hard",
		sha,
	)
}

// createTargetBranchCmd is a helper function to
// create a branch reference to the target branch
// for this pull request.
func createTargetBranchCmd(prTargetBranch string) *exec.Cmd {
	logrus.Trace("returning idkCmd")

	return exec.Command(
		"git",
		"branch",
		"-f",
		prTargetBranch,
		"origin/"+prTargetBranch,
	)
}

// submoduleCmd is a helper function to
// update the registered submodules to
// the expected states for a git repo.
func submoduleCmd() *exec.Cmd {
	logrus.Trace("returning submoduleCmd")

	return exec.Command(
		"git",
		"submodule",
		"update",
		"--init",
	)
}

// versionCmd is a helper function to output
// the client version information.
func versionCmd() *exec.Cmd {
	logrus.Trace("creating git version command")

	// variable to store flags for command
	var flags []string

	// add flag for version git command
	flags = append(flags, "version")

	return exec.Command("git", flags...)
}
