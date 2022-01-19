// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import "testing"

func TestGit_Repo_Validate(t *testing.T) {
	// setup types
	r := &Repo{
		Remote:     "https://github.com/octocat/hello-world.git",
		Submodules: false,
		Tags:       false,
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
	}

	err := r.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
