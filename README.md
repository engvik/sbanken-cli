# sbanken-cli

`sbanken` provides an easy way to interact with your bank from the terminal

## Prerequisites

* Access to [Sbankens API](https://sbanken.no/bruke/utviklerportalen/). 

## Installation

TODO

## Required configuration

The interact with Sbanken you must provide a client ID, client secret and customer ID. See link under [Prerequisites](https://sbanken.no/bruke/utviklerportalen/) on how to obtain the client ID and client secret. The customer ID is you social security number.

There are three ways to pass these values to `sbanken`:

### Configuration file

```yaml
client-id: "my-client-id""
client-secret: "my-client-secret"
customer-id: "my-customer-id"
```

`sbanken --config ~/.config/sbanken/config.yaml accounts list`

### Environment variables

`SBANKEN_CLIENT_ID="my-client-id" SBANKEN_CLIENT_SECRET="my-client-secret" CUSTOMER_ID="my-customer-id" sbanken accounts list`

### Global application options

`sbanken --client-id="my-client-id" --client-secret="my-client-secret" --customer-id="my-customer-id" accounts list`

## Usage

```
$ sbanken --help
NAME:
   sbanken - interact with sbanken through the command line

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   accounts        interact with accounts
   cards           interact with cards
   efakturas       interact with efakturas
   payments        interact with payments
   standingorders  interact with standing orders
   transactions    interact with transactions
   transfers       interact with transfers
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --client-id value      the client id [$SBANKEN_CLIENT_ID]
   --client-secret value  the client secret [$SBANKEN_CLIENT_SECRET]
   --customer-id value    customer id [$SBANKEN_CUSTOMER_ID]
   --config value         path to YAML config
   --help, -h             show help (default: false)
   --version, -v          print the version (default: false)
```
