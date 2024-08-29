// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestGit_execCmd(t *testing.T) {
	// setup types
	e := exec.Command("echo", "hello")

	err := execCmd(e)
	if err != nil {
		t.Errorf("execCmd returned err: %v", err)
	}
}

func TestGit_fetchCmdWithTags(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"fetch",
		"--tags",
		"--depth",
		"10",
		"origin",
		"refs/heads/main",
	)

	got := fetchCmd("refs/heads/main", true, "10")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("fetchTagsCmd is %v, want %v", got, want)
	}
}

func TestGit_fetchCmdNoTags(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"fetch",
		"--no-tags",
		"--depth",
		"100",
		"origin",
		"refs/heads/main",
	)

	got := fetchCmd("refs/heads/main", false, "")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("fetchNoTagsCmd is %v, want %v", got, want)
	}
}

func TestGit_initCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"init",
	)

	got := initCmd()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("initCmd is %v, want %v", got, want)
	}
}

func TestGit_defaultBranchCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"config",
		"--global",
		"init.defaultBranch",
		"main",
	)

	got := defaultBranchCmd("main")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("defaultBranchCmd is %v, want %v", got, want)
	}
}

func TestGit_remoteAddCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"remote",
		"add",
		"origin",
		"https://github.com/go-vela/vela-git-test.git",
	)

	got := remoteAddCmd("https://github.com/go-vela/vela-git-test.git")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("remoteAddCmd is %v, want %v", got, want)
	}
}

func TestGit_remoteVerboseCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"remote",
		"--verbose",
	)

	got := remoteVerboseCmd()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("remoteVerboseCmd is %v, want %v", got, want)
	}
}

func TestGit_resetCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"reset",
		"--hard",
		"ee1e671529ad86a11ed628a04b37829e71783682",
	)
	want.Env = append(want.Env, "GIT_LFS_SKIP_SMUDGE=1")

	got := resetCmd("ee1e671529ad86a11ed628a04b37829e71783682")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("resetCmd is %v, want %v", got, want)
	}
}

func TestGit_submoduleCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"submodule",
		"update",
		"--init",
	)

	got := submoduleCmd()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("submoduleCmd is %v, want %v", got, want)
	}
}

func TestGit_getLFSCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"lfs",
		"pull",
	)

	got := getLFSCmd()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("getLFSCmd is %v, want %v", got, want)
	}
}
