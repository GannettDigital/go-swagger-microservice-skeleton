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
<dt>Name</dt>
<dd>A short name for the service.  It must be a valid go variable name</dd>
<dt>CamelName</dt>
  <dd>A capitalized camel case version of Name</dd>
<dt>Description
  <dd>A description of the service</dd>
<dt>Version
  <dd>The initial version of the service</dd>
</dl>

``` shell
    $ cd $GOPATH/src/github.com/GannettDigital
    $ boilr template use swagger-microservice my-echo-service
    [?] Please choose a value for "Name" [default: "echo"]: echo
    [?] Please choose a value for "Version" [default: "1.0.0"]:
    [✔] Created CHANGELOG.md
    [✔] Created Makefile
    [?] Please choose a value for "Description" [default: "A generic echo service"]:
    [✔] Created README.md
    [✔] Created config/config.go
    [✔] Created glide.yaml
    [✔] Created swag/Makefile
    [?] Please choose a value for "CamelName" [default: "Echo"]: MyEchoService
    [✔] Created swag/restapi/configure_echo.go
    [✔] Created swagger.yaml
    [✔] Successfully executed the project template swagger-microservice in /Users/emoritz/work/src/github.com/GannettDigital/my-echo-service
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
