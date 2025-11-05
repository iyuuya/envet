# envet

## Installation
```sh
go install github.com/ysugimoto/envet/cmd/envet@latest
```

## Usage

```sh
$ envet
envet is a tool to manage environment variables stored in the system keychain.

Usage:
  envet [command]

Available Commands:
  add              add a new environment variable
  completion       Generate the autocompletion script for the specified shell
  export           print export environment variables
  help             Help about any command
  list-keys        list keys in a namespace
  list-namespaces  list all namespaces
  remove           remove an environment variable
  remove-namespace remove a namespace
  run              run command with environment variables
  unset            print unset environment variables

Flags:
  -h, --help   help for envet

Use "envet [command] --help" for more information about a command.
```

### Completion

```sh
eval "$(envet completion zsh)"  # for zsh
```

see also: `envet completion`

supports: bash, zsh, fish, powershell (cobra generated)

## Integration Example

### fzf
```sh
envet export $(envet lsn | fzf --preview 'envet lsk {}')
```

### mise
```toml
# mise.toml
[hooks.enter]
shell = "zsh"
script = """
eval "$(envet export hello-ns)"
"""

[hooks.leave]
shell = "zsh"
script = """
eval "$(envet unset hello-ns)"
"""
```

### direnv
```
# .envrc
direnv_load envet run aws-retrip direnv dump
```
