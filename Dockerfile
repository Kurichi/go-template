FROM public.ecr.aws/docker/library/golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux go build -o main ./cmd/main.go

FROM public.ecr.aws/amazonlinux/amazonlinux:latest

COPY --from=builder /app/main /app/main

EXPOSE ${PORT}

CMD ["/app/main"]