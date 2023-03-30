FROM golang:latest
WORKDIR /app
COPY ./cmd/app/ .
EXPOSE 8080
CMD ["./main"]