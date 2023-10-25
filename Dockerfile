FROM golang:latest AS go

WORKDIR /app
ENV GOFLAGS=-buildvcs=false
COPY go.mod .
RUN go install github.com/cosmtrek/air@latest
RUN go mod download

FROM go AS app
WORKDIR /app
COPY . .
ENTRYPOINT ["air", "-c", "air.toml"]