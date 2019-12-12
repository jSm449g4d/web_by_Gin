
FROM python
COPY . .

FROM golang
COPY . .

WORKDIR /

RUN go get github.com/gin-gonic/gin
RUN go get github.com/shirou/gopsutil/cpu
RUN go get github.com/shirou/gopsutil/mem

#download for fallback 
RUN wget -N https://github.com/twbs/bootstrap/releases/download/v4.4.1/bootstrap-4.4.1-dist.zip
RUN unzip bootstrap-4.4.1-dist.zip
RUN rm bootstrap-4.4.1-dist.zip
RUN mv ./bootstrap-4.4.1-dist ./static/bootstrap-4.4.1-dist
RUN wget -N https://code.jquery.com/jquery-3.4.1.min.js
RUN mv ./jquery-3.4.1.min.js ./static/bootstrap-4.4.1-dist/js/jquery-3.4.1.min.js

CMD go run *.go --ipport "0.0.0.0:8080"
