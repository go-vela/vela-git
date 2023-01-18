// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import "testing"

func TestGit_Build_Validate(t *testing.T) {
	// setup types
	b := &Build{
		Path: "/home/octocat_hello-world_1",
		Ref:  "refs/heads/master",
		Sha:  "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
	}

	err := b.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGit_Build_Validate_NoPath(t *testing.T) {
	// setup types
	b := &Build{
		Ref: "refs/heads/master",
		Sha: "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
	}

	err := b.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestGit_Build_Validate_NoRef(t *testing.T) {
	// setup types
	b := &Build{
		Path: "/home/octocat_hello-world_1",
		Sha:  "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
	}

	err := b.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestGit_Build_Validate_NoSha(t *testing.T) {
	// setup types
	b := &Build{
		Path: "/home/octocat_hello-world_1",
		Ref:  "refs/heads/master",
	}

	err := b.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
