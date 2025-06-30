FROM golang:1 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /app

FROM gcr.io/distroless/static:nonroot

COPY --from=build /app /
CMD ["/app"]
