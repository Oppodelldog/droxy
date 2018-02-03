[![Go Report Card](https://goreportcard.com/badge/github.com/Oppodelldog/docker-proxy-command)](https://goreportcard.com/report/github.com/Oppodelldog/docker-proxy-command)

# The idea
This tool should help you in creating variants of commands that proxy execution into docker containers.

**WHY?**  
When working on many different projects that require different software or different versions you might want
to have all those tools from one hand, maybe a configuration file.

**HOW?**
1. Commit dev-tools to your project by a config file
2. Bootstrap all command executables from that config file
3. Just use your new tools

## the config (toml)

```TOML
    Version="1"

    [[command]]
      name = "basic command"  # name of the command
      isTemplate = true       # this command can be used as a template, no command will be created
      addGroups = true        # add current systems groups
      impersonate = true      # use executing user and group for execution in the container
      workDir = "/app"        # define working directory
      removeContainer=true    # remove container after command has finished
      isInteractive=true      # enable interaction with the called command

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

#### Pros and Contras

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


**symlinks**  
**Pro:** low disc space consumption  
**Con:** config will only be found when defined by environment variable  
  
**hardlinks**  
**Pro:** config will be found  
**Con:** ?  
  
**clones**  
**Pro:** config will be found  
**Con:** needs more disc space  
  

### bootstrap containers (?!)
    docker-proxy prepare

all required images are pulled
