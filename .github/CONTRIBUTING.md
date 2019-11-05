# Contributing

We'd love to accept your contributions to this project! There are just a few guidelines you need to follow.

## Feature Requests

Feature Requests should be opened up as [Issues](/issues/new/) on this repository!

## Issues

[Issues](/issues/new/) are always welcome!

## Pull Requests

**NOTE: We recommend you start by opening a new issue describing the bug or feature you're intending to fix. Even if you think it's relatively minor, it's helpful to know what people are working on.**

We are always welcome to new PRs! You can follow the below guide for learning how you can contribute to the project!

## Getting Started

### Prerequisites

* [Review the commit guide we follow](https://chris.beams.io/posts/git-commit/#seven-rules) - ensure your commits follow our standards
* [Docker](https://docs.docker.com/install/) - building block for local development
* [Docker Compose](https://docs.docker.com/compose/install/) - start up local development
* [Golang](https://golang.org/dl/) - for source code and [dependency management](https://github.com/golang/go/wiki/Modules)
* [Make](https://www.gnu.org/software/make/) - start up local development

### Setup

* [Fork](/fork) this repository

* Clone this repository to your workstation:

```bash
# Clone the project
git clone git@github.com:go-vela/vela-git.git $HOME/go-vela/vela-git
```

* Navigate to the repository code:

```bash
# Change into the project directory
cd $HOME/go-vela/vela-git
```

* Point the original code at your fork:

```bash
# Add a remote branch pointing to your fork
git remote add fork https://github.com/your_fork/vela-git
```

### Running Locally

* Navigate to the repository code:

```bash
# Change into the project directory
cd $HOME/go-vela/vela-git
```

* Build the repository code:

```bash
# Build the code with `make`
make build
```

* Run the repository code:

```bash
# Run the code with `make`
make run
```

### Development

coming soon!
