FROM golang:1.17-alpine as builder

WORKDIR /app/src/

COPY . .

RUN go build -o ./runner

FROM alpine:leates

WORKDIR /app/

COPY --from=builder ./runner .

ENTRYPOINT ["./runner"]
