FROM golang:latest as builder
WORKDIR /app
COPY . .

# Download dependencies 
RUN go mod download

# CGO_ENABLED=0 reduces size 
RUN CGO_ENABLED=0 GOOS=linux go build -o work-it .

# Stage 2: Create a minimal runtime image
FROM alpine:latest
WORKDIR /bin
COPY --from=builder /app/work-it .
CMD ["/bin/work-it"]
