FROM golang:1.16.4 AS go-build

ENV GO111MODULE=on

WORKDIR /go/src/gitlab.com/utopiops-water/application_name


COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o application_name .

FROM alpine:3.9.5
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=go-build /go/src/gitlab.com/utopiops-water/application_name .
ENTRYPOINT ["./application_name"]