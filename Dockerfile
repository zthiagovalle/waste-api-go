FROM golang:1.20-alpine3.16 as base
RUN apk update
WORKDIR /usr/src/app
COPY . .
RUN go build -o server ./cmd/api

FROM alpine:3.16 as binary
COPY --from=base /usr/src/app .
EXPOSE 3000

CMD [ "./server" ]