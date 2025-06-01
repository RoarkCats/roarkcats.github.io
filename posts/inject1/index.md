---
title: Inject 1 - Podman MC
date: 2025-4-26
description: Podman rootless public modded minecraft server
categories: ['inject', 'minecraft', 'podman']
draft: false # Change to true to not render the post in on the website
---

# Inject 1

### Intro
This week Jeffrey has requested that I setup a modded Minecraft server accessible to the public for team building! He has added a twist, however, the server must be setup with Podman (an alt to Docker) and rootless at that!

I'm not sure how doable this is, but we'll give it a shot

- Required Steps
    - Setup Podman
    - Make Podman rootless
    - Setup modded Minecraft container
    - Test public connection

### System Setup
I chose Ubuntu 24 Noble because that's what I use managing my own MC server so it's what I'm more used to. I could have chosen something new but I wanted this to be a learning experience about Podman more so than an OS learning experience.

Ubuntu typically comes in 3 variants I noticed: `default`, `cloud`, `desktop`. Desktop should be obvious but I was wondering what the Cloud variant was about; apparently its the same as Default except that it's installed and compatible with `cloud-init` which you can read more about here https://cloud-init.io. 

After creation, in the configuration tab I attached the network to `public0` and changed the resource limits to allow for 8gb of ram and 4 cpu cores. This should be good for a solid Minecraft server especially if multithreading is used effectively. It should be ok to half these values though (or double!)

**Note**: only works when VM is off, also interface never finishes loading when saved, just click save, wait a sec, and reload page

### Podman
By default Podman is built to be rootless, or rather built to run from the user that started its containers, so out of the box iut will be 'rootless' as long as you don't run it from root or [have multiple existing users which may conflict](https://opensource.com/article/19/2/how-does-rootless-podman-work)

One major downside to running rootless is the lack of access to privileged ports below 1024 (can't sue default 80 for a website)

Podman containers may not survive a reboot, so you must create a service

Guide I intended to follow (verbose): https://github.com/containers/podman/blob/main/docs/tutorials/rootless_tutorial.md
Guide I actually followed (well explained): https://youtu.be/69dADRzXpqk

Commands run: 
```bash
# Ubuntu 20.10 and newer
apt-get update
apt-get -y install podman
apt-get -y install podman-compose
apt-get -y install passt # was already installed
apt-get -y install nano # if not installed
# pasta should be selected in podman by default so no need to change settings
reboot now
# avoids potential errors
# Now to make a service that will restart when your system does
useradd -m -s /bin/bash podmanuser
#  home folder, shell (bash), (username)
passwd podmanuser # optional
su podmanuser # not optional!
cd # goto home dir
nano compose.yaml # your stuff
exit # return to root
nano /lib/systemd/system/podman-compose.service
systemctl --system daemon-reload
systemctl enable podman-compose.service
systemctl start podman-compose.service
# dont forget to address privileged port range if your user id 
```

`podman-compose.service`: 
```ini
[Unit]
Description=Podman-compose
Wants=network-online.target
After=network-online.target

[Service]
User=podmanuser
Group=podmanuser
Type=oneshot
RemainAfterExit=true
ExecStartPre=
ExecStart=/usr/bin/podman-compose -f /home/podmanuser/compose.yaml up -d
ExecStop=/usr/bin/podman-compose stop

[Install]
WantedBy=default.target
```

Podman side demo: 
```bash
podman run -d -p 8080:80/tcp --name webserver docker.io/library/httpd
#   in bkg, ports, (ext:int ports), name, (name), container image
# test cmd running site on 8080 public, just went to ip:8080 and "It works!"
podman ps -a
# list all containers
podman stop webserver
podman rm webserver
# but now lets remove it
# for compose same thing:
nano compose.yaml # your stuff
podman-compose up -d
#            run, in bkg
```

`compose.yaml`: 
```yaml
version: '3'

services:
  webserver:
    image: docker.io/library/httpd
    container_name: webserver
    ports:
      - '8080:80'
    restart: unless-stopped
```

### Minecraft

I found this compose setup online and modified it a bit to my needs, the [docker image](https://hub.docker.com/r/itzg/minecraft-server) does have its own [documentation](https://docker-minecraft-server.readthedocs.io/en/latest/) though on its use

`compose.yaml`:
```yaml
services:
    minecraft-server:
        image: docker.io/itzg/minecraft-server
        container_name: minecraft
        restart: unless-stopped
        userns_mode: keep-id
        stdin_open: true
        tty: true
        ports:
            - "25565:25565"
        environment:
            SERVER_NAME: "Podcraft"
            MOTD: "Testing..."
            EULA: "TRUE"
            TYPE: NEOFORGE
            VERSION: "1.20.1"
            NEOFORGE_VERSION: "latest"
            MAX_MEMORY: 8G
            VIEW_DISTANCE: 16
            MAX_PLAYERS: 16
            SERVER_PORT: 25565
        volumes:
            - /home/podmanuser/neoforge/data:/data
```

Commands run:
```bash
su podmanuser
cd
nano compose.yaml
podman-compose up -d
podman-compose ps # check status
```

While this all does work, it seems I'll need to spend a little more time on the service setup part as it seems not to work with the previous simple test setup