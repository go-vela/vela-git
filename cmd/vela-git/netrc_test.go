// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"

	"github.com/spf13/afero"
)

func TestGit_Netrc_Validate(t *testing.T) {
	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Username: "octocat",
		Password: "superSecretPassword",
	}

	err := n.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestGit_Netrc_Validate_NoMachine(t *testing.T) {
	// setup types
	n := &Netrc{
		Username: "octocat",
		Password: "superSecretPassword",
	}

	err := n.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestGit_Netrc_Validate_NoUsername(t *testing.T) {
	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Password: "superSecretPassword",
	}

	err := n.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestGit_Netrc_Validate_NoPassword(t *testing.T) {
	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Username: "octocat",
	}

	err := n.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestGit_Netrc_Write(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Username: "octocat",
		Password: "superSecretPassword",
	}

	err := n.Write()
	if err != nil {
		t.Errorf("Write returned err: %v", err)
	}
}

func TestGit_Netrc_Write_Error(t *testing.T) {
	// setup filesystem
	appFS = afero.NewReadOnlyFs(afero.NewMemMapFs())

	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Username: "octocat",
		Password: "superSecretPassword",
	}

	err := n.Write()
	if err == nil {
		t.Errorf("Write should have returned err")
	}
}

func TestGit_Netrc_Write_NoMachine(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	n := &Netrc{
		Username: "octocat",
		Password: "superSecretPassword",
	}

	err := n.Write()
	if err != nil {
		t.Errorf("Write returned err: %v", err)
	}
}

func TestGit_Netrc_Write_NoUsername(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Password: "superSecretPassword",
	}

	err := n.Write()
	if err != nil {
		t.Errorf("Write returned err: %v", err)
	}
}

func TestGit_Netrc_Write_NoPassword(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Username: "octocat",
	}

	err := n.Write()
	if err != nil {
		t.Errorf("Write returned err: %v", err)
	}
}
