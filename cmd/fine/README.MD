# fine

`Fine` is a powerful CLI tool for building [Fine](https://gitlab.sys.hxsapp.net/hxs/fine) application with convenience.

## 1. Install

## 1) PreCompiled Binary

You can also install `gf` tool using pre-built binaries: <https://github.com/gogf/gf/releases>

1. Database support

   |     DB     | builtin support |                                                                             remarks                                                                              |
   |:----------:|:---------------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------------:|
   |   mysql    |       yes       |                                                                                -                                                                                 |

## 2) Manually Install

```shell
go install gitlab.sys.hxsapp.net/hxs/fine/cmd/fine@latest # latest version
```

## 2. Commands

```html
$ fine
Usage:
[command]

Available Commands:
completion  Generate the autocompletion script for the specified shell
genctrl     Generate controller and method by defining api structs
gendao      Generate relevant data models based on table definitions
genservice  Generate service file by parse struct and associated functions from packages
help        Help about any command

Flags:
-h, --help   help for this command

Use " [command] --help" for more information about a command.
```