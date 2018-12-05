FROM golang:1.11-alpine
RUN apk add --no-cache git gcc musl-dev
WORKDIR /src/app
COPY . .
RUN go mod download
RUN GOOS=linux GO111MODULE=on go build -o /app .

FROM alpine:3.8
COPY --from=0 /app /app
ENTRYPOINT ["/app"]