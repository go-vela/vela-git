// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := cli.NewApp()

	// Plugin Information

	app.Name = "vela-git"
	app.HelpName = "vela-git"
	app.Usage = "Vela Git plugin for cloning repositories"
	app.Copyright = "Copyright (c) 2019 Target Brands, Inc. All rights reserved."
	app.Authors = []cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	// Plugin Metadata

	app.Compiled = time.Now()
	app.Action = run

	// Plugin Flags

	app.Flags = []cli.Flag{

		// Default Flags

		cli.StringFlag{
			EnvVar: "PARAMETER_COMMIT,BUILD_COMMIT",
			Name:   "commit",
			Usage:  "git commit sha",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_PATH,BUILD_WORKSPACE",
			Name:   "path",
			Usage:  "git clone path",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_REF,BUILD_REF",
			Name:   "ref",
			Usage:  "git commit ref",
			Value:  "refs/heads/master",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_REMOTE,REPOSITORY_CLONE",
			Name:   "remote",
			Usage:  "git remote url",
		},

		// Netrc Flags
		//
		// https://www.gnu.org/software/inetutils/manual/html_node/The-_002enetrc-file.html

		cli.StringFlag{
			EnvVar: "PARAMETER_NETRC_MACHINE,VELA_NETRC_MACHINE",
			Name:   "netrc.machine",
			Usage:  "remote machine name to communicate with",
			Value:  "github.com",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_NETRC_USERNAME,VELA_NETRC_USERNAME",
			Name:   "netrc.username",
			Usage:  "user name for communication with the remote machine",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_NETRC_PASSWORD,VELA_NETRC_PASSWORD",
			Name:   "netrc.password",
			Usage:  "password for communication with the remote machine",
		},

		// Optional Flags

		cli.BoolFlag{
			EnvVar: "PARAMETER_SUBMODULES",
			Name:   "submodules",
			Usage:  "git update submodules",
		},
		cli.BoolFlag{
			EnvVar: "PARAMETER_TAGS",
			Name:   "tags",
			Usage:  "git fetch tags",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

// run executes the CLI
func run(c *cli.Context) error {
	plugin := Plugin{
		// default arguments
		Default: &Default{
			Path:   c.String("path"),
			Ref:    c.String("ref"),
			Remote: c.String("remote"),
			Sha:    c.String("commit"),
		},
		// netrc arguments
		Netrc: &Netrc{
			Machine:  c.String("netrc.machine"),
			Username: c.String("netrc.username"),
			Password: c.String("netrc.password"),
		},
		// optional arguments
		Optional: &Optional{
			Submodules: c.Bool("submodules"),
			Tags:       c.Bool("tags"),
		},
	}

	return plugin.Exec()
}
