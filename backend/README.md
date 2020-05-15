# backend

> appboot backend

## Usage

### Docker

#### Run

```shell
// replace 192.168.0.1 with your host IP
docker run -d --name backend -e HOST_IP=192.168.0.1 -p 8888:8888 appboot/backend
```

#### Get WS_URL

```sh
$ docker logs backend
WS_URL: ws://192.168.0.1:8888/appboot
```
