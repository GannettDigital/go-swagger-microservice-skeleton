REPO=$$(basename $$(pwd))

BUILD_DEPS=golang.org/x/tools/cmd/cover \
	github.com/Masterminds/glide \
	github.com/golang/lint/golint \
	github.com/go-swagger/go-swagger/cmd/swagger \
	github.com/go-openapi/runtime


all: fmt bootstrap swag deps test

fmt:
	go fmt $$(glide nv)

vet:
	go vet $$(glide nv)

lint:
	glide nv | xargs -n1 golint -set_exit_status

bootstrap:
	go get -v $(BUILD_DEPS)

deps:
	glide install

swag:
	swagger generate server -t swag
	# If you need to build a swagger client, do the following:
	# swagger generate client -t swag/backend -f taxo.swagger.yaml -A taxo -c taxo


clean:
	@make -C swag clean

run:
	go install ./swag/cmd/${REPO}-server/
	${REPO}-server --port 8000

test: vet lint
	go test $$(glide novendor)

.PHONY: all fmt vet lint bootstrap deps swag clean run test
