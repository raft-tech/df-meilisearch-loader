#syntax=docker/dockerfile:1.4
ARG BASE_REGISTRY=registry1.dso.mil

FROM $BASE_REGISTRY/ironbank/google/golang/golang-1.20:1.20.8 AS builder

USER root

# Install Certificate
RUN dnf update -y && dnf install -y ca-certificates

WORKDIR /app

COPY . ./
RUN go mod download
RUN go build -o meilisearch_loader cmd/main.go

FROM $BASE_REGISTRY/ironbank/opensource/debian/debian:11.7

WORKDIR /app

COPY --from=builder /app/meilisearch_loader ./
RUN chmod -R +x ./
ENV GOPATH /app
RUN addgroup --gid 3000 appuser && adduser --system --no-create-home --uid 1001 --gecos "" --gid 3000 appuser
RUN chown -R appuser ./
RUN chmod -R u=rwx ./
USER appuser

EXPOSE 8080

CMD ["/app/meilisearch_loader"]
