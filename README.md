Torino
===

# Installation

```bash

# Install Docker
sudo pacman -S docker

# https://docs.docker.com/engine/installation/linux/archlinux/
sudo gpasswd -a USER docker

# logout, login
newgrp docker

# Start the docker service
sudo systemctl start docker

# Start docker on system boot
sudo systemctl enable docker

# to close

sudo systemctl stop docker

# You should be able to run docker as your user, try
docker ps

# Now, run
docker version

# And do
export DOCKER_API_VERSION=1.24 # Whatever version matches the server and the API.


```
