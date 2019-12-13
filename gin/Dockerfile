
FROM golang

WORKDIR /
COPY . .

RUN apt update -q;apt upgrade -yq
RUN apt install -yq wget unzip golang gcc git

RUN go get github.com/gin-gonic/gin
RUN go get github.com/shirou/gopsutil/cpu
RUN go get github.com/shirou/gopsutil/mem
RUN go get github.com/artdarek/go-unzip

RUN go build -o gin.elf *.go

CMD ./gin.elf --ipport "0.0.0.0:8080"
