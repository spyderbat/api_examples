# Using the PicoAgent

## Arguments

`-c` / `--config`: The configuration file location to use

### Optional

`-u` / `--url`: A custom API url to use \
`-p` / `--prefix`: A custom prefix to use for generating the source ID \
`-m` / `--muid`: A custom machine ID to use for analytics \
`-i` / `--input`: Read bat data from standard input instead of sepcified bats \
`-v` / `--verbose`: Print logs of bat output and data sending

## Configuration Fields

`API Key`: The API key \
`Org ID`: The organization ID \
`Source`: The source ID \
`Source Desc`: The source description \
`Bats`: A list of bats to run, in the form of commands \
`URL`: The API url to use \
`Prefix`: The prefix to use for generating the source ID \
`MUID`: The machine ID to use if none is sourced from the bat \
`Timeout`: The delay between each send of collected data

## Configuration Handling and Startup Logic

```
read and validate config file
validate API key and organization ID
if config file does not contain a source ID:
    get list of sources
    generate a unique source ID
    create source with source ID
    save source ID to config
validate source ID (including that the prefix matches)
run and capture output from all bats specified
on bat output:
    validate minimum required output fields
    send output to the API with the source, key, and organization
```