FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
CMD [ "./main" ]