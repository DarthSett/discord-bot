FROM golang:1.11.12-alpine3.10

RUN apk add --no-cache git alpine-sdk pkgconfig opus-dev opusfile-dev

WORKDIR /build

ADD . .
RUN go mod download
RUN go mod verify
RUN go build -o /build/bin .


FROM wernight/youtube-dl

RUN apk update
RUN youtube-dl -U
RUN apk add --no-cache ca-certificates opus-dev opusfile-dev

WORKDIR /bot

COPY --from=0 /build /server

ENTRYPOINT ["/server/bin"]