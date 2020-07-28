// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := cli.NewApp()

	// Plugin Information

	app.Name = "vela-git"
	app.HelpName = "vela-git"
	app.Usage = "Vela Git plugin for cloning repositories"
	app.Copyright = "Copyright (c) 2020 Target Brands, Inc. All rights reserved."
	app.Authors = []*cli.Author{
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

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_LOG_LEVEL", "VELA_LOG_LEVEL", "GIT_LOG_LEVEL"},
			FilePath: string("/vela/parameters/log_level,/vela/secrets/downstream/log_level"),
			Name:     "log.level",
			Usage:    "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:    "info",
		},

		// Build Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_SHA", "BUILD_COMMIT"},
			FilePath: string("/vela/parameters/build/sha,/vela/secrets/downstream/build/sha"),
			Name:     "build.sha",
			Usage:    "git commit sha",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_PATH", "BUILD_WORKSPACE"},
			FilePath: string("/vela/parameters/build/path,/vela/secrets/downstream/build/path"),
			Name:     "build.path",
			Usage:    "git clone path",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_REF", "BUILD_REF"},
			FilePath: string("/vela/parameters/build/ref,/vela/secrets/downstream/build/ref"),
			Name:     "build.ref",
			Usage:    "git commit ref",
			Value:    "refs/heads/master",
		},

		// Netrc Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_NETRC_MACHINE", "VELA_NETRC_MACHINE"},
			FilePath: string("/vela/parameters/netrc/machine,/vela/secrets/downstream/netrc/machine"),
			Name:     "netrc.machine",
			Usage:    "remote machine name to communicate with",
			Value:    "github.com",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_NETRC_USERNAME", "VELA_NETRC_USERNAME", "GIT_USERNAME"},
			FilePath: string("/vela/parameters/netrc/username,/vela/secrets/downstream/netrc/username"),
			Name:     "netrc.username",
			Usage:    "user name for communication with the remote machine",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_NETRC_PASSWORD", "VELA_NETRC_PASSWORD", "GIT_PASSWORD"},
			FilePath: string("/vela/parameters/netrc/password,/vela/secrets/downstream/netrc/password"),
			Name:     "netrc.password",
			Usage:    "password for communication with the remote machine",
		},

		// Repo Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_REMOTE", "REPOSITORY_CLONE"},
			FilePath: string("/vela/parameters/repo/remote,/vela/secrets/downstream/repo/remote"),
			Name:     "repo.remote",
			Usage:    "git remote url",
		},
		&cli.BoolFlag{
			EnvVars:  []string{"PARAMETER_SUBMODULES"},
			FilePath: string("/vela/parameters/repo/submodules,/vela/secrets/downstream/repo/submodules"),
			Name:     "repo.submodules",
			Usage:    "git update submodules",
		},
		&cli.BoolFlag{
			EnvVars:  []string{"PARAMETER_TAGS"},
			FilePath: string("/vela/parameters/repo/tags,/vela/secrets/downstream/repo/tags"),
			Name:     "repo.tags",
			Usage:    "git fetch tags",
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

	logrus.WithFields(logrus.Fields{
		"code":     "https://github.com/go-vela/vela-git",
		"docs":     "https://go-vela.github.io/docs/plugins/registry/git",
		"registry": "https://hub.docker.com/r/target/vela-git",
	}).Info("Vela Git Plugin")

	// create the plugin
	p := &Plugin{
		// build configuration
		Build: &Build{
			Path: c.String("build.path"),
			Ref:  c.String("build.ref"),
			Sha:  c.String("build.sha"),
		},
		// netrc configuration
		Netrc: &Netrc{
			Machine:  c.String("netrc.machine"),
			Username: c.String("netrc.username"),
			Password: c.String("netrc.password"),
		},
		// repo configuration
		Repo: &Repo{
			Remote:     c.String("repo.remote"),
			Submodules: c.Bool("repo.submodules"),
			Tags:       c.Bool("repo.tags"),
		},
	}

	// validate the plugin
	err := p.Validate()
	if err != nil {
		return err
	}

	// execute the plugin
	return p.Exec()
}
