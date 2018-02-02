# WIP - proof of concept

# The idea
Create a tool that helps creating variants of commands that proxy execution into docker containers.

HOW AND WHY?!

1. Commit dev-tools to your project by a config file
2. Bootstrap all command executables from that config file
3. Reduce custom commands to docker into one configuration file per project.

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

### create symlink commands
    docker-proxy symlinks

symlinks are created and point to the docker-proxy exe, which will emulate the configured command
depending on which symlink was called.

In case of the configuration sample it would be a symlink called **php**.

If the folder containing the symlinks is added to the $PATH (for example by direnv)
you can execute exactly the configured php -version by just calling **php**.


### bootstrap containers (?!)
    docker-proxy prepare

all required images are pulled