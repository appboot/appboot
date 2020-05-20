# backend

> appboot backend

## Usage

### Docker

#### Run

```shell
// replace 192.168.0.1 with your host IP or domain
docker run -d --name backend \
 -e HOST_IP=192.168.0.1 \
 -v $HOME/.ssh:/root/.ssh \
 -v $HOME/.gitconfig:/root/.gitconfig \
 -v $HOME/.appboot:/root/.appboot \
 -p 8888:8888 \
 appboot/backend
```

with [jenkinsapi](https://github.com/CatchZeng/jenkinsapi)

```shell
// replace 192.168.0.1 with your host IP or domain
docker run -d --name backend \
 -e HOST_IP=192.168.0.1 \
 -v $HOME/.ssh:/root/.ssh \
 -v $HOME/.gitconfig:/root/.gitconfig \
 -v $HOME/.appboot:/root/.appboot \
 -v $HOME/.jenkinsapi:/root/.jenkinsapi \
 -p 8888:8888 \
 appboot/backend:jenkinsapi-latest
```

#### Get WS_URL

```sh
$ docker logs backend
WS_URL: ws://192.168.0.1:8888/appboot
```