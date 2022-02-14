# appboot

![Go](https://github.com/appboot/appboot/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/appboot/appboot/branch/master/graph/badge.svg)](https://codecov.io/gh/appboot/appboot)
[![Go Report Card](https://goreportcard.com/badge/github.com/appboot/appboot)](https://goreportcard.com/report/github.com/appboot/appboot)
[![Release](https://img.shields.io/github/release/appboot/appboot.svg)](https://github.com/appboot/appboot/releases)
[![GoDoc](https://godoc.org/github.com/appboot/appboot?status.svg)](https://pkg.go.dev/github.com/appboot/appboot?tab=doc)

<p align="center">
  <br>
  <img width="20%" src="./assets/logo.svg" alt="logo">
  <br>
  <br>
</p>

[中文请戳](./README-zh.md)

Appboot, which means application BootLoader, is a universal platform for creating applications.

Appboot creates applications based on [templates](https://github.com/appboot/templates). You can create custom templates to meet different requirements.

A template contains **code, configuration(description, parameters, custom scripts)**, and more.

In custom scripts, you can **do anything** like **commit code, configure CI&CD, DevOps**, etc.

## Install

```sh
# Go 1.16+
go install github.com/appboot/appboot@v0.7.0

# Go version < 1.16
go get -u github.com/appboot/appboot@v0.7.0
```

## Quick Start

### Command Line Tool

```shell
appboot create
```

![](https://cdn.jsdelivr.net/gh/appboot/resources@master/appboot.gif)

### appboot web

```sh
$ docker run -it \
-e API_URL=http://127.0.0.1:8000 \
-v $HOME/appboot:/root/.appboot \
-p 8000:8000 \
-p 3000:80 \
appboot/appboot
```

Visit <http://localhost:3000/> to **update templates** and **create projects**.

![](https://cdn.jsdelivr.net/gh/appboot/resources@master/appboot-web.gif)

> Note:
>
> - `-e API_URL`: API URL. appboot docker includes frontend and backend, so `API_URL` is the address where appboot docker is deployed, where the API backend service port is `8000`, and the frontend port is `80`.
> - `-v $HOME/appboot:/root/.appboot`: Map the appboot's **working directory** to `$HOME/appboot`, so that [configuration](#Configuration) and **Data** can be persisted locally.
>   Once the `test` project is created, it can be found in the working directory
>
> ```sh
> ❯ tree -a -L 2 $HOME/appboot
> /Users/catchzeng/appboot
> ├── workspace
> │ └── test
> └── templates
> ├── GO-CMD
> ├── README-CN.md
> ├── README.md
> ├── SwiftUI
> └── VUE
>
> 6 directories, 2 files
> ```

> - `-p 8000:8000`: map the `8000` port of the API backend service to the local `8000`
> - `-p 3000:80`: map the frontend `80` port to the local `3000`

## Configuration

Appboot configuration file(`config.yaml`) is placed in the `$HOME/.appboot/` directory. If you do not have this file, you can create the file yourself.

The current configuration file supports **templateSource** and **templateRoot**.

- templateSource: specifies the source repository for obtaining templates. The default value is <https://github.com/appboot/templates.git>. If the default repository cannot meet the requirements, you can specify your own repository. For example, `templateSource: https://github.com/CatchZeng/templates.git`.

- templateRoot: specifies the storage and retrieval path of the template. The default value is `$HOME/.appboot/templates/`. For example, `templateRoot: /Users/catchzeng/Desktop/templates` means set the template path to my desktop templates directory.

```shell
$ cat $HOME/.appboot/config.yaml
templateSource: https://github.com/CatchZeng/templates.git
templateRoot: /Users/catchzeng/Desktop/templates
```
