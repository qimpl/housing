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

Check and fix your code using:

```sh
$ make coding-style
```
