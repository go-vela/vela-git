// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/mail"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"

	_ "github.com/joho/godotenv/autoload"

	"github.com/go-vela/vela-git/version"
)

//nolint:funlen // extensive CLI flag configuration makes function long but readable
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
	cmd := &cli.Command{
		Name:      "vela-git",
		Usage:     "Vela Git plugin for cloning repositories",
		Copyright: "Copyright 2019 Target Brands, Inc. All rights reserved.",
		Authors: []any{
			&mail.Address{
				Name:    "Vela Admins",
				Address: "vela@target.com",
			},
		},
		// Plugin Metadata
		Version: v.Semantic(),
		Action:  run,
	}

	// Plugin Flags
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "log.level",
			Value: "info",
			Usage: "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_LOG_LEVEL"),
				cli.EnvVar("GIT_LOG_LEVEL"),
				cli.File("/vela/parameters/git/log_level"),
				cli.File("/vela/secrets/git/log_level"),
			),
		},

		// Build Flags
		&cli.StringFlag{
			Name:  "build.branch",
			Value: "main",
			Usage: "the repo branch for the build used during git init",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_BRANCH"),
				cli.EnvVar("GIT_BRANCH"),
				cli.EnvVar("VELA_PULL_REQUEST_SOURCE"),
				cli.EnvVar("VELA_BUILD_BRANCH"),
				cli.File("/vela/parameters/git/branch"),
				cli.File("/vela/secrets/git/branch"),
			),
		},
		&cli.StringFlag{
			Name:  "build.sha",
			Usage: "commit sha to clone from the repo",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_SHA"),
				cli.EnvVar("GIT_SHA"),
				cli.EnvVar("VELA_BUILD_COMMIT"),
				cli.File("/vela/parameters/git/sha"),
				cli.File("/vela/secrets/git/sha"),
			),
		},
		&cli.StringFlag{
			Name:  "build.path",
			Usage: "local path to clone the repo to",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_PATH"),
				cli.EnvVar("GIT_PATH"),
				cli.EnvVar("VELA_BUILD_WORKSPACE"),
				cli.File("/vela/parameters/git/path"),
				cli.File("/vela/secrets/git/path"),
			),
		},
		&cli.StringFlag{
			Name:  "build.ref",
			Value: "refs/heads/main",
			Usage: "commit reference to clone from the repo",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_REF"),
				cli.EnvVar("GIT_REF"),
				cli.EnvVar("VELA_BUILD_REF"),
				cli.File("/vela/parameters/git/ref"),
				cli.File("/vela/secrets/git/ref"),
			),
		},
		&cli.StringFlag{
			Name:  "build.depth",
			Value: "100",
			Usage: "enables fetching the repository with the specified depth",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_DEPTH"),
				cli.EnvVar("GIT_DEPTH"),
				cli.File("/vela/parameters/git/depth"),
				cli.File("/vela/secrets/git/depth"),
			),
		},

		// Netrc Flags
		&cli.StringFlag{
			Name:  "netrc.machine",
			Value: "github.com",
			Usage: "remote machine name to communicate with",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_MACHINE"),
				cli.EnvVar("GIT_MACHINE"),
				cli.EnvVar("VELA_NETRC_MACHINE"),
				cli.File("/vela/parameters/git/machine"),
				cli.File("/vela/secrets/git/machine"),
			),
		},
		&cli.StringFlag{
			Name:  "netrc.username",
			Usage: "user name for communication with the remote machine",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_USERNAME"),
				cli.EnvVar("GIT_USERNAME"),
				cli.EnvVar("VELA_NETRC_USERNAME"),
				cli.File("/vela/parameters/git/username"),
				cli.File("/vela/secrets/git/username"),
			),
		},
		&cli.StringFlag{
			Name:  "netrc.password",
			Usage: "password for communication with the remote machine",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_PASSWORD"),
				cli.EnvVar("GIT_PASSWORD"),
				cli.EnvVar("VELA_NETRC_PASSWORD"),
				cli.File("/vela/parameters/git/password"),
				cli.File("/vela/secrets/git/password"),
			),
		},

		// Repo Flags
		&cli.StringFlag{
			Name:  "repo.remote",
			Usage: "the remote (clone URL) for the repo being cloned",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_REMOTE"),
				cli.EnvVar("GIT_REMOTE"),
				cli.EnvVar("VELA_REPO_CLONE"),
				cli.File("/vela/parameters/git/remote"),
				cli.File("/vela/secrets/git/remote"),
			),
		},
		&cli.BoolFlag{
			Name:  "repo.submodules",
			Usage: "enables fetching submodules for the repo being cloned",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_SUBMODULES"),
				cli.EnvVar("GIT_SUBMODULES"),
				cli.File("/vela/parameters/git/submodules"),
				cli.File("/vela/secrets/git/submodules"),
			),
		},
		&cli.BoolFlag{
			Name:  "repo.tags",
			Usage: "enables fetching tags for the repo being cloned",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_TAGS"),
				cli.EnvVar("GIT_TAGS"),
				cli.File("/vela/parameters/git/tags"),
				cli.File("/vela/secrets/git/tags"),
			),
		},
		&cli.BoolFlag{
			Name:  "repo.lfs",
			Usage: "enables resolving LFS objects",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_LFS"),
				cli.EnvVar("GIT_LFS"),
				cli.File("/vela/parameters/git/lfs"),
				cli.File("/vela/secrets/git/lfs"),
			),
		},
	}

	err = cmd.Run(context.Background(), os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(ctx context.Context, c *cli.Command) error {
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
			LFS:        c.Bool("repo.lfs"),
		},
	}

	// validate the plugin
	err := p.Validate()
	if err != nil {
		return err
	}

	// execute the plugin
	return p.Exec(ctx)
}
