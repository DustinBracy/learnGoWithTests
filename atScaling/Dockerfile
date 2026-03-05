FROM golang:1.24-alpine

WORKDIR /app

ARG bin_to_build 

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr ./atScaling/cmd/${bin_to_build}/

CMD ["./svr"]