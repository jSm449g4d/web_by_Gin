
FROM golang

WORKDIR /
COPY . .

RUN go get github.com/gin-gonic/gin
RUN go get github.com/shirou/gopsutil/cpu
RUN go get github.com/shirou/gopsutil/mem

CMD go run *.go --ipport "0.0.0.0:8080"
