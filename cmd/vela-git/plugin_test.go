// SPDX-License-Identifier: Apache-2.0

package main

import "testing"

func TestGit_Plugin_Exec(t *testing.T) {
	// setup directory
	dir := t.TempDir()

	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "master",
			Path:   dir,
			Ref:    "refs/heads/master",
			Sha:    "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
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

	err := p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}
}

func TestGit_Plugin_Exec_Submodules(t *testing.T) {
	// setup directory
	dir := t.TempDir()

	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "master",
			Path:   dir,
			Ref:    "refs/heads/master",
			Sha:    "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/octocat/hello-world.git",
			Submodules: true,
			Tags:       false,
		},
	}

	err := p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}
}

func TestGit_Plugin_Exec_Tags(t *testing.T) {
	// setup directory
	dir := t.TempDir()

	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "master",
			Path:   dir,
			Ref:    "refs/heads/master",
			Sha:    "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/octocat/hello-world.git",
			Submodules: false,
			Tags:       true,
		},
	}

	err := p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}
}

func TestGit_Plugin_Exec_LFS(t *testing.T) {
	// setup directory
	dir := t.TempDir()

	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "master",
			Path:   dir,
			Ref:    "refs/heads/main",
			Sha:    "efd6f0c16e6593c5468037ae408c52b4980a2666",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/go-vela/community.git",
			Submodules: false,
			Tags:       true,
			LFS:        true,
		},
	}

	err := p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}
}

func TestGit_Plugin_Validate(t *testing.T) {
	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "master",
			Path:   "/home/octocat_hello-world_1",
			Ref:    "refs/heads/master",
			Sha:    "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
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
