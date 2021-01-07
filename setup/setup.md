# Setup

## Create 3 Cento8.2 VMs

## Setup K3s

We need to enable container features in the kernel, edit /boot/cmdline.txt and add the following to the end of the line:
```sh
 cgroup_enable=cpuset cgroup_memory=1 cgroup_enable=memory
```
and reboot...

Run this script on the Server node

```sh
curl -sfL https://get.k3s.io | sh -
```
A kubeconfig file is written to /etc/rancher/k3s/k3s.yaml and the service is automatically started or restarted. The install script will install K3s and additional utilities, such as kubectl, crictl, k3s-killall.sh, and k3s-uninstall.sh

K3S_TOKEN is created at /var/lib/rancher/k3s/server/node-token on your server. To install on worker nodes we should pass K3S_URL along with K3S_TOKEN or K3S_CLUSTER_SECRET environment variables

Grab the join key from this node with:
```sh
sudo cat /var/lib/rancher/k3s/server/node-token
```

Change the owner of the kubeconfig

```sh
sudo chown -R pi:pi /etc/rancher/k3s/
```

Connect to the agents and install 
```sh
export K3S_URL="https://ip:6443"
export K3S_TOKEN="XXX"
``` 


## Create NFS

### In the Server
```sh
sudo apt install nfs-kernel-server
```

Create the shared directory
```sh
sudo mkdir -p /home_nfs
```

Add the clients addresses
vi `/etc/exports`
```sh
/home_nfs	192.168.X.XXX(rw,sync,no_all_squash,root_squash)
/home_nfs   192.168.X.XXX(rw,sync,no_all_squash,root_squash)
```

Add the shared folder in an "open" user:group 
```sh
sudo chown -R nobody: /home_nfs
sudo chmod -R 777 /home_nfs
```

Restart the service
```sh
sudo systemctl restart nfs-kernel-server
```

### In the Clients 
```sh
sudo apt install nfs-common
```

Mount the filesystem

```sh
mount -t nfs 192.168.1.233:/home_nfs /home_nfs
```

## Setup hyperledger fabric

Create the namespace

```sh
kubectl create ns hyperledger
```

## Bonus: Compile the hyperledger images in ARM

1. Enable docker experimental
```sh
vim ~/.docker/config.json
add "experimental": "enabled"
```

2. Create new builder for docker
```sh
sudo apt-get install qemu-user-static -y
docker run --rm --privileged multiarch/qemu-user-static --reset -p yes i
docker buildx rm multibuilder
docker buildx create --name multibuilder
docker buildx ls
docker buildx inspect multibuilder
docker buildx inspect multibuilder --bootstrap
docker buildx use multibuilder
docker ps -a
```