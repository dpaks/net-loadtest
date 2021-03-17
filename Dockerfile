FROM ubuntu:20.04

RUN apt update && DEBIAN_FRONTEND=noninteractive TZ=Asia/Kolkata apt install -y \
    netcat \
    vim \
    golang-go \
    net-tools

ADD server.go /

ENTRYPOINT ["go"]
CMD ["run", "server.go"]
