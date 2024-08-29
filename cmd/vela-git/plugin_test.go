// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"path"
	"testing"
)

func TestGit_Plugin_Exec(t *testing.T) {
	// setup directory
	dir := t.TempDir()

	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "main",
			Path:   dir,
			Ref:    "refs/heads/main",
			Sha:    "ee1e671529ad86a11ed628a04b37829e71783682",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/go-vela/vela-git-test.git",
			Submodules: false,
			Tags:       false,
			LFS:        false,
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
	// expected file for resolved LFS object is 100kb
	wantFileSizeBytes := int64(100 * 1024)

	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "main",
			Path:   dir,
			Ref:    "refs/heads/main",
			Sha:    "ee1e671529ad86a11ed628a04b37829e71783682",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/go-vela/vela-git-test.git",
			Submodules: true,
			Tags:       false,
		},
	}

	err := p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}

	testFile, err := os.Stat(path.Join(dir, "100k-test.bin"))
	if err != nil {
		t.Errorf("Exec unable to get info on test file")
	}

	if testFile.Size() == wantFileSizeBytes {
		t.Errorf("Exec resulted in unexpected file size in repo object - want %d, got %d", wantFileSizeBytes, testFile.Size())
	}
}

func TestGit_Plugin_Exec_Tags(t *testing.T) {
	// setup directory
	dir := t.TempDir()

	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "main",
			Path:   dir,
			Ref:    "refs/heads/main",
			Sha:    "ee1e671529ad86a11ed628a04b37829e71783682",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/go-vela/vela-git-test.git",
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
	// expected file for resolved LFS object is 100kb
	wantFileSizeBytes := int64(100 * 1024)

	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "main",
			Path:   dir,
			Ref:    "refs/heads/main",
			Sha:    "ee1e671529ad86a11ed628a04b37829e71783682",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/go-vela/vela-git-test.git",
			Submodules: false,
			Tags:       true,
			LFS:        true,
		},
	}

	err := p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}

	testFile, err := os.Stat(path.Join(dir, "100k-test.bin"))
	if err != nil {
		t.Errorf("Exec unable to get info on test file")
	}

	if testFile.Size() != wantFileSizeBytes {
		t.Errorf("Exec resulted in unexpected file size in repo object - want %d, got %d", wantFileSizeBytes, testFile.Size())
	}
}

func TestGit_Plugin_Validate(t *testing.T) {
	// setup types
	p := &Plugin{
		Build: &Build{
			Branch: "main",
			Path:   "/home/go-vela_vela-git-test_1",
			Ref:    "refs/heads/main",
			Sha:    "ee1e671529ad86a11ed628a04b37829e71783682",
		},
		Netrc: &Netrc{
			Machine:  "github.com",
			Username: "octocat",
			Password: "superSecretPassword",
		},
		Repo: &Repo{
			Remote:     "https://github.com/go-vela/vela-git-test.git",
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
			Remote:     "https://github.com/go-vela/vela-git-test.git",
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
			Path: "/home/go-vela_vela-git-test_1",
			Ref:  "refs/heads/main",
			Sha:  "ee1e671529ad86a11ed628a04b37829e71783682",
		},
		Netrc: &Netrc{},
		Repo: &Repo{
			Remote:     "https://github.com/go-vela/vela-git-test.git",
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
			Path: "/home/go-vela_vela-git-test_1",
			Ref:  "refs/heads/main",
			Sha:  "ee1e671529ad86a11ed628a04b37829e71783682",
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
