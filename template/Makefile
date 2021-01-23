GO=CGO_ENABLED=0 GO111MODULE=off go

gProject:
	$(GO) build -o /go/src/Gproject/demo $(target)

demo:
	docker build \
		-f ./Dockerfile \
		-t demo \
		.