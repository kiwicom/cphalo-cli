# CPHalo CLI

A command line tool for accessing CloudPassage Halo API

[![coverage report](https://gitlab.com/kiwicom/cphalo-cli/badges/master/pipeline.svg)](https://gitlab.com/kiwicom/cphalo-cli/pipelines)
[![pipeline status](https://gitlab.com/kiwicom/cphalo-cli/badges/master/coverage.svg)](https://gitlab.com/kiwicom/cphalo-cli/commits/master)
[![mit license](https://img.shields.io/badge/license-MIT-green.svg)](https://gitlab.com/kiwicom/cphalo-cli/blob/master/LICENSE)
[![go report](https://goreportcard.com/badge/gitlab.com/kiwicom/cphalo-cli)](https://goreportcard.com/report/gitlab.com/kiwicom/cphalo-cli)
[![go doc](https://godoc.org/gitlab.com/kiwicom/cphalo-cli?status.svg)](https://godoc.org/gitlab.com/kiwicom/cphalo-cli)
[![contribute](https://img.shields.io/badge/contributions-welcome-brightgreen.svg)](https://gitlab.com/kiwicom/cphalo-cli/forks/new)


## Usage

### Install

**Download code**

```bash
go get gitlab.com/kiwicom/cphalo-cli
```

**Install tool**

```bash
go install gitlab.com/kiwicom/cphalo-cli
```

### Configure credentials

Following command will ask you about credentials and save them to `~.cphalo.yaml` file.

```bash
cphalo config
```

Headless usage

```bash
cphalo config --key=my-key --secret=my-secret
```

You can also change the path of config:

```bash
cphalo config --path=/some/other/config.yaml
```

## Supported command

To check which commands are support just use `help`:

```bash
cphalo -h
``` 

## Docker usage

```bash
docker run --rm -it \
    --env CPHALO_APPLICATION_KEY=key \
    --env CPHALO_APPLICATION_SECRET=secret \
    registry.gitlab.com/kiwicom/cphalo-cli
```
