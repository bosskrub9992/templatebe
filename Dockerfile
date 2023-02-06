FROM golang:1.20-alpine as builder

WORKDIR /app

RUN apk update && apk add git

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

ADD cmd ./cmd
ADD lib ./lib
ADD src ./src

WORKDIR /app/cmd
RUN go build -o /ms-service

FROM alpine:3.17

COPY --from=builder /ms-service .

# #Set default timezone
# RUN apk add --no-cache tzdata
# ENV TZ Asia/Bangkok
# RUN ln -sf /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
# RUN echo "Asia/Bangkok" > /etc/timezone

EXPOSE 8080

CMD [ "/ms-service" ]