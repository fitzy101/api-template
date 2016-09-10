# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.5.1

###### Install some unix toools that may be commonly used.
RUN echo "deb http://us.archive.ubuntu.com/ubuntu/ precise main universe" >> /etc/apt/source.list
RUN echo | apt-get update
RUN echo | apt-get install uuid-runtime
RUN rm -rf /var/lib/apt/lists/*

# add golibrary dependencies.
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/schema

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/fitzy101/api-template

# Build the api.
RUN go install github.com/fitzy101/api-template

# The entrypoint will now be the executable created in the previous command.
ENTRYPOINT /go/bin/api-template
