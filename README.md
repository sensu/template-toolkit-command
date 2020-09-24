[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/jspaleta/template-toolkit-command)
![goreleaser](https://github.com/jspaleta/template-toolkit-command/workflows/goreleaser/badge.svg)
[![Go Test](https://github.com/jspaleta/template-toolkit-command/workflows/Go%20Test/badge.svg)](https://github.com/jspaleta/template-toolkit-command/actions?query=workflow%3A%22Go+Test%22)
[![goreleaser](https://github.com/jspaleta/template-toolkit-command/workflows/goreleaser/badge.svg)](https://github.com/jspaleta/template-toolkit-command/actions?query=workflow%3Agoreleaser)

# Template Toolkit Command


# Template Toolkit

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Sensuctl command asset registration](#sensuctl-command-asset-registration)
  - [Sensuctl command usage](#sensuctl-command-usage)
- [Installation from source](#installation-from-source)
- [Contributing](#contributing)

## Overview

template-toolkit-command is a Sensuctl command plugin using the
[Sensu Plugin SDK][2].  This Sensuctl command plugin allows you to test and
validate Sensu plugin template strings, to ensure the templating logic works
correctly before you use them in other plugins.

## Usage examples
```
Sensuctl command plugin for validating and testing Sensu plugin templates

Usage:
  template-toolkit-command [flags]
  template-toolkit-command [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -t, --template string   A template string, in Golang text/template format
  -d, --dump-names        Dump the event object names that can be used in a Golang template
  -h, --help              help for template-toolkit-command

Use "template-toolkit-command [command] --help" for more information about a command.
```

## Configuration

### Sensuctl command asset registration

Sensuctl command assets are the best way to make use of this plugin. If
you're not using an asset, please consider doing so! You can use the
following command to add the asset:

```
sensuctl command add jspaleta/template-toolkit-command
```

### Sensuctl command usage
Requires event json representation stdin to operate correctly.

To test the output of a template:
```
cat event.json | sensuctl command execute jspaleta/template-toolkit-command -- --template "{{ .Check.Name}}"
Executing command with --template {{ .Check.Name }}
Template String Output: "keepalive"

sensuctl event info webserver01 check-http --format json | sensuctl command execute jspaleta/template-toolkit-command -- --template "Server: {{.Entity.Name}} Check: {{.Check.Name}} Status: {{.Check.State}}"
Executing command with --template Server: {{.Entity.Name}} Check: {{.Check.Name}} Status: {{.Check.State}}
Template String Output: Server: "webserver01 Check: check-http Status: passing"
```

To see the variable names available for use in a template:
```
sensuctl event info webserver01 check-http | sensuctl command execute jspaleta/template-toolkit-command -- --dump-names
.Entity{
    .EntityClass: "agent",
    .System:      .System{
        .Hostname:        "webserver01",
        .OS:              "linux",
[...]
.Check{
    .Command:           "check-http.rb -u http://webserver01/",
    .Handlers:          {"slack", "email"},
    .HighFlapThreshold: 0x0,
[...]
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an
asset. If you would like to compile and install the plugin from source or
contribute to it, download the latest version or create an executable from
this source.

From the local path of the template-toolkit-command repository:

```
go build
```

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
