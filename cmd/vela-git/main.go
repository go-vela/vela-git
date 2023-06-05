// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-vela/vela-git/version"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// capture application version information
	v := version.New()

	// serialize the version information as pretty JSON
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.Fatal(err)
	}

	// output the version information to stdout
	fmt.Fprintf(os.Stdout, "%s\n", string(bytes))

	// create new CLI application
	app := cli.NewApp()

	// Plugin Information

	app.Name = "vela-git"
	app.HelpName = "vela-git"
	app.Usage = "Vela Git plugin for cloning repositories"
	app.Copyright = "Copyright (c) 2023 Target Brands, Inc. All rights reserved."
	app.Authors = []*cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	// Plugin Metadata

	app.Action = run
	app.Compiled = time.Now()
	app.Version = v.Semantic()

	// Plugin Flags

	app.Flags = []cli.Flag{

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_LOG_LEVEL", "GIT_LOG_LEVEL"},
			FilePath: "/vela/parameters/git/log_level,/vela/secrets/git/log_level",
			Name:     "log.level",
			Usage:    "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:    "info",
		},

		// Build Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_BRANCH", "GIT_BRANCH", "VELA_PULL_REQUEST_SOURCE", "VELA_BUILD_BRANCH"},
			FilePath: "/vela/parameters/git/branch,/vela/secrets/git/branch",
			Name:     "build.branch",
			Usage:    "the repo branch for the build used during git init",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_SHA", "GIT_SHA", "VELA_BUILD_COMMIT"},
			FilePath: "/vela/parameters/git/sha,/vela/secrets/git/sha",
			Name:     "build.sha",
			Usage:    "commit sha to clone from the repo",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_PATH", "GIT_PATH", "VELA_BUILD_WORKSPACE"},
			FilePath: "/vela/parameters/git/path,/vela/secrets/git/path",
			Name:     "build.path",
			Usage:    "local path to clone the repo to",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_REF", "GIT_REF", "VELA_BUILD_REF"},
			FilePath: "/vela/parameters/git/ref,/vela/secrets/git/ref",
			Name:     "build.ref",
			Usage:    "commit reference to clone from the repo",
			Value:    "refs/heads/master",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_DEPTH", "GIT_DEPTH"},
			FilePath: "/vela/parameters/git/depth,/vela/secrets/git/depth",
			Name:     "build.depth",
			Usage:    "enables fetching the repository with the specified depth",
			Value:    "100",
		},

		// Netrc Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_MACHINE", "GIT_MACHINE", "VELA_NETRC_MACHINE"},
			FilePath: "/vela/parameters/git/machine,/vela/secrets/git/machine",
			Name:     "netrc.machine",
			Usage:    "remote machine name to communicate with",
			Value:    "github.com",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_USERNAME", "GIT_USERNAME", "VELA_NETRC_USERNAME"},
			FilePath: "/vela/parameters/git/username,/vela/secrets/git/username",
			Name:     "netrc.username",
			Usage:    "user name for communication with the remote machine",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_PASSWORD", "GIT_PASSWORD", "VELA_NETRC_PASSWORD"},
			FilePath: "/vela/parameters/git/password,/vela/secrets/git/password",
			Name:     "netrc.password",
			Usage:    "password for communication with the remote machine",
		},

		// Repo Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_REMOTE", "GIT_REMOTE", "VELA_REPO_CLONE"},
			FilePath: "/vela/parameters/git/remote,/vela/secrets/git/remote",
			Name:     "repo.remote",
			Usage:    "the remote (clone URL) for the repo being cloned",
		},
		&cli.BoolFlag{
			EnvVars:  []string{"PARAMETER_SUBMODULES", "GIT_SUBMODULES"},
			FilePath: "/vela/parameters/git/submodules,/vela/secrets/git/submodules",
			Name:     "repo.submodules",
			Usage:    "enables fetching submodules for the repo being cloned",
		},
		&cli.BoolFlag{
			EnvVars:  []string{"PARAMETER_TAGS", "GIT_TAGS"},
			FilePath: "/vela/parameters/git/tags,/vela/secrets/git/tags",
			Name:     "repo.tags",
			Usage:    "enables fetching tags for the repo being cloned",
		},
	}

	err = app.Run(os.Args)
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
		"docs":     "https://go-vela.github.io/docs/plugins/registry/pipeline/git",
		"registry": "https://hub.docker.com/r/target/vela-git",
	}).Info("Vela Git Plugin")

	// create the plugin
	p := &Plugin{
		// build configuration
		Build: &Build{
			Branch: c.String("build.branch"),
			Path:   c.String("build.path"),
			Ref:    c.String("build.ref"),
			Sha:    c.String("build.sha"),
			Depth:  c.String("build.depth"),
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
