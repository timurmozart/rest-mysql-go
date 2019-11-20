# 1
FROM golang:latest
# 2
WORKDIR /rest-mysql-go
# 3
COPY . /rest-mysql-go
# 4
RUN go get "github.com/go-sql-driver/mysql"
# 5
RUN go get "github.com/gorilla/mux"
# 6
RUN go build .
# 7
EXPOSE 8000
# 8
CMD ["/rest-mysql-go/rest-mysql-go"]
