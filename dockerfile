# Base image
FROM golang

WORKDIR D:\walt\project\costcogogo
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"

# Install Beego and Bee
RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/go-sql-driver/mysql && go get github.com/astaxie/beego/orm 
EXPOSE 8080

CMD ["bee", "run"]