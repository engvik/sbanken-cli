# sbanken-cli

![tests](https://github.com/engvik/sbanken-cli/workflows/main/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/engvik/sbanken-cli)](https://goreportcard.com/report/github.com/engvik/sbanken-cli)

`sbanken` provides an easy way to interact with your bank from the terminal

```
$ sbanken accounts list
+----------------------------------+--------------------+-----------------------------+-------------+----------+-----------+--------------+
| ID                               | TYPE               | NAME                        | NUMBER      |  BALANCE | AVAILABLE | CREDIT LIMIT |
+----------------------------------+--------------------+-----------------------------+-------------+----------+-----------+--------------+
| D9D073EBC72557BA293288BA311BFB14 | Standard account   | Brukskonto                  | 00000000000 | 13371.33 |  13371.33 |            0 |
+----------------------------------+--------------------+-----------------------------+-------------+----------+-----------+--------------+
|                                  |                    |                             |             | 13371.33 |  13371.33 |            0 |
+----------------------------------+--------------------+-----------------------------+-------------+----------+-----------+--------------+
```


## Prerequisites

* Access to [Sbankens API](https://sbanken.no/bruke/utviklerportalen/). 

## Installation

1. Download the [latest release](https://github.com/engvik/sbanken-cli/releases)
2. Extract the executable binary.

## Required configuration

The interact with Sbanken you must provide a client ID, client secret and customer ID. See link under [Prerequisites](https://sbanken.no/bruke/utviklerportalen/) on how to obtain the client ID and client secret. The customer ID is your social security number.

There are three ways to pass these values to `sbanken`:

### Configuration file

```yaml
client-id: "my-client-id"
client-secret: "my-client-secret"
customer-id: "my-customer-id"
```

| **OS**  | **Default config locations**                            |
|---------|---------------------------------------------------------|
| linux   | `$HOME/.config/sbanken/config.yaml`                     |
| darwin  | `$HOME/Library/Application Support/sbanken/config.yaml` |
| windows | `%APPDATA%\sbanken\config.yaml`                         |

You can also specifig the config path:

`sbanken --config ~/.config/sbanken/config.yaml accounts list`

or 

`SBANKEN_CONFIG="~/.config/sbanken/config.yaml" sbanken accounts list`


### Environment variables

`SBANKEN_CLIENT_ID="my-client-id" SBANKEN_CLIENT_SECRET="my-client-secret" CUSTOMER_ID="my-customer-id" sbanken accounts list`

### Global application options

`sbanken --client-id="my-client-id" --client-secret="my-client-secret" --customer-id="my-customer-id" accounts list`

## Optional configuration

The following configuration can be set by passing a global option or in the configuration file.

### Output styles

Global option: `style`
Config field: `style`

Available styles:

```
"bold"
"colored-bright"
"colored-dark"
"colored-black-on-blue-white"
"colored-black-on-cyan-white"
"colored-black-on-green-white"
"colored-black-on-magenta-white"
"colored-black-on-yellow-white"
"colored-black-on-red-white"
"colored-blue-white-on-black"
"colored-cyan-white-on-black"
"colored-green-white-on-black"
"colored-magenta-white-on-black"
"colored-red-white-on-black"
"colored-Yellow-white-on-black"
"double"
"light"
"rounded"
```

## Usage

```
NAME:
   sbanken - provides an easy way to interact with your bank from the terminal

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.1.0

COMMANDS:
   accounts, a        list and read accounts
   cards, c           list cards
   efakturas, e       list, read and pay efakturas
   payments, p        list and read payments
   standingorders, s  list standing orders
   transactions, ta   list transactions
   transfers, tf      transfer money between accounts
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --client-id value, --clid value    the client id [$SBANKEN_CLIENT_ID]
   --client-secret value, -s value    the client secret [$SBANKEN_CLIENT_SECRET]
   --customer-id value, --cuid value  customer id [$SBANKEN_CUSTOMER_ID]
   --style value                      set output style
   --config value, -c value           path to YAML config [$SBANKEN_CONFIG]
   --help, -h                         show help (default: false)
   --version, -v                      print the version (default: false)
```
