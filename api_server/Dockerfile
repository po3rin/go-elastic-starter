FROM golang:1.11.1

WORKDIR /api_server

COPY . .
ENV GO111MODULE=on
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

RUN go get github.com/pilu/fresh
# CMD ["fresh"]
