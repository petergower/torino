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

Problem of accessing STDOUT:
http://stackoverflow.com/questions/24621067/how-to-read-files-and-stdout-from-a-running-docker-container#24629857

Basically: create a file on a shared volume.

```bash

docker run -d -P --name torino -v ./_test:/app alpine ./app/main
```

Ok, so method that worked:

1) Build Go program
2) move it to images/torino/build/bin
3) build the torino Dockerfile image -- TODO: assign a version for each new binary
4) run the image like this: docker run -it torino /bin/torino
... and it works.

```bash

➜  ~ cd Development/go-petergower/torino
➜  torino git:(MVP-1) ✗ ls
images  LICENSE  main.go  README.md  test
➜  torino git:(MVP-1) ✗ cd test
➜  test git:(MVP-1) ✗ go build main.go ../images/torino/build/torino
named files must be .go files
➜  test git:(MVP-1) ✗ go build main.go
➜  test git:(MVP-1) ✗ ls
main  main.go
➜  test git:(MVP-1) ✗ cp main ../images/torino/build/bin
➜  test git:(MVP-1) ✗ cd ../images/torino/build/bin
➜  bin git:(MVP-1) ✗ ls
main  text.txt
➜  bin git:(MVP-1) ✗ mv main torino
➜  bin git:(MVP-1) ✗ cd ..
➜  build git:(MVP-1) ✗ cd ..
➜  torino git:(MVP-1) ✗ ls
build  Dockerfile
➜  torino git:(MVP-1) ✗ cd build
➜  build git:(MVP-1) ✗ ls
bin
➜  build git:(MVP-1) ✗ cd bin
➜  bin git:(MVP-1) ✗ ./torino
Hello!
➜  bin git:(MVP-1) ✗ cd ..
➜  build git:(MVP-1) ✗ cd ..
➜  torino git:(MVP-1) ✗ ls
build  Dockerfile
➜  torino git:(MVP-1) ✗ docker build -t torino .
Cannot connect to the Docker daemon. Is the docker daemon running on this host?
➜  torino git:(MVP-1) ✗ docker build -t torino .
Sending build context to Docker daemon 1.638 MB
Step 1 : FROM alpine
---> baa5d63471ea
Step 2 : COPY build/ /
---> 4299157d89ac
Removing intermediate container b9453e59db2c
Successfully built 4299157d89ac
➜  torino git:(MVP-1) ✗ docker run -it torino sh
/ # ls
bin      dev      etc      home     lib      linuxrc  media    mnt      proc     root     run      sbin     srv      sys      tmp      usr      var
/ # cd bin
/bin # ls
ash            conspy         echo           grep           ln             mount          pipe_progress  sed            text.txt
base64         cp             ed             gunzip         login          mountpoint     printenv       setserial      torino
bbconfig       cpio           egrep          gzip           ls             mpstat         ps             sh             touch
busybox        date           false          hostname       lzop           mv             pwd            sleep          true
cat            dd             fatattr        ionice         makemime       netstat        reformime      stat           umount
catv           df             fdflush        iostat         mkdir          nice           rev            stty           uname
chgrp          dmesg          fgrep          ipcalc         mknod          pidof          rm             su             usleep
chmod          dnsdomainname  fsync          kbd_mode       mktemp         ping           rmdir          sync           watch
chown          dumpkmap       getopt         kill           more           ping6          run-parts      tar            zcat
/bin # ./torino
Hello!
/bin # exit
➜  torino git:(MVP-1) ✗ docker run -it torino /bin/torino
Hello!
➜  torino git:(MVP-1) ✗ docker run -it torino /bin/torino
Hello!
➜  torino git:(MVP-1) ✗


```
