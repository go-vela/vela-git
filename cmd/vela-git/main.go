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

		cli.StringFlag{
			EnvVar: "PARAMETER_LOG_LEVEL,VELA_LOG_LEVEL,GIT_LOG_LEVEL",
			Name:   "log.level",
			Usage:  "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:  "info",
		},

		// Build Flags

		cli.StringFlag{
			EnvVar: "PARAMETER_SHA,BUILD_COMMIT",
			Name:   "build.sha",
			Usage:  "git commit sha",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_PATH,BUILD_WORKSPACE",
			Name:   "build.path",
			Usage:  "git clone path",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_REF,BUILD_REF",
			Name:   "build.ref",
			Usage:  "git commit ref",
			Value:  "refs/heads/master",
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

		// Repo Flags

		cli.StringFlag{
			EnvVar: "PARAMETER_REMOTE,REPOSITORY_CLONE",
			Name:   "repo.remote",
			Usage:  "git remote url",
		},
		cli.BoolFlag{
			EnvVar: "PARAMETER_SUBMODULES",
			Name:   "repo.submodules",
			Usage:  "git update submodules",
		},
		cli.BoolFlag{
			EnvVar: "PARAMETER_TAGS",
			Name:   "repo.tags",
			Usage:  "git fetch tags",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(c *cli.Context) error {
	// set the log level for the plugin
	switch c.String("log.level") {
	case "t", "trace", "Trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "w", "warn", "Warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "i", "info", "Info", "INFO":
		fallthrough
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	// validate the CLI configuration
	err := validate(c)
	if err != nil {
		return err
	}

	// create the plugin object
	p := Plugin{
		// build configuration
		Build: &Build{
			Path: c.String("build.path"),
			Ref:  c.String("build.ref"),
			Sha:  c.String("build.sha"),
		},
		// netrc arguments
		Netrc: &Netrc{
			Machine:  c.String("netrc.machine"),
			Username: c.String("netrc.username"),
			Password: c.String("netrc.password"),
		},
		Repo: &Repo{
			Remote:     c.String("repo.remote"),
			Submodules: c.Bool("repo.submodules"),
			Tags:       c.Bool("repo.tags"),
		},
	}

	// execute the plugin
	return p.Exec()
}
