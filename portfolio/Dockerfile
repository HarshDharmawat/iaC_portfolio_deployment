# we name this stage 'builder' so we can refer to it later
FROM golang:1.22-alpine AS builder

WORKDIR /app

# copy files and build
COPY . .

# we add CGO_ENABLED=0 to create a "statically linked" binary
# this guarantees it runs perfectly on ultra minimal environments
RUN CGO_ENABLED=0 go build -o portfolio .



# we start fresh with a tiny, 5MB alpine Linux image
FROM alpine:latest

WORKDIR /app

# we copy only the compiled 'portfolio' file 
COPY --from=builder /app/portfolio .

EXPOSE 8080

CMD ["./portfolio"]