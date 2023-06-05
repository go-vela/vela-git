// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

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
				Path: "/home/octocat_hello-world_1",
				Ref:  "refs/heads/master",
				Sha:  "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
			},
		},
		{
			name:    "failure with no branch",
			failure: true,
			build: &Build{
				Path: "/home/octocat_hello-world_1",
				Ref:  "refs/heads/master",
				Sha:  "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
			},
		},
		{
			name:    "failure with no path",
			failure: true,
			build: &Build{
				Branch: "master",
				Ref:    "refs/heads/master",
				Sha:    "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
			},
		},
		{
			name:    "failure with no ref",
			failure: true,
			build: &Build{
				Branch: "master",
				Path:   "/home/octocat_hello-world_1",
				Sha:    "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
			},
		},
		{
			name:    "failure with no sha",
			failure: true,
			build: &Build{
				Branch: "master",
				Path:   "/home/octocat_hello-world_1",
				Ref:    "refs/heads/master",
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
