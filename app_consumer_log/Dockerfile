FROM golang:1.15 AS builder

WORKDIR $GOPATH/src/app_consumer_log

COPY . ./

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN go get -u

RUN swag init

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM scratch

COPY --from=builder /go/src/app_consumer_log/main ./

EXPOSE 8080

ENTRYPOINT ["./main"]