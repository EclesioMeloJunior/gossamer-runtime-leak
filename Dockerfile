FROM golang:1.19 AS go-base
WORKDIR /src
ARG CGO_ENABLED=1
ENV CGO_ENABLED="${CGO_ENABLED}"
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 6060
RUN GOOS=linux GOARCH=amd64 go build -o /src/build/app
ENTRYPOINT ["/src/build/app"]
