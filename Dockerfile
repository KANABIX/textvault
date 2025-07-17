# Build stage
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o textvault

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/textvault .
COPY static/ static/
EXPOSE 8080
CMD ["./textvault"] 