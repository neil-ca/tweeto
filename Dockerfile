FROM golang:alpine AS build

WORKDIR /go/src/build

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./main ./main.go

FROM alpine:latest

COPY --from=build /go/src/build /go/src/build/.env /go/src/build/uploads ./

VOLUME /uploads
RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]
