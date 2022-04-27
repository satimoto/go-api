FROM golang:alpine AS build-env

RUN mkdir /app
WORKDIR /app

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o /go/bin/app cmd/api/main.go

FROM scratch

COPY --from=build-env /go/bin/app /go/bin/app
EXPOSE 9000
CMD [ "/go/bin/app" ]