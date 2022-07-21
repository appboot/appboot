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

Appboot 取 application BootLoader 之意, 是一个通用的应用创建平台。

Appboot 基于[模板](https://github.com/appboot/templates)创建应用。你可以自定义模板来满足个性化的需求。

模板里包含**代码**、**配置**(模板描述、参数列表、自定义脚本)等。

在自定义的脚本里，你可以做**提交代码、配置 CI&CD、DevOps** 等任何事情。

---

## 安装

```sh
# Go 1.16+
go install github.com/appboot/appboot@v0.8.0

# Go version < 1.16
go get -u github.com/appboot/appboot@v0.8.0
```

## 快速开始

### 命令行工具

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

访问 <http://localhost:3000/>，即可**更新模板**并**创建项目**。

![](https://cdn.jsdelivr.net/gh/appboot/resources@master/appboot-web.gif)

> 注：
>
> - `-e API_URL`：API 的地址。appboot docker 包含了前后端，所以 `API_URL` 就是部署 appboot docker 的地址，其中 API 后端服务端口为 `8000`，前端端口为 `80`。
> - `-v $HOME/appboot:/root/.appboot`：将 appboot 的**工作目录**映射到本机的 `$HOME/appboot`，这样就可以将[配置](#配置)和**数据**持久化到本地。
>   当创建完 `test` 项目后，即可在工作目录中找到
>
>   ```sh
>   ❯ tree -a -L 2 $HOME/appboot
>   /Users/catchzeng/appboot
>   ├── workspace
>   │   └── test
>   └── templates
>       ├── GO-CMD
>       ├── README-CN.md
>       ├── README.md
>       ├── SwiftUI
>       └── VUE
>
>   6 directories, 2 files
>   ```

> - `-p 8000:8000`：将 API 后端服务的 `8000` 端口映射为本机的 `8000`
> - `-p 3000:80`：将前端的 `80` 端口映射为本机的 `3000`

## 配置

appboot 配置文件 `config.yaml` 位于 `$HOME/.appboot/` 目录下。如果没有该文件，你可以自行创建它。

当前配置文件支持 **templateSource** 和 **templateRoot**。

- templateSource: 指定获取模板的代码仓库，默认值是 <https://github.com/appboot/templates.git> 。如果默认的仓库没有你想要的模板，可以指定为自己的代码仓库。

- templateRoot: 指定了模板的存储和获取路径，默认值是 `$HOME/.appboot/templates/`。比如设置 `templateRoot: /Users/catchzeng/Desktop/templates` 表示指定模板的路径为我的桌面。

```shell
$ cat $HOME/.appboot/config.yaml
templateSource: https://github.com/CatchZeng/templates.git
templateRoot: /Users/catchzeng/Desktop/templates
```
