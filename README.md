# Street Fair

![Street Fair](docs/street-fair.jpeg)

The objective of this project is to centralize information about street fairs in the city of São Paulo.



## Running

This app is written in Golang. If you don't have experience with it: https://golang.org/doc/install.
It depends on a Golang version with module support (at least 1.11, latest version is recommended).
Should you interoperate with Golang versions >= 1.11, don't forget to proper set your _GO111MODULE_ environment variable (https://dev.to/maelvls/why-is-go111module-everywhere-and-everything-about-go-modules-24k)

### Install deps

- Install `docker` and `docker-compose`

*Attention*: you should first set `$GOBIN` environment variable, add it to the `$PATH`, and then run the commands bellow:


# How to Run

* **Install dependencies**

``` bash
$ make clean install
```

* **To running local dependencies** if you want to run the application using local dependencies with docker, run the command below before running the API

``` bash
$ make deps/up
```

* **To run import CSV**

``` bash
$ make run/import file=docs/example.csv
```

* **To stop**

``` bash
$ make deps/down
```

* **To run API**

``` bash
$ make run/api
```
* **To help**

``` bash
$ make help
```

## Test coverage

* **To generate**
``` bash
$ make coverage
```

* **To show in browser**
``` bash
$ make coverage-html
```

## Runing tests
Just run:
``` bash
$ make test
```

# Log
All logs are stored in the fair.log file.
To show logs:

``` bash
$ make show/logs
```

Then commit the contents of `docs/swagger` dir.

# Swagger

When change/add any HTTP API stuff, just run:

``` bash
$ make swagger/api 
```

Then commit the contents of `docs/swagger` dir.


# Docker

This tasks below are used from CircleCI flow but FTW:

Generate `test` image to run tests and coverage:

```sh
make docker/build
```

With the `test` image we can run test and coverage inside CircleCI:

```sh
make docker/test
make docker/coverage
```

Generate final image getting binary from `test` image:

```sh
make docker/image
```
