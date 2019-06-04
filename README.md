# Simple NEO4J client in GO

Based on the snippet of code shown [here](https://github.com/neo4j/neo4j-go-driver).

## Prerequisites

### Installing openssl

Install openssl (including develop and static):

```
sudo yum install openssl
sudo yum install openssl-devel
sudo yum install openssl-static
```

### Installing seabolt
Download seabolt [here](https://github.com/neo4j-drivers/seabolt/releases)

Install seabolt:

`sudo rpm -i seabolt-1.7.4-Linux-centos-7.rpm` 

Add seabolt to pkg-config via:

`export PKG_CONFIG_PATH=/usr/local/share/pkgconfig`

Add seabolt library path via:

`export LD_LIBRARY_PATH=/usr/local/lib64`


### Run neo4j
The easiest way to run neo4j is by running it inside a docker container.

```
docker run \
    --publish=7474:7474 --publish=7687:7687 \
    --volume=/home/bi76au/data/tmp/neo4j-db:/data \
    --env=NEO4J_AUTH=none \
    neo4j
```

## Building the client

`go build neo4j-client.go`

or with static linking

`go build --tags seabolt_static neo4j-client.go`


### Running from GoLand
Add `--tags seabolt_static` to the Go tool arguments. This will compile with libraries statically linked.

Otherwise you have to add the  environment variable `LD_LIBRARY_PATH`.


