# WIP - proof of concept

# The idea
Create a tool that helps creating variants of commands that proxy execution into docker containers.

HOW?!

1. Commit dev-tools to your project by a config file
2. Bootstrap all command executables from that config file


## the config (toml)
    Version="1"

    [[command]]
        name = "php"
        entryPoint = "php"
        image = "php:7.1.13"
        addGroups = true
        impersonate = true
        workDir = "/app"
        removeContainer=true
        volumes = [
            "${APP_DIR}:/app",
            "${SSH_AUTH_SOCK}:/run/ssh.sock",
            "/etc/passwd:/etc/passwd:ro",
            "/etc/group:/etc/group:ro",
            "/run/docker.sock:/run/docker.sock"
        ]
        envvars = [
            "SSH_AUTH_SOCK:/run/ssh.sock",
            "DOCKER_HOST=unix:///run/docker.sock"
        ]

### create symlink commands
    docker-proxy symlinks

symlinks are created and point to the docker-proxy exe, which will emulate the configured command
depending on which symlink was called.

In case of the configuration sample it would be a symlink called **php**.

If the folder containing the symlinks is added to the $PATH (for example by direnv)
you can execute exactly the configured php -version by just calling **php**.
