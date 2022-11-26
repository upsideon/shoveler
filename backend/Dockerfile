FROM golang:1.19-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod verify
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /shoveler

CMD ["/shoveler"]
