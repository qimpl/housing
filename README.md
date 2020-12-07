# Housing API

Manage housings through HTTP.

Part of the main qimpl application.

## API Documentation

Swagger API documentation can be found at `/v1/swagger`

### Generate Documentation

The Swagger API needs to be generated when modifications are done.

Run the next command to update this documentation:

```sh
# swaggo/swag cli is required
# go get -u github.com/swaggo/swag/cmd/swag
$ swag init
```

## Getting started

Run `make init` to copy all configuration files

Then you can run:

```sh
$ docker-compose up
```

## Code Quality

Your commits names should follow the [@commitlint/config-conventional](https://github.com/conventional-changelog/commitlint/tree/master/@commitlint/config-conventional) rules.

Check and fix your code using:

```sh
$ make coding-style
```

## Pre-commit

Install pre-commit following the [official documentation](https://pre-commit.com/#installation)

Setup your pre-commit hooks using:

```sh
# pre-commit hooks
$ pre-commit install

# message commit hooks
$ pre-commit install --hook-type commit-msg
```

For additional information check [pre-commit docs](https://pre-commit.com)
