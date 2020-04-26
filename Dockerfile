FROM golang:1.11.2 as dev

WORKDIR /go/src/github.com/

# Install Dependecies
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/go-sql-driver/mysql

COPY . .


#build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.9.2 as prod

WORKDIR /root/

RUN apk add --update mysql mysql-client && rm -f /var/cache/apk/*
# copy main binary
COPY --from=dev /go/src/github.com/main .

ENTRYPOINT ["./main"]
