[![Go Report Card](https://goreportcard.com/badge/github.com/Oppodelldog/droxy)](https://goreportcard.com/report/github.com/Oppodelldog/droxy) [![License](https://img.shields.io/badge/License-BSD--3-blue.svg)](https://raw.githubusercontent.com/Oppodelldog/droxy/master/LICENSE) [![Linux build](http://nulldog.de:12080/api/badges/Oppodelldog/droxy/status.svg)](http://nulldog.de:12080/Oppodelldog/droxy) [![Windows build](https://ci.appveyor.com/api/projects/status/qpe2889fbk1bw7lf/branch/master?svg=true)](https://ci.appveyor.com/project/Oppodelldog/droxy/branch/master) [![Coverage Status](https://coveralls.io/repos/github/Oppodelldog/droxy/badge.svg?branch=master)](https://coveralls.io/github/Oppodelldog/droxy?branch=master)
# Droxy
> create commands that proxy to docker

## The idea
This tool should help you in creating variants of commands that proxy execution into docker containers.

## Getting started
To get started you either download a precompiled binary from [releases](https://github.com/Oppodelldog/droxy/releases).
If you would like to build the tool from source code, read the contribution part of this document.

## Configuration
In the config file, you define the commands you want to create.
The config file must be named ```droxy.toml```.

> droxy.toml

```TOML
    Version="1"

    [[command]]
      name = "basic command"  # name of the command
      isTemplate = true       # this command can be used as a template, no command will be created
      addGroups = true        # add current systems groups  (linux only)
      impersonate = true      # use executing user and group for execution in the container (linux only)
      workDir = "/app"        # define working directory
      removeContainer=true    # remove container after command has finished
      isInteractive=true      # enable interaction with the called command
      RequireEnvVars=false    # if true, not defined env vars that are configured will lead to an error

      # volume mappings
      volumes = [
          "${HOME}:${HOME}",
          "${SSH_AUTH_SOCK}:/run/ssh.sock",
          "/etc/passwd:/etc/passwd:ro",
          "/etc/group:/etc/group:ro",
          "/run/docker.sock:/run/docker.sock"
      ]

      # environment variable mappings
      envvars = [
          "HOME:${HOME}",
          "SSH_AUTH_SOCK:/run/ssh.sock",
          "DOCKER_HOST=unix:///run/docker.sock"
      ]

    [[command]]
        template = "basic command"  # apply settings from template 'basic command' to this command
    	name = "php"                # name of the command which is created by calling 'docker-proxy symlinks'
    	entryPoint = "php"          # basic binary to execute inside the container
    	image = "php:7.1.13"        # docker image the container is run on

    [[command]]
        template = "basic command"
    	name = "phpstorm-php-unittest-integration"
    	entryPoint = "php"
    	image = "php:7.1.13"

    	# replace  127.0.0.1 with docker ip of host
        replaceArgs = [
            [
              "-dxdebug.remote_host=127.0.0.1",
              "-dxdebug.remote_host=172.17.0.1"
            ]
        ]

        # ensure xdebug will startup and communicate
        additionalArgs = ["-dxdebug.remote_autostart=1"]

```


### create commands
So the idea is to create "real" commands in form of binaries.
Why? Well for a bash user also bash files would apply, but when you are trying to trick an IDE to use such a bash
file as a real executable some may fail. So in general it should be a good idea to create real binaries.

**What are the options?**  
There are three options to create custom docker-proxy commands into a directory:  

* **symlinks**  
    this sub-command creates a symlink for every command defined in the config into the current directory
* **clones**  
    this sub-command creates copies of docker-proxy in the config into the current directory, renaming to the appropriate command name.
* **hardlinks**  
    this sub-command creates a hardlink for every command defined in the config into the current directory

```bash
docker-proxy symlinks
```
```bash
docker-proxy hardlinks
```
```bash
docker-proxy clones
```

use ```-f``` to force file creation. This will delete files with command names before creation.


## Contribute
Feature requests and pull requests are welcome.

### How to
Clone the repository into your go folder.
The path should look like this ```.../go/src/github.com/Oppodelldog/droxy.```

There are some make targets that help you with several tasks, first: ```setup```:

```shell
make setup
```
This installs necessary go tools to get further jobs done and installs vendors.

Now you are ready to go.
```make install``` builds and installs the **droxy** command in **.../go/bin**
so you can directly use or test it.

> If you have no make installed, execute the commands of the ```setup``` task manually.

