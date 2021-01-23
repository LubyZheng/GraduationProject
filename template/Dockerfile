FROM golang

RUN mkdir /go/src/Gproject

COPY . /go/src/Gproject

WORKDIR /go/src/Gproject

RUN make gProject

ENTRYPOINT ["./demo"]