##### Stage 1 #####
FROM golang:1.16.5-alpine as builder

### Create new directly and set it as working directory
RUN mkdir -p /app
WORKDIR /app

### Copy Go application dependency files
COPY go.mod .
COPY go.sum .

### Download Go application module dependencies
RUN go mod download

### Copy source code for building the application
COPY . .

### Build the Go app for a linux OS
RUN GOOS=linux go build -o ./out/b2w-star-wars .

##### Stage 2 #####

### Define the running image
FROM alpine:3.13.1

### Set working directory
WORKDIR /application

RUN adduser -S -D -H -h /application appuser
USER appuser

### Copy built binary application from 'builder' image
COPY --from=builder /app/out/b2w-star-wars .

### Run the binary application
CMD ["./b2w-star-wars"]
