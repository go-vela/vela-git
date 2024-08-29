## Description

> **NOTE:** This plugin is automatically injected into your pipeline for the source repository.

This plugin enables you to clone repositories in a Vela pipeline to your build workspace.

Source Code: https://github.com/go-vela/vela-git

Registry: https://hub.docker.com/r/target/vela-git

## Usage

> **NOTE:**
>
> Users should refrain from using latest as the tag for the Docker image.
>
> It is recommended to use a semantically versioned tag instead.

Sample of cloning a repository:

```yaml
steps:
  - name: clone_vela-git-test
    image: target/vela-git:latest
    pull: always
    parameters:
      path: vela-git-test
      ref: refs/heads/main
      remote: https://github.com/go-vela/vela-git-test.git
      sha: ee1e671529ad86a11ed628a04b37829e71783682
```

Sample of cloning a repository with submodules:

```diff
steps:
  - name: clone_vela-git-test
    image: target/vela-git:latest
    pull: always
    parameters:
      path: vela-git-test
      ref: refs/heads/main
      remote: https://github.com/go-vela/vela-git-test.git
      sha: ee1e671529ad86a11ed628a04b37829e71783682
+     submodules: true
```

Sample of cloning a repository with tags:

```diff
steps:
  - name: clone_vela-git-test
    image: target/vela-git:latest
    pull: always
    parameters:
      path: vela-git-test
      ref: refs/heads/main
      remote: https://github.com/go-vela/vela-git-test.git
      sha: ee1e671529ad86a11ed628a04b37829e71783682
+     tags: true
```

Sample of cloning a repository and resolving LFS objects:

```diff
steps:
  - name: clone_vela-git-test
    image: target/vela-git:latest
    pull: always
    parameters:
      path: vela-git-test
      ref: refs/heads/main
      remote: https://github.com/go-vela/vela-git-test.git
      sha: ee1e671529ad86a11ed628a04b37829e71783682
+     lfs: true
```

## Secrets

> **NOTE:** Users should refrain from configuring sensitive information in your pipeline in plain text.

### Internal

Users can use [Vela internal secrets](https://go-vela.github.io/docs/tour/secrets/) to substitute these sensitive values at runtime:

```diff
steps:
  - name: clone_vela-git-test
    image: target/vela-git:latest
    pull: always
+   secrets: [ git_username, git_password ]
    parameters:
-     username: octocat
-     password: superSecretPassword
      path: /home/go-vela_vela-git-test_1
      ref: refs/heads/master
      remote: https://github.com/go-vela/vela-git-test.git
      sha: ee1e671529ad86a11ed628a04b37829e71783682
```

> This example will add the secrets to the `clone_vela-git-test` step as environment variables:
>
> * `GIT_USERNAME=<value>`
> * `GIT_PASSWORD=<value>`

### External

The plugin accepts the following files for authentication:

| Parameter  | Volume Configuration                                          |
| ---------- | ------------------------------------------------------------- |
| `password` | `/vela/parameters/git/password`, `/vela/secrets/git/password` |
| `username` | `/vela/parameters/git/username`, `/vela/secrets/git/username` |

Users can use [Vela external secrets](https://go-vela.github.io/docs/concepts/pipeline/secrets/origin/) to substitute these sensitive values at runtime:

```diff
steps:
  - name: clone_vela-git-test
    image: target/vela-git:latest
    pull: always
    parameters:
-     username: octocat
-     password: superSecretPassword
      path: /home/go-vela_vela-git-test_1
      ref: refs/heads/main
      remote: https://github.com/go-vela/vela-git-test.git
      sha: ee1e671529ad86a11ed628a04b37829e71783682
```

> This example will read the secret values in the volume stored at `/vela/secrets/`

## Parameters

> **NOTE:**
>
> The plugin supports reading all parameters via environment variables or files.
>
> Any values set from a file take precedence over values set from the environment.

The following parameters are used to configure the image:

| Name         | Description                             | Required | Default             | Environment Variables                                                                   |
|--------------|-----------------------------------------| -------- |---------------------|-----------------------------------------------------------------------------------------|
| `branch`     | initial branch to clone from repository | `true`   | `master`            | `PARAMETER_BRANCH`<br>`GIT_BRANCH`<br>`VELA_PULL_REQUEST_SOURCE`<br>`VELA_BUILD_BRANCH` |
| `log_level`  | set the log level for the plugin        | `true`   | `info`              | `PARAMETER_LOG_LEVEL`<br>`GIT_LOG_LEVEL`                                                |
| `machine`    | machine name to communicate with        | `true`   | `github.com`        | `PARAMETER_MACHINE`<br>`GIT_MACHINE`<br>`VELA_NETRC_MACHINE`                            |
| `password`   | password for authentication             | `true`   | **set by Vela**     | `PARAMETER_PASSWORD`<br>`GIT_PASSWORD`<br>`VELA_NETRC_PASSWORD`                         |
| `username`   | user name for authentication            | `true`   | **set by Vela**     | `PARAMETER_USERNAME`<br>`GIT_USERNAME`<br>`VELA_NETRC_USERNAME`                         |
| `path`       | local path to clone repository to       | `true`   | **set by Vela**     | `PARAMETER_PATH`<br>`GIT_PATH`<br>`VELA_BUILD_WORKSPACE`                                |
| `ref`        | reference generated for commit          | `true`   | `refs/heads/master` | `PARAMETER_REF`<br>`GIT_REF`<br>`VELA_BUILD_REF`                                        |
| `remote`     | full url for cloning repository         | `true`   | **set by Vela**     | `PARAMETER_REMOTE`<br>`GIT_REMOTE`<br>`VELA_REPO_CLONE`                                 |
| `sha`        | SHA-1 hash generated for commit         | `true`   | **set by Vela**     | `PARAMETER_SHA`<br>`GIT_SHA`<br>`VELA_BUILD_COMMIT`                                     |
| `submodules` | enables fetching of submodules          | `false`  | `false`             | `PARAMETER_SUBMODULES`<br>`GIT_SUBMODULES`                                              |
| `tags`       | enables fetching of tags                | `false`  | `false`             | `PARAMETER_TAGS`<br>`GIT_TAGS`                                                          |
| `depth`      | enables fetching with a specific depth  | `false`  | `100`               | `PARAMETER_DEPTH`<br>`GIT_DEPTH`                                                        |

## Template

COMING SOON!

## Troubleshooting

You can start troubleshooting this plugin by tuning the level of logs being displayed:

```diff
steps:
  - name: clone_vela-git-test
    image: target/vela-git:latest
    pull: always
    parameters:
+     log_level: trace
      path: vela-git-test
      ref: refs/heads/main
      remote: https://github.com/go-vela/vela-git-test.git
      sha: ee1e671529ad86a11ed628a04b37829e71783682
```

Below are a list of common problems and how to solve them:
