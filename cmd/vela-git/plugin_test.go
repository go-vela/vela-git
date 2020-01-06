// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGit_Plugin_Exec(t *testing.T) {
	// setup directory
	dir, err := ioutil.TempDir("/tmp", "vela_git_plugin_")
	if err != nil {
		t.Errorf("unable to create temp directory: %v", err)
	}

	// defer cleanup of directory
	defer func() {
		err := os.RemoveAll(dir)
		if err != nil {
			logrus.Fatalf("unable to remove temp directory %s: %v", dir, err)
		}
	}()

	// setup types
	p := &Plugin{
		Build: &Build{
			Path: dir,
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

	err = p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}
}

func TestGit_Plugin_Exec_Submodules(t *testing.T) {
	// setup directory
	dir, err := ioutil.TempDir("/tmp", "vela_git_plugin_")
	if err != nil {
		t.Errorf("unable to create temp directory: %v", err)
	}

	// defer cleanup of directory
	defer func() {
		err := os.RemoveAll(dir)
		if err != nil {
			logrus.Fatalf("unable to remove temp directory %s: %v", dir, err)
		}
	}()

	// setup types
	p := &Plugin{
		Build: &Build{
			Path: dir,
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
			Submodules: true,
			Tags:       false,
		},
	}

	err = p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}
}

func TestGit_Plugin_Exec_Tags(t *testing.T) {
	// setup directory
	dir, err := ioutil.TempDir("/tmp", "vela_git_plugin_")
	if err != nil {
		t.Errorf("unable to create temp directory: %v", err)
	}

	// defer cleanup of directory
	defer func() {
		err := os.RemoveAll(dir)
		if err != nil {
			logrus.Fatalf("unable to remove temp directory %s: %v", dir, err)
		}
	}()

	// setup types
	p := &Plugin{
		Build: &Build{
			Path: dir,
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
			Tags:       true,
		},
	}

	err = p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}
}

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
