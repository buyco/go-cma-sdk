DOCKER_BUILD := docker run --rm -u `id -u` -v ${PWD}:/sdk openapitools/openapi-generator-cli:v6.0.0 generate -i sdk/api_files/dcsa/logistic.tracking.v1.json
GO_CLIENT := -g go -o /sdk/${API_TYPE}/cma \
			--git-repo-id=go-cma-sdk --git-user-id=buyco \
			--additional-properties=packageName=cma \
			--additional-properties=isGoSubmodule=true \
			--additional-properties=generateInterfaces=true

## generate: Clean and generate SDK from DCSA file.
generate:
	${MAKE} clean go-sdk

go-sdk:
	${DOCKER_BUILD} ${GO_CLIENT}

clean:
	rm -rf ./${API_TYPE}/cma

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo