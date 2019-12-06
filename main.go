// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "git"
	app.Usage = "Git plugin for cloning repositories"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "remote",
			Usage:  "git remote url",
			EnvVar: "PARAMETER_REMOTE,REPOSITORY_CLONE",
		},
		cli.StringFlag{
			Name:   "path",
			Usage:  "git clone path",
			EnvVar: "PARAMETER_PATH,BUILD_WORKSPACE",
		},
		cli.StringFlag{
			Name:   "commit",
			Usage:  "git commit sha",
			EnvVar: "PARAMETER_COMMIT,BUILD_COMMIT",
		},
		cli.StringFlag{
			Name:   "ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "PARAMETER_REF,BUILD_REF",
		},
		cli.BoolFlag{
			Name:   "tags",
			Usage:  "git fetch tags",
			EnvVar: "PARAMETER_TAGS",
		},
		cli.BoolFlag{
			Name:   "submodules",
			Usage:  "git update submodules",
			EnvVar: "PARAMETER_SUBMODULES",
		},
		cli.StringFlag{
			Name:   "netrc-machine",
			Usage:  "netrc machine",
			EnvVar: "VELA_NETRC_MACHINE",
		},
		cli.StringFlag{
			Name:   "netrc-username",
			Usage:  "netrc username",
			EnvVar: "VELA_NETRC_USERNAME",
		},
		cli.StringFlag{
			Name:   "netrc-password",
			Usage:  "netrc password",
			EnvVar: "VELA_NETRC_PASSWORD",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}

}

func run(c *cli.Context) error {
	plugin := Plugin{
		Remote:        c.String("remote"),
		Path:          c.String("path"),
		CommitSha:     c.String("commit"),
		CommitRef:     c.String("ref"),
		Tags:          c.Bool("tags"),
		Submodules:    c.Bool("submodules"),
		NetrcMachine:  c.String("netrc-machine"),
		NetrcUsername: c.String("netrc-username"),
		NetrcPassword: c.String("netrc-password"),
	}

	return plugin.Exec()
}
