FROM alpine:latest

MAINTAINER Pankaj Rai <pankaj.cdac@gmail.com>

WORKDIR "/opt"

ADD .docker_build/battleship /opt/bin/battleship

CMD ["/opt/bin/battleship"]

