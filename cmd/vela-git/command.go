// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
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

// fetchCmd is a helper function to
// download all objects, including tags,
// from the ref for a git repo.
func fetchCmd(ctx context.Context, ref string, includeTags bool, depth string) *exec.Cmd {
	logrus.Trace("returning fetchCmd")

	args := []string{"fetch"}

	if includeTags {
		args = append(args, "--tags")
	} else {
		args = append(args, "--no-tags")
	}

	if depth != "" {
		args = append(args, "--depth", depth)
	} else {
		args = append(args, "--depth", "100")
	}

	args = append(args, "origin", ref)

	return exec.CommandContext(
		ctx,
		"git",
		args...,
	)
}

// initCmd is a helper function to
// create an empty or reinitialize
// an existing git repo.
func initCmd(ctx context.Context) *exec.Cmd {
	logrus.Trace("returning initCmd")

	return exec.CommandContext(
		ctx,
		"git",
		"init",
	)
}

// remoteAddCmd is a helper function to
// add a remote for a git repo.
func remoteAddCmd(ctx context.Context, remote string) *exec.Cmd {
	logrus.Trace("returning remoteAddCmd")

	return exec.CommandContext(
		ctx,
		"git",
		"remote",
		"add",
		"origin",
		remote,
	)
}

// remoteVerboseCmd is a helper function to
// output al remotes for a git repo.
func remoteVerboseCmd(ctx context.Context) *exec.Cmd {
	logrus.Trace("returning remoteVerboseCmd")

	return exec.CommandContext(
		ctx,
		"git",
		"remote",
		"--verbose",
	)
}

// defaultBranchCmd is a helper function
// to set init.defaultBranch in git
// available to override default branch
// name when initializing a new repo.
func defaultBranchCmd(ctx context.Context, branch string) *exec.Cmd {
	logrus.Trace("returning defaultBranchCmd")

	return exec.CommandContext(
		ctx,
		"git",
		"config",
		"--global",
		"init.defaultBranch",
		branch,
	)
}

// resetCmd is a helper function to
// hard reset the current HEAD to
// the sha for a git repo.
func resetCmd(ctx context.Context, sha string) *exec.Cmd {
	logrus.Trace("returning resetCmd")

	cmd := exec.CommandContext(
		ctx,
		"git",
		"reset",
		"--hard",
		sha,
	)

	// skip resolving LFS objects by default
	// https://github.com/git-lfs/git-lfs/blob/main/docs/man/git-lfs-smudge.adoc
	cmd.Env = append(cmd.Env, "GIT_LFS_SKIP_SMUDGE=1")

	return cmd
}

// getLFSCmd is a helper function to
// resolve LFS objects.
func getLFSCmd(ctx context.Context) *exec.Cmd {
	logrus.Trace("returning command to pull LFS objects")

	return exec.CommandContext(
		ctx,
		"git",
		"lfs",
		"pull",
	)
}

// submoduleCmd is a helper function to
// update the registered submodules to
// the expected states for a git repo.
func submoduleCmd(ctx context.Context) *exec.Cmd {
	logrus.Trace("returning submoduleCmd")

	return exec.CommandContext(
		ctx,
		"git",
		"submodule",
		"update",
		"--init",
	)
}

// versionCmd is a helper function to output
// the client version information.
func versionCmd(ctx context.Context) *exec.Cmd {
	logrus.Trace("creating git version command")

	// variable to store flags for command
	var flags []string

	// add flag for version git command
	flags = append(flags, "version")

	return exec.CommandContext(ctx, "git", flags...)
}
