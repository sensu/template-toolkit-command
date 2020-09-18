[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/jspaleta/template-toolkit-command)
![goreleaser](https://github.com/jspaleta/template-toolkit-command/workflows/goreleaser/badge.svg)
[![Go Test](https://github.com/jspaleta/template-toolkit-command/workflows/Go%20Test/badge.svg)](https://github.com/jspaleta/template-toolkit-command/actions?query=workflow%3A%22Go+Test%22)
[![goreleaser](https://github.com/jspaleta/template-toolkit-command/workflows/goreleaser/badge.svg)](https://github.com/jspaleta/template-toolkit-command/actions?query=workflow%3Agoreleaser)

# Template Toolkit Command


# Template Toolkit

## Table of Contents
- [Overview](#overview)
- [Files](#files)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Handler definition](#handler-definition)
  - [Annotations](#annotations)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

template-toolkit-command is a Sensuctl command plugin using the [Sensu Plugin SDK][2].
This Sensuctl command plugin allows you to test and validate Sensu plugin template strings, to ensure the templating logic works correctly before you use them in other plugins.


The Template Toolkit is a [Sensu Handler][6] that ...

## Files

## Usage examples

## Configuration

### Sensuctl command asset registration

Sensuctl command Assets are the best way to make use of this plugin. If you're not using an asset, please consider doing so! You can use the following command to add the asset:

```
sensuctl command add jspaleta/template-toolkit-command
```

#### Examples


## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the template-toolkit-command repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/handler-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/handler-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/handlers/
[7]: https://github.com/sensu-community/handler-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
