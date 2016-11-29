Torino
===

Start the docker daemon

```bash

# <BS>https://wiki.archlinux.org/index.php/Docker#Installation
sudo gpasswd -a USER docker

# logout, login
newgrp docker

systemctl start docker

# to close

systemctl stop docker

sudo dockerd

```
