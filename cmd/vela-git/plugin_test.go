package main

import (
	"testing"

	"github.com/spf13/afero"
)

func TestGit_writeNetrc(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	n := &Netrc{
		Machine:  "github.com",
		Username: "octocat",
		Password: "superSecretPassword",
	}

	err := writeNetrc(n.Machine, n.Username, n.Password)
	if err != nil {
		t.Errorf("writeNetrc returned err: %v", err)
	}
}
