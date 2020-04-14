
FROM golang

# Init
ENV APP /var/www
WORKDIR $APP
COPY . .
RUN apt update -q;apt upgrade -yq

# Javascript
WORKDIR $APP/static
RUN apt install npm -yq
RUN npm install -g npm
RUN npm install --save jquery react react-dom react-bootstrap bootstrap firebase

# Golang
WORKDIR $APP/Gin
RUN apt install -yq wget golang gcc git curl
RUN go get github.com/gin-gonic/gin && \
    go get github.com/shirou/gopsutil/cpu && \
    go get github.com/shirou/gopsutil/mem && \
    go get github.com/artdarek/go-unzip

#server hosting
WORKDIR $APP
RUN go build -o ./gin.elf ./*.go
CMD bash -c "./gin.elf --IP_port=0.0.0.0:8080"
