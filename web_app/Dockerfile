
FROM golang
#FROM ubuntu:19.10

#server initialize
ENV APP /web_app
WORKDIR $APP
COPY . .
RUN apt update -q;apt upgrade -yq

#npm
WORKDIR $APP/static
RUN apt install -yq npm
RUN npm install -g npm
RUN npm install --save jquery bootstrap

#golang
WORKDIR $APP
RUN apt install -yq wget golang gcc git curl
RUN go get github.com/gin-gonic/gin && \
    go get github.com/shirou/gopsutil/cpu && \
    go get github.com/shirou/gopsutil/mem && \
    go get github.com/artdarek/go-unzip
RUN go build -o ./gin.elf ./*.go

#server hosting
WORKDIR $APP
CMD ./gin.elf --IP_port "0.0.0.0:8080"
