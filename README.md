# go-swagger-microservice-skeleton

This is a [boilr](https://github.com/tmrts/boilr) template for
generating a [go-swagger](https://goswagger.io/) based microservice.

# Installation

Use the following commands to install boilr and the template

``` shell
    $ go get github.com/tmrts/boilr
    $ boilr init
    $ boilr template download GannettDigital/go-swagger-microservice-skeleton swagger-microservice
```

# Generation

## Variables

<dl>

<dt>RepoName</dt>
<dd>The name of the repository, i.e. github.com/GannettDigital/{{ RepoName }}.<dd>

<dt>ShortName</dt>
<dd>A short, one word name for the service.  It must be a valid go variable name</dd>

<dt>CamelName</dt>
<dd>A capitalized camel case version of ShortName</dd>

<dt>Description
<dd>A description of the service</dd>

<dt>Version
<dd>The initial version of the service</dd>

</dl>

``` shell
    $ cd $GOPATH/src/github.com/GannettDigital
    $ boilr template use swagger-microservice my-echo-service
```

# Running the echo service

To build and run the echo service do the following:

``` shell
    $ make # builds and runs test suite
    $ make run # starts the web service on port :8080
```

# Customization

go-swagger reads the `swagger.yaml` file and generates the models and
API handlers for your service.

This skeleton describes a simple OEmbed service that you can modify or
replace.

Start by altering the `swagger.yaml` to fit the resources required by
your service.

You will need to delete the `swag/restapi/configure_echo.go` to allow
go-swagger to regenerate it based on your new swagger.yaml file.

``` shell
    $ rm swag/restapi/configure_*.go
    $ make clean # removes any generated code
    $ make # build and tests the project
```

go-swagger created a brand new dummy service for your `swagger.yaml`.
