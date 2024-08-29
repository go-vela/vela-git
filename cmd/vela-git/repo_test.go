// SPDX-License-Identifier: Apache-2.0

package main

import "testing"

func TestGit_Repo_Validate(t *testing.T) {
	// setup types
	r := &Repo{
		Remote:     "https://github.com/go-vela/vela-git-test.git",
		Submodules: false,
		Tags:       false,
		LFS:        false,
	}

	err := r.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGit_Repo_Validate_NoRemote(t *testing.T) {
	// setup types
	r := &Repo{
		Submodules: false,
		Tags:       false,
		LFS:        false,
	}

	err := r.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
