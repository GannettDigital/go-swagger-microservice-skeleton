# go-swagger-microservice-skeleton

This is the bare minimum in order to generate a swagger service using https://goswagger.io/

# Creating a new service

Let's say your creating a new service called `uw-dummy-service`.

``` shell
$ make
Enter Project Name: uw-dummy-service
```

This creates a project in `$GOPATH/github.com/GannettDigital/uw-dummy-service`.

Next, copy your service's swagger.yaml in
`$GOPATH/github.com/GannettDigital/uw-dummy-service` and run `make`

Finally, customize the project as described by
https://goswagger.io/generate/server.html
