// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os/exec"
	"slices"
	"testing"
)

func TestGit_execCmd(t *testing.T) {
	// setup types
	e := exec.CommandContext(t.Context(), "echo", "hello")

	err := execCmd(e)
	if err != nil {
		t.Errorf("execCmd returned err: %v", err)
	}
}

func TestGit_fetchCmdWithTags(t *testing.T) {
	// setup types
	want := exec.CommandContext(
		t.Context(),
		"git",
		"fetch",
		"--tags",
		"--depth",
		"10",
		"origin",
		"refs/heads/main",
	)

	got := fetchCmd(t.Context(), "refs/heads/main", true, "10")

	if want.Path != got.Path {
		t.Errorf("fetchCmd Path is %s, want %s", got.Path, want.Path)
	}

	if !slices.Equal(got.Args, want.Args) {
		t.Errorf("fetchCmd Args are %v, want %v", got.Args, want.Args)
	}
}

func TestGit_fetchCmdNoTags(t *testing.T) {
	// setup types
	want := exec.CommandContext(
		t.Context(),
		"git",
		"fetch",
		"--no-tags",
		"--depth",
		"100",
		"origin",
		"refs/heads/main",
	)

	got := fetchCmd(t.Context(), "refs/heads/main", false, "")

	if want.Path != got.Path {
		t.Errorf("fetchNoTagsCmd Path is %s, want %s", got.Path, want.Path)
	}

	if !slices.Equal(got.Args, want.Args) {
		t.Errorf("fetchNoTagsCmd Args are %v, want %v", got.Args, want.Args)
	}
}

func TestGit_initCmd(t *testing.T) {
	// setup types
	want := exec.CommandContext(
		t.Context(),
		"git",
		"init",
	)

	got := initCmd(t.Context())

	if want.Path != got.Path {
		t.Errorf("initCmd Path is %s, want %s", got.Path, want.Path)
	}

	if !slices.Equal(got.Args, want.Args) {
		t.Errorf("initCmd Args are %v, want %v", got.Args, want.Args)
	}
}

func TestGit_defaultBranchCmd(t *testing.T) {
	// setup types
	want := exec.CommandContext(
		t.Context(),
		"git",
		"config",
		"--global",
		"init.defaultBranch",
		"main",
	)

	got := defaultBranchCmd(t.Context(), "main")

	if want.Path != got.Path {
		t.Errorf("defaultBranchCmd Path is %s, want %s", got.Path, want.Path)
	}

	if !slices.Equal(got.Args, want.Args) {
		t.Errorf("defaultBranchCmd Args are %v, want %v", got.Args, want.Args)
	}
}

func TestGit_remoteAddCmd(t *testing.T) {
	// setup types
	want := exec.CommandContext(
		t.Context(),
		"git",
		"remote",
		"add",
		"origin",
		"https://github.com/go-vela/vela-git-test.git",
	)

	got := remoteAddCmd(t.Context(), "https://github.com/go-vela/vela-git-test.git")

	if want.Path != got.Path {
		t.Errorf("remoteAddCmd Path is %s, want %s", got.Path, want.Path)
	}

	if !slices.Equal(got.Args, want.Args) {
		t.Errorf("remoteAddCmd Args are %v, want %v", got.Args, want.Args)
	}
}

func TestGit_remoteVerboseCmd(t *testing.T) {
	// setup types
	want := exec.CommandContext(
		t.Context(),
		"git",
		"remote",
		"--verbose",
	)

	got := remoteVerboseCmd(t.Context())

	if want.Path != got.Path {
		t.Errorf("remoteVerboseCmd Path is %s, want %s", got.Path, want.Path)
	}

	if !slices.Equal(got.Args, want.Args) {
		t.Errorf("remoteVerboseCmd Args are %v, want %v", got.Args, want.Args)
	}
}

func TestGit_resetCmd(t *testing.T) {
	// setup types
	want := exec.CommandContext(
		t.Context(),
		"git",
		"reset",
		"--hard",
		"ee1e671529ad86a11ed628a04b37829e71783682",
	)
	want.Env = append(want.Env, "GIT_LFS_SKIP_SMUDGE=1")

	got := resetCmd(t.Context(), "ee1e671529ad86a11ed628a04b37829e71783682")

	if want.Path != got.Path {
		t.Errorf("resetCmd Path is %s, want %s", got.Path, want.Path)
	}

	if !slices.Equal(got.Args, want.Args) {
		t.Errorf("resetCmd Args are %v, want %v", got.Args, want.Args)
	}
}

func TestGit_submoduleCmd(t *testing.T) {
	// setup types
	want := exec.CommandContext(
		t.Context(),
		"git",
		"submodule",
		"update",
		"--init",
	)

	got := submoduleCmd(t.Context())

	if want.Path != got.Path {
		t.Errorf("submoduleCmd Path is %s, want %s", got.Path, want.Path)
	}

	if !slices.Equal(got.Args, want.Args) {
		t.Errorf("submoduleCmd Args are %v, want %v", got.Args, want.Args)
	}
}

func TestGit_getLFSCmd(t *testing.T) {
	// setup types
	want := exec.CommandContext(
		t.Context(),
		"git",
		"lfs",
		"pull",
	)

	got := getLFSCmd(t.Context())

	if want.Path != got.Path {
		t.Errorf("getLFSCmd Path is %s, want %s", got.Path, want.Path)
	}

	if !slices.Equal(got.Args, want.Args) {
		t.Errorf("getLFSCmd Args are %v, want %v", got.Args, want.Args)
	}
}
