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

func TestGit_fetchTagsCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"fetch",
		"--tags",
		"origin",
		"refs/heads/master",
	)

	got := fetchTagsCmd("refs/heads/master")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("fetchTagsCmd is %v, want %v", got, want)
	}
}

func TestGit_fetchNoTagsCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"fetch",
		"--no-tags",
		"origin",
		"refs/heads/master",
	)

	got := fetchNoTagsCmd("refs/heads/master")

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

func TestGit_remoteAddCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		"git",
		"remote",
		"add",
		"origin",
		"https://github.com/octocat/hello-world.git",
	)

	got := remoteAddCmd("https://github.com/octocat/hello-world.git")

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
		"7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
	)

	got := resetCmd("7fd1a60b01f91b314f59955a4e4d4e80d8edf11d")

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
