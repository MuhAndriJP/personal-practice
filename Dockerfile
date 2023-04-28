FROM railwayapp/base:stable
RUN apt-get update && \
    apt-get install -y golang
# FROM golang:1.17-alpine
WORKDIR /gateway-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main
CMD ./main