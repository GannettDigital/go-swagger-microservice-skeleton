NAME=go-microservice-skeleton

default: generate

generate:
	@read -p "Enter Project Name: " project; \
	export ALLIGATOR=$$project && go generate
