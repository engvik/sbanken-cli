# sbanken-cli

`sbanken` provides an easy way to interact with your bank from the terminal

```
$ sbanken -c ~/.config/sbanken/config.yaml accounts list
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

1. Download the [latest release](https://github.com/engvik/sbanken-cli)releases)
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

`sbanken --config ~/.config/sbanken/config.yaml accounts list`

or 

`SBANKEN_CONFIG="~/.config/sbanken/config.yaml" sbanken accounts list`


### Environment variables

`SBANKEN_CLIENT_ID="my-client-id" SBANKEN_CLIENT_SECRET="my-client-secret" CUSTOMER_ID="my-customer-id" sbanken accounts list`

### Global application options

`sbanken --client-id="my-client-id" --client-secret="my-client-secret" --customer-id="my-customer-id" accounts list`

## Usage

```
NAME:
   sbanken - interact with sbanken through the command line

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

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
   --config value, -c value           path to YAML config
   --help, -h                         show help (default: false)
   --version, -v                      print the version (default: false)
```
