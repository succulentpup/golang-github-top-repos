# it wont' work as the main application has lot of pseudo code
FROM golang:1.11-buster
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get ioutil
RUN go get -u github.com/go-sql-driver/mysql
RUN go build -o main .
CMD ["./main"]
