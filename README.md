# appboot

<p align="center">
  <br>
  <img width="20%" src="./assets/logo.svg" alt="logo">
  <br>
  <br>
</p>

[中文请戳](./README-zh.md)

Appboot, which means application BootLoader, is a general application creation platform. 

Appboot creates applications based on [templates](https://github.com/appboot/templates). You can create custom templates to meet different requirements. 

In the template, you can include code, lint configuration, CI&CD, custom scripts, etc.

---

## Install

```shell
go get -u github.com/appboot/appboot
```

## Quick Start

### Command Line Tool

```shell
appboot create
```

![](https://github.com/appboot/resources/blob/master/appboot.gif?raw=true)

## Configuration

Appboot configuration file(`config.yaml`) is placed in the `$HOME/.appboot/` directory. If you do not have this file, you can create the file yourself.

Currently in `config.yaml` you can configure **templateRoot** and **templateSource**.

- templateRoot: specifies the storage and retrieval path of the template. The default is `$HOME/.appboot/templates/`. For example, `templateRoot: /Users/catchzeng/Desktop/templates` means set the template path to my desktop templates directory.

- templateSource: indicates the source repository for obtaining templates. The default is <https://github.com/appboot/templates.git>. If the default repository cannot meet the requirements, you can specify your own repository. For example, `templateSource: https://github.com/CatchZeng/templates.git`.

```shell
$ cat $HOME/.appboot/config.yaml 
templateSource: https://github.com/CatchZeng/templates.git
templateRoot: /Users/catchzeng/Desktop/templates
```
