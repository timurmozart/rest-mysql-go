# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info

WORKDIR /app

RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"
RUN git clone https://github.com/timurmozart/rest-mysql-go
RUN cd rest-mysql-go
RUN ls
RUN go build .


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/go/rest-mysql-go/rest-mysql-go"]
