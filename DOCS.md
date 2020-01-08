## Description

**NOTE: This plugin is automatically injected into your pipeline for the source repository.**

This plugin enables you to clone repositories in a Vela pipeline to your build workspace.

Source Code: https://github.com/go-vela/vela-git
Registry: https://hub.docker.com/r/target/vela-git

## Usage

**NOTE: This plugin is automatically injected into your pipeline for the source repository.**

Sample of cloning a repository:

```yaml
steps:
  - name: clone_helloworld
    image: target/vela-git:v0.3.0
    parameters:
      path: hello-world
      ref: refs/heads/master
      remote: https://github.com/octocat/hello-world.git
      sha: 7fd1a60b01f91b314f59955a4e4d4e80d8edf11d
```

Sample of cloning a repository with submodules:

```diff
steps:
  - name: clone_helloworld
    image: target/vela-git:v0.3.0
    parameters:
      path: hello-world
      ref: refs/heads/master
      remote: https://github.com/octocat/hello-world.git
      sha: 7fd1a60b01f91b314f59955a4e4d4e80d8edf11d
+     submodules: true
```

Sample of cloning a repository with tags:

```diff
steps:
  - name: clone_helloworld
    image: target/vela-git:v0.3.0
    parameters:
      path: hello-world
      ref: refs/heads/master
      remote: https://github.com/octocat/hello-world.git
      sha: 7fd1a60b01f91b314f59955a4e4d4e80d8edf11d
+     tags: true
```

## Secrets

**NOTE: Users should refrain from configuring sensitive information in your pipeline in plain text.**

You can use Vela secrets to substitute sensitive values at runtime:

```diff
steps:
  - name: clone_helloworld
    image: target/vela-git:v0.3.0
+   secrets: [ vela_netrc_username, vela_netrc_password ]
    parameters:
-     netrc_username: octocat
-     netrc_password: superSecretPassword
      path: /home/octocat_hello-world_1
      ref: refs/heads/master
      remote: https://github.com/octocat/hello-world.git
      sha: 7fd1a60b01f91b314f59955a4e4d4e80d8edf11d
```

## Parameters

The following parameters are used to configure the image:

| Name         | Description                       | Required | Default             |
| ------------ | --------------------------------- | -------- | ------------------- |
| `log_level`  | set the log level for the plugin  | `true`   | `info`              |
| `path`       | local path to clone repository to | `true`   | `N/A`               |
| `ref`        | reference generated for commit    | `true`   | `refs/heads/master` |
| `remote`     | full url for cloning repository   | `true`   | `N/A`               |
| `sha`        | SHA-1 hash generated for commit   | `true`   | `N/A`               |
| `submodules` | enables fetching of submodules    | `false`  | `N/A`               |
| `tags`       | enables fetching of tags          | `false`  | `N/A`               |

## Template

COMING SOON!

## Troubleshooting

Below are a list of common problems and how to solve them:
