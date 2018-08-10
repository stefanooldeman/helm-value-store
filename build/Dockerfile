FROM library/ubuntu:16.04

RUN apt update && apt-get install -y software-properties-common
RUN add-apt-repository ppa:gophers/archive
RUN apt update
RUN apt-get install -y golang-1.9-go git bzr mercurial libzookeeper-mt-dev
# libzookeeper-st-dev
RUN apt-get install curl

RUN ["/bin/bash", "-c", "mkdir -p /root/go/{bin,src,pkg}"]
ENV GOPATH=/root/go
ENV PATH="${GOPATH}/bin:/usr/lib/go-1.9/bin:${PATH}"
RUN ls -l /root/go/bin

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
