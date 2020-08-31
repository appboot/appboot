# appboot

<p align="center">
  <br>
  <img width="20%" src="./assets/logo.svg" alt="logo">
  <br>
  <br>
</p>

Appboot 取 application BootLoader 之意, 是一个通用的应用创建平台。

Appboot 基于[模板](https://github.com/appboot/templates)创建应用。你可以自定义模板来满足个性化的需求。

在模板里，你可以包含代码、lint 配置、CI&CD、自定义脚本等。

---

## 安装

```shell
go get -u github.com/appboot/appboot
```

## 快速开始

### 命令行工具

```shell
appboot create
```

![](https://github.com/appboot/resources/blob/master/appboot.gif?raw=true)

## 配置项

appbctl 配置文件 `config.yaml` 位于 `$HOME/.appboot/` 目录下。如果没有该文件，你可以自行创建它。

当前，在配置文件中支持 **templateRoot** 和 **templateSource** 的配置。

- templateRoot: 指定了模板的存储和获取路径，默认值是 `$HOME/.appboot/templates/`。比如设置 `templateRoot: /Users/catchzeng/Desktop/templates` 表示指定模板的路径为我的桌面。

- templateSource: 指定获取模板的代码仓库，默认值是 <https://github.com/appboot/templates.git>. 如果默认的仓库没有你想要的模板，你可以指定为自己的代码仓库。比如设置 `templateSource: https://github.com/CatchZeng/templates.git`。

```shell
$ cat $HOME/.appboot/config.yaml
templateSource: https://github.com/CatchZeng/templates.git
templateRoot: /Users/catchzeng/Desktop/templates
```
