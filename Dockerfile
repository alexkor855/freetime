FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/freetime

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Asia/Almaty /usr/share/zoneinfo/Asia/Almaty

ENV TZ Asia/Almaty

WORKDIR /app

COPY --from=builder /app/freetime /app/freetime

EXPOSE 1324

CMD ["./freetime"]