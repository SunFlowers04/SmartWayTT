FROM golang:1.22.12-alpine as builder
WORKDIR /build
COPY . .
ENV GOCACHE=/cache-docker/gocache
ENV GOMODCACHE=/cache-docker/gocache
RUN apk add -U --no-cache ca-certificates
RUN --mount=type=cache,target="/cache-docker/gocache" go mod tidy
COPY /.env.developer /main/.env
RUN --mount=type=cache,target="/cache-docker/gocache" CGO_ENABLED=0 GOOS=linux go build -o /main/build src/cmd/app/main.go
#ENTRYPOINT ["/main/build"]

FROM scratch
COPY --from=builder /main/.env /.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /main/build /bin/build
ENTRYPOINT ["/bin/build"]