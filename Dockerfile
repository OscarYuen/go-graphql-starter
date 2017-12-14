# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/OscarYuen/go-graphql-example/
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/neelance/graphql-go
RUN go get github.com/neelance/graphql-go/relay 
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/sqlite

RUN cd $SRC_DIR;go build -o go-server; cp go-server /app/
ENTRYPOINT ["./go-server"]