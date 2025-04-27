---
title: Inject 1
date: 2025-4-26
description: Podman public modded minecraft server
categories: ['inject', 'minecraft', 'podman']
draft: false # Change to true to not render the post in on the website
---

# Inject 1
This week Jeffrey has requested that I setup a modded Minecraft server accessible to the public for team building! He has added a twist, however, the server must be setup with Podman (an alt to Docker) and rootless at that!

I'm not sure how doable this is, but we'll give it a shot

- Requirement Steps
    - Setup Podman
    - Make Podman rootless
    - Setup modded Minecraft container
    - Test public connection

I chose Ubuntu 24 Noble because that's what I use managing my own MC server so it's what I'm more used to. I could have chosen something new but I wanted this to be a learning experience about Podman more so than an OS learning experience.

Ubuntu typically comes in 3 variants I noticed: `default`, `cloud`, `desktop`. Desktop should be obvious but I was wondering what the Cloud variant was about; apparently its the same as Default except that it's installed and compatible with `cloud-init` which you can read more about here https://cloud-init.io. 

After creation, in the configuration tab I attached the network to `public0` and changed the resource limits to allow for 8gb of ram and 4 cpu cores. This should be good for a solid Minecraft server especially if multithreading is used effectively. It should be ok to half these values though (or double!) **Note**: only works when VM is off, also another interface were it will never load when saved, just have to save and reload page

`wip`