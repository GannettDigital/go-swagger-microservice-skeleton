# {{ RepoName }}
===================

{{ Description }}

# Running {{ RepoName }}

To build and run the echo service do the following:

``` shell
    $ make # builds and runs test suite
    $ make run # starts the web service on port :8080
```


# make commands

## make

Generates swagger code, installs dependencies, and run tests

## make clean

Removes any generated swagger code excluding the
`swag/restapi/configure_*.go` file.

## make run

Installs the server binary and runs it
