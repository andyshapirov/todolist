FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./  
COPY internal/ ./internal/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todolist main.go


FROM alpine AS runner

WORKDIR /app

COPY --from=builder /app/todolist ./
COPY web/ ./web/

EXPOSE 7540

CMD ["./todolist"]