# Street Fair

![Street Fair](docs/street-fair.jpeg)

The objective of this project is to centralize information about street fairs in the city of São Paulo.

# How to Run

* **Install dependencies**

``` bash
$ make clean install
```

* **To running local dependencies** if you want to run the application using local dependencies with docker, run the command below before running the API

``` bash
$ make test/deps/up
```

* **To run import CSV**

``` bash
$ make import file=path-to-file.csv
```

* **To stop**

``` bash
$ make test/deps/down
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
