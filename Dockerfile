FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /arvancloud-challenge-app cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build /arvancloud-challenge-app ./

CMD ["./arvancloud-challenge-app"]

