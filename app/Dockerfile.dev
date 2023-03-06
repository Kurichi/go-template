# for development
FROM golang:1.20.1-alpine3.17
RUN apk update && apk add --no-cache git && apk add --no-cache curl
ENV GO_ENV /go/src/app/
WORKDIR ${GO_ENV}
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# development tools
RUN go install github.com/cosmtrek/air@v1.29.0
CMD ["air", "-c", ".air.toml"]