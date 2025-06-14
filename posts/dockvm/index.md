---
title: Docker VMs & Bind Mounting
date: 2025-5-14
description: Using Docker containers as VMs to access external files via bind mounting
categories: ['docker', 'containers', 'mounting', 'experiment']
draft: false # Change to true to not render the post in on the website
---

guides (note they dont have '.' after build which is needed?)
- https://danielmangum.com/posts/simple-linux-command-line-using-docker/ 
- https://flaviocopes.com/docker-access-files-outside-container/ 

example `Dockerfile` (no extension):

```dockerfile
FROM ubuntu:latest
```

cmds from within the folder:

```bash
docker build . -t dockvm
docker run -t -d --name my_dvm -v //c/Users/Roark/external_dir:/internal_dir dockvm
docker ps
docker exec -i -t my_dvm /bin/bash

docker kill my_dvm
docker rm my_dvm
docker ps -a

docker start my_dvm
```

so build sets up the `Dockerfile`, run creates a new container running in the bkg `-d` with a unique `--name` that has access `-v` to an external directory internally under an alias, `ps` lists the containers to show it running (or not), and `exec` connects you to the terminal of the container 'vm', `kill` stops it immediately and `rm` removes to so that the `--name` doesnt overlap next time, can verify its gone by checking `ps -a` to list all containers, can also just use `start` next time instead of `run` so you wont have to remove it but you won't be able to change the dir

side note docker build not really nec here, docker compose or pull should be fine for this