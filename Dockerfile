# start with latest goland docker image
FROM golang:latest AS builder

LABEL maintainer="Sakshat <asakshat453@gmail.com>"

# current app dir
WORKDIR /app

# cp go mod and sum files
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# copy source to the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# start new contianer with postgres alpine image
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# copy the compiled binary from the first container (to reduce size)
COPY --from=builder /app/main .

# copy the .env file
COPY --from=builder /app/.env .

# open port 8080
EXPOSE 8080

# run the compiled go binary
CMD ["./main"] 