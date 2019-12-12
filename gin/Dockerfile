
FROM ubuntu:19.10

WORKDIR /
COPY . .

RUN apt update -q;apt upgrade -yq
RUN apt install -yq wget unzip golang gcc git

#download for fallback 
RUN wget -N https://github.com/twbs/bootstrap/releases/download/v4.4.1/bootstrap-4.4.1-dist.zip
RUN unzip bootstrap-4.4.1-dist.zip
RUN rm bootstrap-4.4.1-dist.zip
RUN mkdir -p ./static
RUN mv ./bootstrap-4.4.1-dist ./static/bootstrap-4.4.1-dist
RUN wget -N https://code.jquery.com/jquery-3.4.1.min.js
RUN mkdir -p ./static/bootstrap-4.4.1-dist/js
RUN mv ./jquery-3.4.1.min.js ./static/bootstrap-4.4.1-dist/js/jquery-3.4.1.min.js

RUN go get github.com/gin-gonic/gin
RUN go get github.com/shirou/gopsutil/cpu
RUN go get github.com/shirou/gopsutil/mem

CMD go run *.go --ipport "0.0.0.0:8080"
