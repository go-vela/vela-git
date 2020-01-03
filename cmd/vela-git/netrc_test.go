package main

import "testing"

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
