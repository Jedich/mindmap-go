FROM golang:alpine AS service_builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build ./cmd/server

FROM scratch

WORKDIR /bin/mindmap

EXPOSE 3000

COPY --from=service_builder --chmod=700 /build .
COPY ./config* ./config

CMD  ["./build"]