FROM golang:1.23-alpine AS base

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY config/ config/
COPY internal/ internal/
COPY migrations/ migrations/
COPY pkg/ pkg/

# Build phase
FROM base AS api
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/*.go

FROM base AS migration
RUN CGO_ENABLED=0 GOOS=linux go build -o migration ./cmd/migration/*.go

# Final phase
FROM gcr.io/distroless/static-debian12:latest AS final

ENV TZ=Asia/Jakarta

COPY --from=api /build/api /api
COPY --from=migration /build/migration /migration
COPY --from=base /build/migrations/*.sql /data/migrations/