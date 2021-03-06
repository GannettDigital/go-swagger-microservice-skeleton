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
<dd>The name of the repository, i.e. github.com/GannettDigital/{{ RepoName }}.</dd>

<dt>SnakeName</dt>
<dd>A <a href="https://en.wikipedia.org/wiki/Snake_case">snake case</a> version of the RepoName</dd>

<dt>CamelName</dt>
<dd>A capitalized <a href="https://en.wikipedia.org/wiki/Camel_case">camel case</a> version of RepoName</dd>

<dt>Description
<dd>A description of the service</dd>

<dt>Version
<dd>The initial version of the service</dd>

</dl>

``` shell
    $ cd $GOPATH/src/github.com/GannettDigital
    $ boilr template use swagger-microservice my-echo-service
    [?] Please choose a value for "RepoName" [default: "echo-service"]: my-echo-service
    [?] Please choose a value for "Version" [default: "1.0.0"]:
    [✔] Created CHANGELOG.md
    [✔] Created Makefile
    [?] Please choose a value for "Description" [default: "A generic echo service"]:
    [✔] Created README.md
    [✔] Created config/config.go
    [✔] Created glide.yaml
    [✔] Created swag/Makefile
    [?] Please choose a value for "SnakeName" [default: "echo_service"]: my_echo_service
    [?] Please choose a value for "CamelName" [default: "EchoService"]: MyEchoService
    [✔] Created swag/restapi/configure_my_echo_service.go
    [✔] Created swagger.yaml
    [✔] Successfully executed the project template swagger-microservice in /Users/emoritz/work/src/github.com/GannettDigital/my-echo-service
```

# Running the echo service

To build and run the echo service do the following:

``` shell
    $ cd my-echo-service
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

You will need to delete the `swag/restapi/configure_*.go` file to
allow go-swagger to regenerate it based on your new swagger.yaml file.

``` shell
    $ rm swag/restapi/configure_*.go # delete the default echo server
    $ rm handlers/getoembed # removed the default handler
    $ make clean # removes any generated code
    $ make # build and tests the project
```

This project encourages you to move your handlers and middleware setup
out of the `configure_*.go` file and into the `handlers/` package.

If you significantly change your `swagger.yaml` file you may need to
delete the `configure_*.go` files and have go-swagger regenerate it.
In this event, you'll want the changes to `configure_*.go` to be
minimal.
