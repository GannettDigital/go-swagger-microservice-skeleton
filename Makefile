NAME=$$(basename $$(pwd))

all: deps

deps: swag
	glide up

swag:
	swagger generate server

clean:
	rm -f glide.lock
	rm -rf {models,cmd,vendor}
	find . -type file -path './restapi/*' -not -path './restapi/configure_*.go' -delete
	find . -type d -empty -delete


# This will remove all customization to the skeleton, you it at your own peril.
cleanall: clean
	rm -rf restapi

run:
	go install ./cmd/${NAME}-server/
	${NAME}-server
