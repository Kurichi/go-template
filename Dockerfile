# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22

FROM public.ecr.aws/docker/library/golang:${GO_VERSION} AS builder
WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  CGO_ENABLED=0 GOOS=linux go build -o /bin/server ./cmd/main.go

FROM alpine:latest as final

ARG UID=10001
RUN adduser \
  --disabled-password \ 
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  appuser
USER appuser

COPY --from=builder /bin/server /bin/

EXPOSE ${PORT}

ENTRYPOINT [ "/bin/server" ]