FROM centos:7

MAINTAINER lubyzheng@outlook.com

#ENV GOROOT /usr/local/go
#ENV GOPATH /code
#ENV PATH $PATH:$GOROOT/bin

RUN yum install -y java-1.8.0-openjdk* #Java
#RUN yum install -y curl gcc gcc-c++

#RUN curl -s -o go.tar.gz https://storage.googleapis.com/golang/go1.16.linux-amd64.tar.gz
#RUN tar --remove-files -C /usr/local/ -zxf go.tar.gz

RUN mkdir code

COPY ./code /code
COPY ./docker /
COPY ./input /input

#WORKDIR /code

ENTRYPOINT ["./execute"]
