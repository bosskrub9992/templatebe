FROM golang:1.20-alpine as builder

WORKDIR /app

RUN apk update && apk add git

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

ADD cmd ./cmd
ADD config ./config
ADD src ./src

RUN go build ./cmd/serve

FROM alpine:3.17

COPY --from=builder /app ./app

# # #Set default timezone
# # RUN apk add --no-cache tzdata
# # ENV TZ Asia/Bangkok
# # RUN ln -sf /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
# # RUN echo "Asia/Bangkok" > /etc/timezone

EXPOSE 8080

CMD [ "./app/serve" ]