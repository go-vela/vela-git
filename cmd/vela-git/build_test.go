// SPDX-License-Identifier: Apache-2.0

package main

import "testing"

func TestGit_Build_Validate(t *testing.T) {
	tests := []struct {
		name    string
		failure bool
		build   *Build
	}{
		{
			name:    "success",
			failure: false,
			build: &Build{
				Branch: "main",
				Path:   "/home/go-vela_vela-git-test_1",
				Ref:    "refs/heads/main",
				Sha:    "ee1e671529ad86a11ed628a04b37829e71783682",
			},
		},
		{
			name:    "failure with no branch",
			failure: true,
			build: &Build{
				Path: "/home/go-vela_vela-git-test_1",
				Ref:  "refs/heads/main",
				Sha:  "ee1e671529ad86a11ed628a04b37829e71783682",
			},
		},
		{
			name:    "failure with no path",
			failure: true,
			build: &Build{
				Branch: "main",
				Ref:    "refs/heads/main",
				Sha:    "ee1e671529ad86a11ed628a04b37829e71783682",
			},
		},
		{
			name:    "failure with no ref",
			failure: true,
			build: &Build{
				Branch: "main",
				Path:   "/home/go-vela_vela-git-test_1",
				Sha:    "ee1e671529ad86a11ed628a04b37829e71783682",
			},
		},
		{
			name:    "failure with no sha",
			failure: true,
			build: &Build{
				Branch: "main",
				Path:   "/home/go-vela_vela-git-test_1",
				Ref:    "refs/heads/main",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.build.Validate()

			if test.failure {
				if err == nil {
					t.Errorf("Validate for %s should have returned err", test.name)
				}

				return
			}

			if err != nil {
				t.Errorf("Validate for %s returned err: %v", test.name, err)
			}
		})
	}
}
