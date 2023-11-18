FROM golang:1.20.5

RUN go Install github.com/astaxie/beego && go get github.com/beego/bee
RUN go get github.com/astaxie/beego/orm
RUN go get github.com/go-sql-driver/mysql

EXPOSE 8080
# Install server application
CMD ["bee", "run"]