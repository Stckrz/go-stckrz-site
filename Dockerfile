# build stage

FROM golang:1.24.5-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web ./cmd/web

#runtime stage
FROM gcr.io/distroless/base-debian12

USER nonroot:nonroot

WORKDIR /app

COPY --from=build /src/web /app/web

# templates & static assets
COPY --from=build /src/internal/templates /app/internal/templates
COPY --from=build /src/public /app/public
COPY --from=build /src/internal/posts /app/internal/posts

EXPOSE 8080
CMD ["/app/web"]

