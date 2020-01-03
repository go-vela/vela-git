package main

import (
	"testing"
)

func TestGit_Plugin_Validate(t *testing.T) {
	// setup types
	p := &Plugin{
		Build: &Build{
			Path: "/home/octocat_hello-world_1",
			Ref:  "refs/heads/master",
			Sha:  "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/octocat/hello-world.git",
			Submodules: false,
			Tags:       false,
		},
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGit_Plugin_Validate_NoBuild(t *testing.T) {
	// setup types
	p := &Plugin{
		Build: &Build{},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/octocat/hello-world.git",
			Submodules: false,
			Tags:       false,
		},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestGit_Plugin_Validate_NoNetrc(t *testing.T) {
	// setup types
	p := &Plugin{
		Build: &Build{
			Path: "/home/octocat_hello-world_1",
			Ref:  "refs/heads/master",
			Sha:  "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		},
		Netrc: &Netrc{},
		Repo: &Repo{
			Remote:     "https://github.com/octocat/hello-world.git",
			Submodules: false,
			Tags:       false,
		},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestGit_Plugin_Validate_NoRepo(t *testing.T) {
	// setup types
	p := &Plugin{
		Build: &Build{
			Path: "/home/octocat_hello-world_1",
			Ref:  "refs/heads/master",
			Sha:  "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestGit_validateBuild(t *testing.T) {
	// setup types
	b := &Build{
		Path: "/home/octocat_hello-world_1",
		Ref:  "refs/heads/master",
		Sha:  "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
	}

	err := validateBuild(b)
	if err != nil {
		t.Errorf("validateBuild returned err: %v", err)
	}
}

func TestGit_validateBuild_NoPath(t *testing.T) {
	// setup types
	b := &Build{
		Ref: "refs/heads/master",
		Sha: "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
	}

	err := validateBuild(b)
	if err == nil {
		t.Errorf("validateBuild should have returned err")
	}
}

func TestGit_validateBuild_NoRef(t *testing.T) {
	// setup types
	b := &Build{
		Path: "/home/octocat_hello-world_1",
		Sha:  "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
	}

	err := validateBuild(b)
	if err == nil {
		t.Errorf("validateBuild should have returned err")
	}
}

func TestGit_validateBuild_NoSha(t *testing.T) {
	// setup types
	b := &Build{
		Path: "/home/octocat_hello-world_1",
		Ref:  "refs/heads/master",
	}

	err := validateBuild(b)
	if err == nil {
		t.Errorf("validateBuild should have returned err")
	}
}

func TestGit_validateNetrc(t *testing.T) {
	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Username: "octocat",
		Password: "superSecretPassword",
	}

	err := validateNetrc(n)
	if err != nil {
		t.Errorf("validateNetrc returned err: %v", err)
	}
}

func TestGit_validateNetrc_NoMachine(t *testing.T) {
	// setup types
	n := &Netrc{
		Username: "octocat",
		Password: "superSecretPassword",
	}

	err := validateNetrc(n)
	if err == nil {
		t.Errorf("validateNetrc should have returned err")
	}
}

func TestGit_validateNetrc_NoUsername(t *testing.T) {
	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Password: "superSecretPassword",
	}

	err := validateNetrc(n)
	if err == nil {
		t.Errorf("validateNetrc should have returned err")
	}
}

func TestGit_validateNetrc_NoPassword(t *testing.T) {
	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Username: "octocat",
	}

	err := validateNetrc(n)
	if err == nil {
		t.Errorf("validateNetrc should have returned err")
	}
}

func TestGit_validateRepo(t *testing.T) {
	// setup types
	r := &Repo{
		Remote:     "https://github.com/octocat/hello-world.git",
		Submodules: false,
		Tags:       false,
	}

	err := validateRepo(r)
	if err != nil {
		t.Errorf("validateRepo returned err: %v", err)
	}
}

func TestGit_validateRepo_NoRemote(t *testing.T) {
	// setup types
	r := &Repo{
		Submodules: false,
		Tags:       false,
	}

	err := validateRepo(r)
	if err == nil {
		t.Errorf("validateRepo should have returned err")
	}
}
