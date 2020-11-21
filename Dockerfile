FROM golang:1.15
ARG DEBIAN_MIRROR="huaweicloud"
#ARG DEBIAN_MIRROR="aliyun"
COPY . /build
WORKDIR /build/cli
RUN GO111MODULE="on" GOPROXY="https://goproxy.io" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o sshtunnel
RUN sed -i "s@http://ftp.debian.org@https://mirrors.$DEBIAN_MIRROR.com@g" /etc/apt/sources.list && \
sed -i "s@http://security.debian.org@https://mirrors.$DEBIAN_MIRROR.com@g" /etc/apt/sources.list && \
sed -i "s@http://deb.debian.org@https://mirrors.$DEBIAN_MIRROR.com@g" /etc/apt/sources.list && \
apt update && \
apt install -y upx
RUN upx sshtunnel

FROM alpine
ENV WAIT_VERSION 2.7.3
# ENV WAIT_RELEASE https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait
ENV WAIT_RELEASE https://st0n3-dev.obs.cn-south-1.myhuaweicloud.com/docker-compose-wait/release/$WAIT_VERSION/wait
ADD $WAIT_RELEASE /wait
RUN chmod +x /wait

RUN mkdir -p /app
COPY --from=0 /build/cli/sshtunnel /app/
WORKDIR /app
CMD /wait && ./sshtunnel